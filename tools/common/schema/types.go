package schema

import (
	"fmt"
	"regexp"
)

type (
	// ConfigError is an error type that
	// represents a problem with the config
	ConfigError struct {
		msg string
	}
	// UpdateConfig holds the config
	// params for executing a UpdateTask
	UpdateConfig struct {
		DBName        string
		TargetVersion string
		SchemaDir     string
		SchemaName    string
		IsDryRun      bool
	}
	// SetupConfig holds the config
	// params need by the SetupTask
	SetupConfig struct {
		SchemaFilePath    string
		SchemaName        string
		InitialVersion    string
		Overwrite         bool // overwrite previous data
		DisableVersioning bool // do not use schema versioning
	}

	// DB is the database interface that's required to be implemented
	// for the schema-tool to work
	DB interface {
		// Exec executes a cql statement
		Exec(stmt string, args ...interface{}) error
		// DropAllTables drops all tables
		DropAllTables() error
		// CreateSchemaVersionTables sets up the schema version tables
		CreateSchemaVersionTables() error
		// ReadSchemaVersion returns the current schema version for the keyspace
		ReadSchemaVersion() (string, error)
		// UpdateSchemaVersion updates the schema version for the keyspace
		UpdateSchemaVersion(newVersion string, minCompatibleVersion string) error
		// WriteSchemaUpdateLog adds an entry to the schema update history table
		WriteSchemaUpdateLog(oldVersion string, newVersion string, manifestMD5 string, desc string) error
		// Close gracefully closes the client object
		Close()
		// Type gives the type of db (e.g. "cassandra", "sql")
		Type() string
	}
)

