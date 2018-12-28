package repo_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
	//. "github.com/smartystreets/goconvey/convey"
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
