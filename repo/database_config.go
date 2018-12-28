package repo

//DatabaseConfig used to create database connection.
type DatabaseConfig struct {
	SlaveDatabaseConnectionString  []string
	MasterDatabaseConnectionString string
	DatabaseDialect                string
	ReadWriteSeparate              bool
	Sharding                       bool
	TableNamePrefix                string
}
