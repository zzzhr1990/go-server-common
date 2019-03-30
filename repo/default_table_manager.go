package repo

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

// DefaultTableManager default 3
type DefaultTableManager struct {
	//masterDBs map[string]*gorm.DB
	Config       *DatabaseConfig
	DefaultTable *gorm.DB
}

//GetTable Imp get default table3
func (manager *DefaultTableManager) GetTable(entity interface{}, read bool) (*gorm.DB, error) {

	if manager.Config.Sharding {
		return nil, errors.New("You should define get table")
	}
	return manager.DefaultTable, nil
}

// Close Database, need override
func (manager *DefaultTableManager) Close() error {
	err := manager.DefaultTable.Close()
	if err != nil {
		return err
	}
	if manager.Config.Sharding {
		return errors.New("You should define close function")
	}
	return nil
}

//Initializate initializate all databases
func (manager *DefaultTableManager) Initializate(config *DatabaseConfig, entity ...interface{}) error {
	//manager.masterDBs = make(map[string]*gorm.DB)
	manager.Config = config
	d := config.DbConnection
	var err error
	if d == nil {
		d, err = gorm.Open(config.DatabaseDialect, config.MasterDatabaseConnectionString)
		if err != nil {
			log.Printf("Cannot Initializate %v", err)
			return err
		}
	}

	manager.DefaultTable = d.Table(config.TableNamePrefix)
	if manager.DefaultTable.Error != nil {
		log.Printf("Cannot Initializate %v", manager.DefaultTable.Error)
		return manager.DefaultTable.Error
	}
	if len(entity) > 0 {
		manager.DefaultTable = manager.DefaultTable.AutoMigrate(entity...)
		if manager.DefaultTable.Error != nil {
			log.Printf("Cannot Initializate %v", manager.DefaultTable.Error)
			return manager.DefaultTable.Error
		}
		return nil
	}
	//manager.defaultTable = d

	return nil
}
