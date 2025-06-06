package history

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/dynamicconfig"
	"go.temporal.io/server/common/goro"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/membership"
	serviceerrors "go.temporal.io/server/common/serviceerror"
)

type (
	cacheEntry struct {
		shardID    int32
		address    rpcAddress
		connection clientConnection
		staleAt    time.Time
	}

	// A cachingRedirector is a redirector that maintains a cache of shard
	// owners, and uses that cache instead of querying membership for each
	// operation. Cache entries are evicted either for shard ownership lost
	// errors, or for any error that might indicate the history instance
	// is no longer available, including timeouts.
	cachingRedirector struct {
		mu struct {
			sync.RWMutex
			cache map[int32]cacheEntry
		}

		connections            connectionPool
		goros                  goro.Group
		historyServiceResolver membership.ServiceResolver
		logger                 log.Logger
		membershipUpdateCh     chan *membership.ChangedEvent
		staleTTL               dynamicconfig.DurationPropertyFn
	}
)

const cachingRedirectorListener = "cachingRedirectorListener"

func newCachingRedirector(
	connections connectionPool,
	historyServiceResolver membership.ServiceResolver,
	logger log.Logger,
	staleTTL dynamicconfig.DurationPropertyFn,
) *cachingRedirector {
	r := &cachingRedirector{
		connections:            connections,
		historyServiceResolver: historyServiceResolver,
		logger:                 logger,
		membershipUpdateCh:     make(chan *membership.ChangedEvent, 1),
		staleTTL:               staleTTL,
	}
	r.mu.cache = make(map[int32]cacheEntry)

	r.goros.Go(r.eventLoop)

	return r
}

func (r *cachingRedirector) stop() {
	r.goros.Cancel()
	r.goros.Wait()
}

func (r *cachingRedirector) clientForShardID(shardID int32) (historyservice.HistoryServiceClient, error) {
	if err := checkShardID(shardID); err != nil {
		return nil, err
	}
	entry, err := r.getOrCreateEntry(shardID)
	if err != nil {
		return nil, err
	}
	return entry.connection.historyClient, nil
}

func (r *cachingRedirector) execute(ctx context.Context, shardID int32, op clientOperation) error {
	if err := checkShardID(shardID); err != nil {
		return err
	}
	opEntry, err := r.getOrCreateEntry(shardID)
	if err != nil {
		return err
	}
	return r.redirectLoop(ctx, opEntry, op)
}

func (r *cachingRedirector) redirectLoop(ctx context.Context, opEntry cacheEntry, op clientOperation) error {
	for {
		if err := common.IsValidContext(ctx); err != nil {
			return err
		}
		opErr := op(ctx, opEntry.connection.historyClient)
		if opErr == nil {
			return opErr
		}
		if maybeHostDownError(opErr) {
			r.cacheDeleteByAddress(opEntry.address)
			return opErr
		}
		var solErr *serviceerrors.ShardOwnershipLost
		if !errors.As(opErr, &solErr) {
			return opErr
		}
		var again bool
		opEntry, again = r.handleSolError(opEntry, solErr)
		if !again {
			return opErr
		}
	}
}

func (r *cachingRedirector) getOrCreateEntry(shardID int32) (cacheEntry, error) {
	r.mu.RLock()
	entry, ok := r.mu.cache[shardID]
	r.mu.RUnlock()
	if ok {
		if entry.staleAt.IsZero() || time.Now().Before(entry.staleAt) {
			return entry, nil
		}
		// Otherwise, check below under write lock.
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// Recheck under write lock.
	entry, ok = r.mu.cache[shardID]
	if ok {
		if entry.staleAt.IsZero() || time.Now().Before(entry.staleAt) {
			return entry, nil
		}
		// Delete and fallthrough below to re-check ownership.
		delete(r.mu.cache, shardID)
	}

	address, err := shardLookup(r.historyServiceResolver, shardID)
	if err != nil {
		return cacheEntry{}, err
	}

	return r.cacheAddLocked(shardID, address), nil
}

func (r *cachingRedirector) cacheAddLocked(shardID int32, addr rpcAddress) cacheEntry {
	// New history instances might reuse the address of a previously live history
	// instance. Since we don't currently close GRPC connections when they become
	// unused or idle, we might have a GRPC connection that has gone into its
	// connection backoff state, due to the previous history instance becoming
	// unreachable. A request on the GRPC connection, intended for the new history
	// instance, would be delayed waiting for the next connection attempt, which
	// could be many seconds.
	// If we're adding a new cache entry for a shard, we take that as a hint that
	// the next request should attempt to connect immediately if required. If the
	// GRPC connection is not in connect backoff, this call has no effect.
	connection := r.connections.getOrCreateClientConn(addr)
	r.connections.resetConnectBackoff(connection)

	entry := cacheEntry{
		shardID:    shardID,
		address:    addr,
		connection: connection,
		// staleAt is left at zero; it's only set when r.staleTTL is set,
		// and after a membership update informs us that this address is no
		// longer the shard owner.
	}
	r.mu.cache[shardID] = entry

	return entry
}

func (r *cachingRedirector) cacheDeleteByAddress(address rpcAddress) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for shardID, entry := range r.mu.cache {
		if entry.address == address {
			delete(r.mu.cache, shardID)
		}
	}
}

func (r *cachingRedirector) handleSolError(opEntry cacheEntry, solErr *serviceerrors.ShardOwnershipLost) (cacheEntry, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cached, ok := r.mu.cache[opEntry.shardID]; ok {
		if cached.address == opEntry.address {
			delete(r.mu.cache, cached.shardID)
		}
	}

	solErrNewOwner := rpcAddress(solErr.OwnerHost)
	if len(solErrNewOwner) != 0 && solErrNewOwner != opEntry.address {
		r.logger.Info("historyClient: updating cache from shard ownership lost error",
			tag.ShardID(opEntry.shardID),
			tag.NewAnyTag("oldAddress", opEntry.address),
			tag.NewAnyTag("newAddress", solErrNewOwner))
		return r.cacheAddLocked(opEntry.shardID, solErrNewOwner), true
	}

	return cacheEntry{}, false
}

func maybeHostDownError(opErr error) bool {
	var unavail *serviceerror.Unavailable
	if errors.As(opErr, &unavail) {
		return true
	}
	return common.IsContextDeadlineExceededErr(opErr)
}

func (r *cachingRedirector) eventLoop(ctx context.Context) error {
	if err := r.historyServiceResolver.AddListener(cachingRedirectorListener, r.membershipUpdateCh); err != nil {
		r.logger.Fatal("Error adding listener", tag.Error(err))
	}
	defer func() {
		if err := r.historyServiceResolver.RemoveListener(cachingRedirectorListener); err != nil {
			r.logger.Warn("Error removing listener", tag.Error(err))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-r.membershipUpdateCh:
			r.staleCheck()
		}
	}
}

func (r *cachingRedirector) staleCheck() {
	staleTTL := r.staleTTL()

	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	for shardID, entry := range r.mu.cache {
		if !entry.staleAt.IsZero() {
			if now.After(entry.staleAt) {
				delete(r.mu.cache, shardID)
			}
			continue
		}
		if staleTTL > 0 {
			addr, err := shardLookup(r.historyServiceResolver, shardID)
			if err != nil || addr != entry.address {
				entry.staleAt = now.Add(staleTTL)
				r.mu.cache[shardID] = entry
			}
		}
	}
}
