package repo_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql

	//. "github.com/smartystreets/goconvey/convey"
	"github.com/zzzhr1990/go-server-common/repo"
)

var (
	dbConnectString = "cdb_outerroot:papayadedadingding123*@tcp(591145eb488e5.gz.cdb.myqcloud.com:4513)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialect         = "mysql"
	//db              *gorm.DB
)

const (
	commonTestTableName = "common_test_table"
	dbName              = "test"
)

//Model Used for test
type Model struct {
	UUID     uint64 `gorm:"primary_key;" qingzhen_multi_table:"true;"`
	Name     string `gorm:"type:varchar(128);"`
	Salt     string `gorm:"type:varchar(32);"`
	Password string `gorm:"type:varchar(128);"`
	Email    string `gorm:"type:varchar(128);"`

	CountryCode string `gorm:"type:varchar(32);index:phone_code_idx"`
	Phone       string `gorm:"type:varchar(128);index:phone_code_idx"`
	CreateTime  uint64
	CreateIP    string `gorm:"type:varchar(128);"`
	// Seems SSID no need

	// "ssid": {},
	Icon          string `gorm:"type:varchar(128);"`
	SpaceUsed     uint64
	SpaceCapacity uint64
	Type          uint `gorm:"primary_key;type:int(9);"`
	Status        uint `gorm:"primary_key;type:int(9);"`
	Version       uint `gorm:"primary_key;type:int(6);"`
}

func cleanUpDatabase(t *testing.T) {
	err := setUpDatabase(t).Close()
	if err != nil {
		t.Error(err)
	}
}

func setUpDatabase(t *testing.T) *gorm.DB {
	d, err := gorm.Open(dialect, dbConnectString)
	if err != nil {
		t.Error(err)
	}
	rows, err := d.Raw("SELECT CONCAT('drop table ',table_name,';') as result FROM `information_schema`.`TABLES` WHERE table_schema='" + dbName + "'").Rows()
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		a := ""
		rows.Scan(&a)
		ee := d.Exec(a).Error
		if ee != nil {
			t.Error(ee)
		}
	}
	return d
}

func createDefaultConfig() *repo.DatabaseConfig {
	dbConfig := &repo.DatabaseConfig{}
	dbConfig.DatabaseDialect = dialect
	dbConfig.MasterDatabaseConnectionString = dbConnectString
	dbConfig.Sharding = false
	dbConfig.TableNamePrefix = "db_test_table"
	return dbConfig
}
