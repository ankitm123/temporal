// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/server/api/token/v1/message.proto

package token

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	v12 "go.temporal.io/server/api/clock/v1"
	v1 "go.temporal.io/server/api/history/v1"
	v11 "go.temporal.io/server/api/persistence/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HistoryContinuation struct {
	state               protoimpl.MessageState   `protogen:"open.v1"`
	RunId               string                   `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	FirstEventId        int64                    `protobuf:"varint,2,opt,name=first_event_id,json=firstEventId,proto3" json:"first_event_id,omitempty"`
	NextEventId         int64                    `protobuf:"varint,3,opt,name=next_event_id,json=nextEventId,proto3" json:"next_event_id,omitempty"`
	IsWorkflowRunning   bool                     `protobuf:"varint,5,opt,name=is_workflow_running,json=isWorkflowRunning,proto3" json:"is_workflow_running,omitempty"`
	PersistenceToken    []byte                   `protobuf:"bytes,6,opt,name=persistence_token,json=persistenceToken,proto3" json:"persistence_token,omitempty"`
	BranchToken         []byte                   `protobuf:"bytes,8,opt,name=branch_token,json=branchToken,proto3" json:"branch_token,omitempty"`
	VersionHistoryItem  *v1.VersionHistoryItem   `protobuf:"bytes,10,opt,name=version_history_item,json=versionHistoryItem,proto3" json:"version_history_item,omitempty"`
	VersionedTransition *v11.VersionedTransition `protobuf:"bytes,11,opt,name=versioned_transition,json=versionedTransition,proto3" json:"versioned_transition,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *HistoryContinuation) Reset() {
	*x = HistoryContinuation{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryContinuation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryContinuation) ProtoMessage() {}

func (x *HistoryContinuation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryContinuation.ProtoReflect.Descriptor instead.
func (*HistoryContinuation) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryContinuation) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *HistoryContinuation) GetFirstEventId() int64 {
	if x != nil {
		return x.FirstEventId
	}
	return 0
}

func (x *HistoryContinuation) GetNextEventId() int64 {
	if x != nil {
		return x.NextEventId
	}
	return 0
}

func (x *HistoryContinuation) GetIsWorkflowRunning() bool {
	if x != nil {
		return x.IsWorkflowRunning
	}
	return false
}

func (x *HistoryContinuation) GetPersistenceToken() []byte {
	if x != nil {
		return x.PersistenceToken
	}
	return nil
}

func (x *HistoryContinuation) GetBranchToken() []byte {
	if x != nil {
		return x.BranchToken
	}
	return nil
}

func (x *HistoryContinuation) GetVersionHistoryItem() *v1.VersionHistoryItem {
	if x != nil {
		return x.VersionHistoryItem
	}
	return nil
}

func (x *HistoryContinuation) GetVersionedTransition() *v11.VersionedTransition {
	if x != nil {
		return x.VersionedTransition
	}
	return nil
}

type RawHistoryContinuation struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	NamespaceId       string                 `protobuf:"bytes,10,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	WorkflowId        string                 `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId             string                 `protobuf:"bytes,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	StartEventId      int64                  `protobuf:"varint,4,opt,name=start_event_id,json=startEventId,proto3" json:"start_event_id,omitempty"`
	StartEventVersion int64                  `protobuf:"varint,5,opt,name=start_event_version,json=startEventVersion,proto3" json:"start_event_version,omitempty"`
	EndEventId        int64                  `protobuf:"varint,6,opt,name=end_event_id,json=endEventId,proto3" json:"end_event_id,omitempty"`
	EndEventVersion   int64                  `protobuf:"varint,7,opt,name=end_event_version,json=endEventVersion,proto3" json:"end_event_version,omitempty"`
	PersistenceToken  []byte                 `protobuf:"bytes,8,opt,name=persistence_token,json=persistenceToken,proto3" json:"persistence_token,omitempty"`
	VersionHistories  *v1.VersionHistories   `protobuf:"bytes,9,opt,name=version_histories,json=versionHistories,proto3" json:"version_histories,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *RawHistoryContinuation) Reset() {
	*x = RawHistoryContinuation{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawHistoryContinuation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawHistoryContinuation) ProtoMessage() {}

func (x *RawHistoryContinuation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawHistoryContinuation.ProtoReflect.Descriptor instead.
func (*RawHistoryContinuation) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *RawHistoryContinuation) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *RawHistoryContinuation) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *RawHistoryContinuation) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *RawHistoryContinuation) GetStartEventId() int64 {
	if x != nil {
		return x.StartEventId
	}
	return 0
}

func (x *RawHistoryContinuation) GetStartEventVersion() int64 {
	if x != nil {
		return x.StartEventVersion
	}
	return 0
}

