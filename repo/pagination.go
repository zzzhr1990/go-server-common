package repo

import (
	"log"
	"math"

	"github.com/jinzhu/gorm"
)

//Paginator for page use
type Paginator struct {
	TotalCount int //`json:"total_record"`
	TotalPage  int // `json:"total_page"`
	Page       int //`json:"page"`
	PageSize   int //`json:"prev_page"`
}

// DoPage ip
func (p *Paginator) DoPage(table *gorm.DB, list interface{}) error {
	return Page(table, p, list)
}

//Page so
func Page(table *gorm.DB, p *Paginator, list interface{}) error {

	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 20
	}
	done := make(chan bool, 1)
	go countRecords(table, done, &p.TotalCount)
	offset := (p.Page - 1) * p.PageSize
	err := table.Limit(p.PageSize).Offset(offset).Find(list).Error
	<-done
	if err != nil {
		log.Printf("Query countRecords %v", err)
		return err
	}
	p.TotalPage = int(math.Ceil(float64(p.TotalCount) / float64(p.PageSize)))
	if p.TotalPage < p.Page {
		p.Page = p.TotalPage
	}

	return nil
}

func countRecords(table *gorm.DB, done chan bool, count *int) {
	err := table.Count(count).Error
	if err != nil {
		log.Printf("Query countRecords %v", err)
	}
	done <- true
}
