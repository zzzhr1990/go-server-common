package repo_test

import (
	"fmt"
	"testing"
	"time"

	//	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zzzhr1990/go-server-common/repo"
)

type TestOverrideTableManager struct {
	repo.DefaultTableManager
	//repo.TableManager
	repo.BaseRepo
}

func (manager *TestOverrideTableManager) GetTable(entity interface{}, read bool) (*gorm.DB, error) {
	return nil, nil
}

func (manager *TestOverrideTableManager) GetMasterTable(entity interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, dbConnectString)
	return db, err
}

func TestDefaultRepo(t *testing.T) {
	Convey("DBinit Test", t, func() {
		db := setUpDatabase(t)
		So(db.Error, ShouldBeNil)
		So(db.Close(), ShouldBeNil)

		xp := &repo.SingleCurdRepo{BaseRepo: repo.BaseRepo{TableManager: &repo.SingleCurdRepo{}}}
		//BaseRepo.TableManager: &TestOverrideTableManager{}
		dbConfig := &repo.DatabaseConfig{}
		dbConfig.DatabaseDialect = dialect
		dbConfig.MasterDatabaseConnectionString = dbConnectString
		dbConfig.Sharding = false
		dbConfig.TableNamePrefix = "db_test_table"
		So(xp.Initializate(dbConfig, &Model{}), ShouldBeNil)
		Convey("DB override Test", func() {
			table, err := xp.GetTable(false, false)
			So(err, ShouldBeNil)
			So(table, ShouldNotBeNil)
			So(xp.DefaultTable, ShouldNotBeNil)
			for a := 1; a < 51; a++ {
				//fmt.
				user := &Model{
					Name:          fmt.Sprintf("name_%v", a),
					Salt:          fmt.Sprintf("salt_%v", a),
					Password:      fmt.Sprintf("password_%v", a),
					Email:         fmt.Sprintf("mail-%v@test.balabala.com", a),
					CountryCode:   "86",
					Phone:         fmt.Sprintf("138001%v", a),
					CreateTime:    uint64(time.Now().Unix()),
					CreateIP:      "127.0.0.1",
					SpaceUsed:     0,
					SpaceCapacity: 6 * 1024 * 1024 * 1024 * 1024,
					Type:          0,
					Status:        0,
					Version:       1,
				}

				So(table.Create(user).Error, ShouldBeNil)
			}

			var users []Model
			pg := &repo.Paginator{PageSize: 5}
			//c := table.Where("uuid > ?", 30)
			//err =
			So(repo.Page(table.Where("uuid > ?", 30), pg, &users), ShouldBeNil)
			So(pg.TotalCount, ShouldEqual, 20)
			So(len(users), ShouldEqual, 5)
			So(pg.TotalPage, ShouldEqual, 4)

			pg.PageSize = 50
			So(repo.Page(table.Where("uuid > ?", 10), pg, &users), ShouldBeNil)
			So(pg.TotalCount, ShouldEqual, 40)
			So(len(users), ShouldEqual, 40)
			So(pg.TotalPage, ShouldEqual, 1)

			pg.PageSize = 39
			So(repo.Page(table.Where("uuid > ?", 10), pg, &users), ShouldBeNil)
			So(pg.TotalCount, ShouldEqual, 40)
			So(len(users), ShouldEqual, 39)
			So(pg.TotalPage, ShouldEqual, 2)
		})

	})
}

func TestOverrideCommonTest(t *testing.T) {
	Convey("Setup", t, func() {
		db := setUpDatabase(t)
		So(db.Error, ShouldBeNil)
		Convey("DBinit Test", func() {
			xp := &TestOverrideTableManager{BaseRepo: repo.BaseRepo{TableManager: &TestOverrideTableManager{}}}
			//BaseRepo.TableManager: &TestOverrideTableManager{}
			dbConfig := &repo.DatabaseConfig{}
			dbConfig.DatabaseDialect = dialect
			dbConfig.MasterDatabaseConnectionString = dbConnectString
			dbConfig.Sharding = false
			dbConfig.TableNamePrefix = "db_test_table"
			e := xp.Initializate(dbConfig)
			So(e, ShouldBeNil)
			Convey("DB override Test", func() {
				res, err := xp.GetTable(false, false)
				So(err, ShouldBeNil)
				So(res, ShouldBeNil)
			})
		})
	})
}