func (x *RawHistoryContinuation) GetEndEventId() int64 {
	if x != nil {
		return x.EndEventId
	}
	return 0
}

func (x *RawHistoryContinuation) GetEndEventVersion() int64 {
	if x != nil {
		return x.EndEventVersion
	}
	return 0
}

func (x *RawHistoryContinuation) GetPersistenceToken() []byte {
	if x != nil {
		return x.PersistenceToken
	}
	return nil
}

func (x *RawHistoryContinuation) GetVersionHistories() *v1.VersionHistories {
	if x != nil {
		return x.VersionHistories
	}
	return nil
}

type Task struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	NamespaceId      string                 `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	WorkflowId       string                 `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId            string                 `protobuf:"bytes,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	ScheduledEventId int64                  `protobuf:"varint,4,opt,name=scheduled_event_id,json=scheduledEventId,proto3" json:"scheduled_event_id,omitempty"`
	Attempt          int32                  `protobuf:"varint,5,opt,name=attempt,proto3" json:"attempt,omitempty"`
	ActivityId       string                 `protobuf:"bytes,6,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	WorkflowType     string                 `protobuf:"bytes,7,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type,omitempty"`
	ActivityType     string                 `protobuf:"bytes,8,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	Clock            *v12.VectorClock       `protobuf:"bytes,9,opt,name=clock,proto3" json:"clock,omitempty"`
	StartedEventId   int64                  `protobuf:"varint,10,opt,name=started_event_id,json=startedEventId,proto3" json:"started_event_id,omitempty"`
	Version          int64                  `protobuf:"varint,11,opt,name=version,proto3" json:"version,omitempty"`
	StartedTime      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=started_time,json=startedTime,proto3" json:"started_time,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *Task) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *Task) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *Task) GetScheduledEventId() int64 {
	if x != nil {
		return x.ScheduledEventId
	}
	return 0
}

func (x *Task) GetAttempt() int32 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *Task) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *Task) GetWorkflowType() string {
	if x != nil {
		return x.WorkflowType
	}
	return ""
}

func (x *Task) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *Task) GetClock() *v12.VectorClock {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *Task) GetStartedEventId() int64 {
	if x != nil {
		return x.StartedEventId
	}
	return 0
}

func (x *Task) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Task) GetStartedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedTime
	}
	return nil
}

type QueryTask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NamespaceId   string                 `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	TaskQueue     string                 `protobuf:"bytes,2,opt,name=task_queue,json=taskQueue,proto3" json:"task_queue,omitempty"`
	TaskId        string                 `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryTask) Reset() {
	*x = QueryTask{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTask) ProtoMessage() {}

func (x *QueryTask) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTask.ProtoReflect.Descriptor instead.
func (*QueryTask) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *QueryTask) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *QueryTask) GetTaskQueue() string {
	if x != nil {
		return x.TaskQueue
	}
	return ""
}

func (x *QueryTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type NexusTask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NamespaceId   string                 `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	TaskQueue     string                 `protobuf:"bytes,2,opt,name=task_queue,json=taskQueue,proto3" json:"task_queue,omitempty"`
	TaskId        string                 `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NexusTask) Reset() {
	*x = NexusTask{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NexusTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusTask) ProtoMessage() {}

func (x *NexusTask) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusTask.ProtoReflect.Descriptor instead.
func (*NexusTask) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *NexusTask) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *NexusTask) GetTaskQueue() string {
	if x != nil {
		return x.TaskQueue
	}
	return ""
}

func (x *NexusTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

// A reference for loading a history event.
type HistoryEventRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Event ID.
	EventId int64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Event batch ID - the first event ID in the batch the event was stored in.
	EventBatchId  int64 `protobuf:"varint,2,opt,name=event_batch_id,json=eventBatchId,proto3" json:"event_batch_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoryEventRef) Reset() {
	*x = HistoryEventRef{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryEventRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryEventRef) ProtoMessage() {}

func (x *HistoryEventRef) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryEventRef.ProtoReflect.Descriptor instead.
func (*HistoryEventRef) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryEventRef) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *HistoryEventRef) GetEventBatchId() int64 {
	if x != nil {
		return x.EventBatchId
	}
	return 0
}

