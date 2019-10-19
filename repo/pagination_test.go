package repo_test

/*
import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/zzzhr1990/go-server-common/repo"
)

func TestPagination(t *testing.T) {
	cleanUpDatabase(t)
	Convey("Pagination Test", t, func() {
		config := createDefaultConfig()
		So(config, ShouldNotBeNil)
		testRepo, err := repo.CreateNew(config, &Model{})
		So(err, ShouldBeNil)
		So(testRepo, ShouldNotBeNil)
		table, err := testRepo.GetTable(false, false)
		So(err, ShouldBeNil)
		So(table, ShouldNotBeNil)
		for a := 1; a < 51; a++ {
			//fmt.
			user := &Model{
				Name:          fmt.Sprintf("name_%v", a),
				Salt:          fmt.Sprintf("salt_%v", a),
				Password:      fmt.Sprintf("password_%v", a),
				Email:         fmt.Sprintf("mail-%v@test.balabala.com", a),
				CountryCode:   "86",
				Phone:         fmt.Sprintf("138001%v", a),
				CreateTime:    time.Now().Unix(),
				CreateIP:      "127.0.0.1",
				SpaceUsed:     0,
				SpaceCapacity: 6 * 1024 * 1024 * 1024 * 1024,
				Type:          0,
				Status:        0,
				Version:       1,
			}

			So(table.Create(user).Error, ShouldBeNil)
		}
		var users []*Model
		pg := &repo.Paginator{PageSize: 5}
		So(pg.DoPage(table.Where("uuid > ?", 30), &users, []string{}), ShouldBeNil)
		So(pg.TotalCount, ShouldEqual, 20)
		So(len(users), ShouldEqual, 5)
		So(pg.TotalPage, ShouldEqual, 4)

		So(users[0].Name, ShouldEqual, "name_31")
		So(users[1].Name, ShouldEqual, "name_32")

		pg.PageSize = 50
		So(repo.Page(table.Where("uuid > ?", 10), pg, &users, []string{}), ShouldBeNil)
		So(pg.TotalCount, ShouldEqual, 40)
		So(len(users), ShouldEqual, 40)
		So(pg.TotalPage, ShouldEqual, 1)
	})
}
*/