const (
	// CLIOptEndpoint is the cli option for endpoint
	CLIOptEndpoint = "endpoint"
	// CLIOptPort is the cli option for port
	CLIOptPort = "port"
	// CLIOptUser is the cli option for user
	CLIOptUser = "user"
	// CLIOptPassword is the cli option for password
	CLIOptPassword = "password"
	// CLIOptAuthenticator is the cli option for allowed authenticator settings
	CLIOptAllowedAuthenticators = "allowed-authenticators"
	// CLIOptTimeout is the cli option for timeout
	CLIOptTimeout = "timeout"
	// CLIOptKeyspace is the cli option for keyspace
	CLIOptKeyspace = "keyspace"
	// CLIOptDatabase is the cli option for database
	CLIOptDatabase = "database"
	// CLIOptDefaultDb is the cli option used as defaultdb to connect to
	CLIOptDefaultDb = "defaultdb"
	// CLIOptPluginName is the cli option for plugin name
	CLIOptPluginName = "plugin"
	// CLIOptConnectAttributes is the cli option for connect attributes (key/values via a url query string)
	CLIOptConnectAttributes = "connect-attributes"
	// CLIOptVersion is the cli option for version
	CLIOptVersion = "version"
	// CLIOptSchemaFile is the cli option for schema file
	CLIOptSchemaFile = "schema-file"
	// CLIOptOverwrite is the cli option for overwrite
	CLIOptOverwrite = "overwrite"
	// CLIOptDisableVersioning is the cli option to disabling versioning
	CLIOptDisableVersioning = "disable-versioning"
	// CLIOptTargetVersion is the cli option for target version
	CLIOptTargetVersion = "version"
	// CLIOptSchemaDir is the cli option for schema directory
	CLIOptSchemaDir = "schema-dir"
	// CLIOptSchemaName is the cli option for which pre-embedded schema to use
	CLIOptSchemaName = "schema-name"
	// CLIOptReplicationFactor is the cli option for replication factor
	CLIOptReplicationFactor = "replication-factor"
	// CLIOptDatacenter is the cli option for NetworkTopologyStrategy datacenter
	CLIOptDatacenter = "datacenter"
	// CLIOptConsistency is the cli option for consistency settings
	CLIOptConsistency = "consistency"
	// CLIOptAddressTranslator is the cli option for address translator for Cassandra
	CLIOptAddressTranslator = "address-translator"
	// CLIOptAddressTranslatorOptions is the cli option for options for address translator
	CLIOptAddressTranslatorOptions = "address-translator-options"
	// CLIOptQuiet is the cli option for quiet mode
	CLIOptQuiet = "quiet"
	// CLIOptForce is the cli option for force mode
	CLIOptForce = "force"

	// CLIFlagEndpoint is the cli flag for endpoint
	CLIFlagEndpoint = CLIOptEndpoint + ", ep"
	// CLIFlagPort is the cli flag for port
	CLIFlagPort = CLIOptPort + ", p"
	// CLIFlagUser is the cli flag for user
	CLIFlagUser = CLIOptUser + ", u"
	// CLIFlagPassword is the cli flag for password
	CLIFlagPassword = CLIOptPassword + ", pw"
	// CLIFlagAllowedAuthenticators is the cli flag for allowed authenticators
	CLIFlagAllowedAuthenticators = CLIOptAllowedAuthenticators + ", aa"
	// CLIFlagTimeout is the cli flag for timeout
	CLIFlagTimeout = CLIOptTimeout + ", t"
	// CLIFlagKeyspace is the cli flag for keyspace
	CLIFlagKeyspace = CLIOptKeyspace + ", k"
	// CLIFlagDatabase is the cli flag for database
	CLIFlagDatabase = CLIOptDatabase + ", db"
	// CLIFlagPluginName is the cli flag for plugin name
	CLIFlagPluginName = CLIOptPluginName + ", pl"
	// CLIFlagConnectAttributes allows arbitrary connect attributes
	CLIFlagConnectAttributes = CLIOptConnectAttributes + ", ca"
	// CLIFlagVersion is the cli flag for version
	CLIFlagVersion = CLIOptVersion + ", v"
	// CLIFlagSchemaFile is the cli flag for schema file
	CLIFlagSchemaFile = CLIOptSchemaFile + ", f"
	// CLIFlagOverwrite is the cli flag for overwrite
	CLIFlagOverwrite = CLIOptOverwrite + ", o"
	// CLIFlagDisableVersioning is the cli flag for disabling versioning
	CLIFlagDisableVersioning = CLIOptDisableVersioning + ", d"
	// CLIFlagTargetVersion is the cli flag for target version
	CLIFlagTargetVersion = CLIOptTargetVersion + ", v"
	// CLIFlagSchemaDir is the cli flag for schema directory
	CLIFlagSchemaDir = CLIOptSchemaDir + ", d"
	// CLIFlagSchemaName is the cli flag that says which pre-embedded schema to use
	CLIFlagSchemaName = CLIOptSchemaName + ", s"
	// CLIFlagReplicationFactor is the cli flag for replication factor
	CLIFlagReplicationFactor = CLIOptReplicationFactor + ", rf"
	// CLIFlagDatacenter is the cli option for NetworkTopologyStrategy datacenter
	CLIFlagDatacenter = CLIOptDatacenter + ", dc"
	// CLIFlagAddressTranslator is the cli option for address translator for Cassandra
	CLIFlagAddressTranslator = CLIOptAddressTranslator + ", at"
	// CLIFlagAddressTranslatorOptions is the cli option for address translator of Cassandra
	CLIFlagAddressTranslatorOptions
	// CLIFlagQuiet is the cli flag for quiet mode
	CLIFlagQuiet = CLIOptQuiet + ", q"
	// CLIFlagForce is the cli flag for force mode
	CLIFlagForce = CLIOptForce + ", f"
	// CLIFlagDisableInitialHostLookup is the cli flag for only using supplied hosts to connect to the database
	CLIFlagDisableInitialHostLookup = "disable-initial-host-lookup"

	// CLIFlagEnableTLS enables cassandra client TLS
	CLIFlagEnableTLS = "tls"
	// CLIFlagTLSCertFile is the optional tls cert file (tls must be enabled)
	CLIFlagTLSCertFile = "tls-cert-file"
	// CLIFlagTLSKeyFile is the optional tls key file (tls must be enabled)
	CLIFlagTLSKeyFile = "tls-key-file"
	// CLIFlagTLSCaFile is the optional tls CA file (tls must be enabled)
	CLIFlagTLSCaFile = "tls-ca-file"
	// CLIFlagTLSDisableHostVerification disable tls host verification (tls must be enabled)
	CLIFlagTLSDisableHostVerification = "tls-disable-host-verification"
	// CLIFlagTLSHostName specifies the host name for host name verification
	CLIFlagTLSHostName = "tls-server-name"
)

var rmspaceRegex = regexp.MustCompile(`\s+`)

// NewConfigError creates and returns an instance of ConfigError
func NewConfigError(msg string) error {
	return &ConfigError{msg: msg}
}

// Error returns a string representation of this error
func (e *ConfigError) Error() string {
	return fmt.Sprintf("Config Error:%v", e.msg)
}