// A completion token for a Nexus operation started from a workflow.
type NexusOperationCompletion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Namespace UUID.
	NamespaceId string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	// Workflow ID.
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Run ID at the time this token was generated.
	RunId string `protobuf:"bytes,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// Reference including the path to the backing Operation state machine and a version + transition count for
	// staleness checks.
	Ref *v11.StateMachineRef `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// Request ID embedded in the NexusOperationScheduledEvent.
	// Allows completing a started operation after a workflow has been reset.
	RequestId     string `protobuf:"bytes,5,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NexusOperationCompletion) Reset() {
	*x = NexusOperationCompletion{}
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NexusOperationCompletion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusOperationCompletion) ProtoMessage() {}

func (x *NexusOperationCompletion) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_token_v1_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusOperationCompletion.ProtoReflect.Descriptor instead.
func (*NexusOperationCompletion) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_token_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *NexusOperationCompletion) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *NexusOperationCompletion) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *NexusOperationCompletion) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *NexusOperationCompletion) GetRef() *v11.StateMachineRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *NexusOperationCompletion) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_temporal_server_api_token_v1_message_proto protoreflect.FileDescriptor

const file_temporal_server_api_token_v1_message_proto_rawDesc = "" +
	"\n" +
	"*temporal/server/api/token/v1/message.proto\x12\x1ctemporal.server.api.token.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a*temporal/server/api/clock/v1/message.proto\x1a,temporal/server/api/history/v1/message.proto\x1a,temporal/server/api/persistence/v1/hsm.proto\"\xd4\x03\n" +
	"\x13HistoryContinuation\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12$\n" +
	"\x0efirst_event_id\x18\x02 \x01(\x03R\ffirstEventId\x12\"\n" +
	"\rnext_event_id\x18\x03 \x01(\x03R\vnextEventId\x12.\n" +
	"\x13is_workflow_running\x18\x05 \x01(\bR\x11isWorkflowRunning\x12+\n" +
	"\x11persistence_token\x18\x06 \x01(\fR\x10persistenceToken\x12!\n" +
	"\fbranch_token\x18\b \x01(\fR\vbranchToken\x12d\n" +
	"\x14version_history_item\x18\n" +
	" \x01(\v22.temporal.server.api.history.v1.VersionHistoryItemR\x12versionHistoryItem\x12j\n" +
	"\x14versioned_transition\x18\v \x01(\v27.temporal.server.api.persistence.v1.VersionedTransitionR\x13versionedTransitionJ\x04\b\a\x10\bJ\x04\b\t\x10\n" +
	"\"\xa9\x03\n" +
	"\x16RawHistoryContinuation\x12!\n" +
	"\fnamespace_id\x18\n" +
	" \x01(\tR\vnamespaceId\x12\x1f\n" +
	"\vworkflow_id\x18\x02 \x01(\tR\n" +
	"workflowId\x12\x15\n" +
	"\x06run_id\x18\x03 \x01(\tR\x05runId\x12$\n" +
	"\x0estart_event_id\x18\x04 \x01(\x03R\fstartEventId\x12.\n" +
	"\x13start_event_version\x18\x05 \x01(\x03R\x11startEventVersion\x12 \n" +
	"\fend_event_id\x18\x06 \x01(\x03R\n" +
	"endEventId\x12*\n" +
	"\x11end_event_version\x18\a \x01(\x03R\x0fendEventVersion\x12+\n" +
	"\x11persistence_token\x18\b \x01(\fR\x10persistenceToken\x12]\n" +
	"\x11version_histories\x18\t \x01(\v20.temporal.server.api.history.v1.VersionHistoriesR\x10versionHistoriesJ\x04\b\x01\x10\x02\"\xd8\x03\n" +
	"\x04Task\x12!\n" +
	"\fnamespace_id\x18\x01 \x01(\tR\vnamespaceId\x12\x1f\n" +
	"\vworkflow_id\x18\x02 \x01(\tR\n" +
	"workflowId\x12\x15\n" +
	"\x06run_id\x18\x03 \x01(\tR\x05runId\x12,\n" +
	"\x12scheduled_event_id\x18\x04 \x01(\x03R\x10scheduledEventId\x12\x18\n" +
	"\aattempt\x18\x05 \x01(\x05R\aattempt\x12\x1f\n" +
	"\vactivity_id\x18\x06 \x01(\tR\n" +
	"activityId\x12#\n" +
	"\rworkflow_type\x18\a \x01(\tR\fworkflowType\x12#\n" +
	"\ractivity_type\x18\b \x01(\tR\factivityType\x12?\n" +
	"\x05clock\x18\t \x01(\v2).temporal.server.api.clock.v1.VectorClockR\x05clock\x12(\n" +
	"\x10started_event_id\x18\n" +
	" \x01(\x03R\x0estartedEventId\x12\x18\n" +
	"\aversion\x18\v \x01(\x03R\aversion\x12=\n" +
	"\fstarted_time\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\vstartedTime\"f\n" +
	"\tQueryTask\x12!\n" +
	"\fnamespace_id\x18\x01 \x01(\tR\vnamespaceId\x12\x1d\n" +
	"\n" +
	"task_queue\x18\x02 \x01(\tR\ttaskQueue\x12\x17\n" +
	"\atask_id\x18\x03 \x01(\tR\x06taskId\"f\n" +
	"\tNexusTask\x12!\n" +
	"\fnamespace_id\x18\x01 \x01(\tR\vnamespaceId\x12\x1d\n" +
	"\n" +
	"task_queue\x18\x02 \x01(\tR\ttaskQueue\x12\x17\n" +
	"\atask_id\x18\x03 \x01(\tR\x06taskId\"R\n" +
	"\x0fHistoryEventRef\x12\x19\n" +
	"\bevent_id\x18\x01 \x01(\x03R\aeventId\x12$\n" +
	"\x0eevent_batch_id\x18\x02 \x01(\x03R\feventBatchId\"\xdb\x01\n" +
	"\x18NexusOperationCompletion\x12!\n" +
	"\fnamespace_id\x18\x01 \x01(\tR\vnamespaceId\x12\x1f\n" +
	"\vworkflow_id\x18\x02 \x01(\tR\n" +
	"workflowId\x12\x15\n" +
	"\x06run_id\x18\x03 \x01(\tR\x05runId\x12E\n" +
	"\x03ref\x18\x04 \x01(\v23.temporal.server.api.persistence.v1.StateMachineRefR\x03ref\x12\x1d\n" +
	"\n" +
	"request_id\x18\x05 \x01(\tR\trequestIdB*Z(go.temporal.io/server/api/token/v1;tokenb\x06proto3"

var (
	file_temporal_server_api_token_v1_message_proto_rawDescOnce sync.Once
	file_temporal_server_api_token_v1_message_proto_rawDescData []byte
)

func file_temporal_server_api_token_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_server_api_token_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_token_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_server_api_token_v1_message_proto_rawDesc), len(file_temporal_server_api_token_v1_message_proto_rawDesc)))
	})
	return file_temporal_server_api_token_v1_message_proto_rawDescData
}

var file_temporal_server_api_token_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_temporal_server_api_token_v1_message_proto_goTypes = []any{
	(*HistoryContinuation)(nil),      // 0: temporal.server.api.token.v1.HistoryContinuation
	(*RawHistoryContinuation)(nil),   // 1: temporal.server.api.token.v1.RawHistoryContinuation
	(*Task)(nil),                     // 2: temporal.server.api.token.v1.Task
	(*QueryTask)(nil),                // 3: temporal.server.api.token.v1.QueryTask
	(*NexusTask)(nil),                // 4: temporal.server.api.token.v1.NexusTask
	(*HistoryEventRef)(nil),          // 5: temporal.server.api.token.v1.HistoryEventRef
	(*NexusOperationCompletion)(nil), // 6: temporal.server.api.token.v1.NexusOperationCompletion
	(*v1.VersionHistoryItem)(nil),    // 7: temporal.server.api.history.v1.VersionHistoryItem
	(*v11.VersionedTransition)(nil),  // 8: temporal.server.api.persistence.v1.VersionedTransition
	(*v1.VersionHistories)(nil),      // 9: temporal.server.api.history.v1.VersionHistories
	(*v12.VectorClock)(nil),          // 10: temporal.server.api.clock.v1.VectorClock
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
	(*v11.StateMachineRef)(nil),      // 12: temporal.server.api.persistence.v1.StateMachineRef
}
var file_temporal_server_api_token_v1_message_proto_depIdxs = []int32{
	7,  // 0: temporal.server.api.token.v1.HistoryContinuation.version_history_item:type_name -> temporal.server.api.history.v1.VersionHistoryItem
	8,  // 1: temporal.server.api.token.v1.HistoryContinuation.versioned_transition:type_name -> temporal.server.api.persistence.v1.VersionedTransition
	9,  // 2: temporal.server.api.token.v1.RawHistoryContinuation.version_histories:type_name -> temporal.server.api.history.v1.VersionHistories
	10, // 3: temporal.server.api.token.v1.Task.clock:type_name -> temporal.server.api.clock.v1.VectorClock
	11, // 4: temporal.server.api.token.v1.Task.started_time:type_name -> google.protobuf.Timestamp
	12, // 5: temporal.server.api.token.v1.NexusOperationCompletion.ref:type_name -> temporal.server.api.persistence.v1.StateMachineRef
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_temporal_server_api_token_v1_message_proto_init() }
func file_temporal_server_api_token_v1_message_proto_init() {
	if File_temporal_server_api_token_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_server_api_token_v1_message_proto_rawDesc), len(file_temporal_server_api_token_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_token_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_token_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_token_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_server_api_token_v1_message_proto = out.File
	file_temporal_server_api_token_v1_message_proto_goTypes = nil
	file_temporal_server_api_token_v1_message_proto_depIdxs = nil
}
