package repo

import (
	"errors"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

// MultiTableManager is a concurrent map with amortized-constant-time loads, stores, and deletes.
// It is safe for multiple goroutines to call a MultiTableManager's methods concurrently.
//
// The zero MultiTableManager is valid and empty.
//
// A MultiTableManager must not be copied after first use.
type MultiTableManager struct {
	tables *sync.Map
	baseDB *gorm.DB
	lock   *sync.RWMutex
	entity []interface{}
	Config *DatabaseConfig
}

//GetTable Imp get default table
func (manager *MultiTableManager) GetTable(entity interface{}, read bool) (*gorm.DB, error) {
	return nil, errors.New("You should overrride GetTable function")
}

//GetMultiTable is get table
func (manager *MultiTableManager) GetMultiTable(tableName string) (*gorm.DB, error) {
	data, ok := manager.tables.Load(tableName)
	if ok {
		db, suc := data.(*gorm.DB)
		if !suc {
			return nil, errors.New("Convert DB instance Error")
		}
		return db, nil
	}
	// miss, lock to prevent concurrent.
	manager.lock.Lock()
	data, ok = manager.tables.Load(tableName)
	// dounle checking
	if ok {
		manager.lock.Unlock()
		db, suc := data.(*gorm.DB)
		if !suc {
			return nil, errors.New("Convert DB instance Error")
		}
		return db, nil
	}
	// create new & switch
	// note: if failed to auto migrate, will just log an error.
	table := manager.baseDB.Table(tableName)
	if table.Error != nil {
		log.Printf("Switch to table %v err, %v", tableName, table.Error)
		manager.lock.Unlock()
		return nil, table.Error
	}
	if len(manager.entity) > 0 {
		tb := table.AutoMigrate(manager.entity...)
		if tb.Error != nil {
			log.Printf("AutoMigrate table %v err, %v", tableName, tb.Error)
		}
		manager.lock.Unlock()
		return nil, tb.Error
	}
	manager.tables.Store(tableName, table)
	manager.lock.Unlock()
	return table, nil
}

//Close is get table when close
func (manager *MultiTableManager) Close() error {
	// Close main table.
	err := manager.baseDB.Close()
	if err != nil {
		log.Printf("Close Table error %v", err)
	}
	manager.tables.Range(closeTables)
	return nil
}

func closeTables(key interface{}, value interface{}) bool {
	str, ok := key.(string)
	if ok {
		val, suc := value.(*gorm.DB)
		if suc {
			err := val.Close()
			if err != nil {
				log.Printf("Close Table %v error %v", str, err)
			}
		}
	}
	return true
}

//Initializate Initializate new
func (manager *MultiTableManager) Initializate(config *DatabaseConfig, entity ...interface{}) error {
	//manager.Config = config
	manager.Config = config
	manager.tables = &sync.Map{}
	manager.lock = new(sync.RWMutex)
	manager.entity = entity
	d := config.DbConnection
	var err error
	if d == nil {
		d, err = gorm.Open(config.DatabaseDialect, config.MasterDatabaseConnectionString)
		if err != nil {
			log.Printf("Cannot Initializate %v", err)
			return err
		}
	}
	manager.baseDB = d
	return nil
}

/*
// CreateNewMultiTableManager instance
func CreateNewMultiTableManager(baseDB *gorm.DB, entity ...interface{}) *MultiTableManager {

	return &MultiTableManager{
		tables: &sync.Map{},
		lock:   new(sync.RWMutex),
		baseDB: baseDB,
		entity: entity,
	}
}
*/

/*
var (
	tableMap = &syncmap.Map{}
)
*/
