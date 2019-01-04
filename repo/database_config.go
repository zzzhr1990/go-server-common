package repo

import (
	"github.com/jinzhu/gorm"
)

//DatabaseConfig used to create database connection.
type DatabaseConfig struct {
	SlaveDatabaseConnectionString  []string
	MasterDatabaseConnectionString string
	DatabaseDialect                string
	ReadWriteSeparate              bool
	Sharding                       bool
	TableNamePrefix                string
	DbConnection                   *gorm.DB
}
