package helper

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"test-privy/core/db"
	"test-privy/helper/constant"

	"github.com/jinzhu/gorm"
)

// Pager is used to get pager data
type Pager struct {
	Total   int
	MaxRows int
	Page    int
	Offset  int
}

// Total for counting data
type Total struct {
	Count uint `gorm:"column:count"`
}

// Paging is used to create a page data list
func (p *Pager) Paging() (pages []int) {
	count := p.Total
	LIMIT := p.MaxRows
	page := p.Page
	allData := math.Ceil(float64(count) / float64(LIMIT))
	iter := int(allData)
	if allData > 5 {
		iter = 5
		inan := []int{1, 2, 3}
		anan := []int{0, 1, 2}
		if in_array(page, inan) {
			return []int{1, 2, 3, 4, 5}
		} else if in_array(count-page, anan) {
			pages = []int{count - 4, count - 3, count - 2, count - 1, count}
		} else {
			if page == int(allData) {
				pages = []int{page - 2, page - 1, page}
			} else if page+1 == int(allData) {
				pages = []int{page - 2, page - 1, page, page + 1}
			} else {
				pages = []int{page - 2, page - 1, page, page + 1, page + 2}
			}
		}
	} else {
		for i := 1; i <= iter; i++ {
			pages = append(pages, i)
		}
	}

	return
}

func in_array(val int, array []int) (exists bool) {
	exists = false
	for _, v := range array {
		if val == v {
			exists = true
			return
		}
	}
	return
}

// GetOffset is used to count a page to data
func (p *Pager) GetOffset() string {
	from := p.MaxRows * (p.Page - 1)
	to := from + p.MaxRows
	p.Offset = from
	if (from == 0) && (p.Total != 0) {
		from = 1
	}
	if to > p.Total {
		to = p.Total
	}
	return fmt.Sprintf("%d to %d of %d", from, to, p.Total)
}

// Count is return a number of rows
func Count(main bool, tbl, col, where string) (int, error) {
	var dbs *gorm.DB
	var err error

	type x struct {
		Count int `gorm:"column:count"`
	}

	if main {
		dbs, err = db.MainDB()
	}

	if err != nil {
		return 0, errors.New("can't get count")
	}

	defer dbs.Close()
	c := new(x)
	err = dbs.Table(tbl).
		Select("COUNT(" + col + ") as count").
		Where(where).
		Scan(&c).Error
	if err != nil {
		return 0, errors.New("can't get count")
	}
	return c.Count, nil
}

// Count is return a number of rows
func CountWithJoin(main bool, tbl, col, where string, join string) (int, error) {
	var dbs *gorm.DB
	var err error

	type x struct {
		Count int `gorm:"column:count"`
	}

	if main {
		dbs, err = db.MainDB()
	}

	if err != nil {
		return 0, errors.New("can't get count")
	}

	defer dbs.Close()
	c := new(x)
	err = dbs.Table(tbl).
		Select("COUNT(" + col + ") as count").
		Where(where).
		Joins(join).
		Scan(&c).Error
	if err != nil {
		return 0, errors.New("can't get count")
	}
	return c.Count, nil
}

// CountWithGroupBy is used to count data with group by
func CountWithGroupBy(main bool, tbl, col, where, group string) (int, error) {
	var dbs *gorm.DB
	var err error

	type x struct {
		Count int `gorm:"column:count"`
	}

	if main {
		dbs, err = db.MainDB()
	}

	if err != nil {
		return 0, errors.New("can't get count")
	}

	defer dbs.Close()
	c := new(x)

	dbs.Raw("SELECT count(b." + col + ") as count from (SELECT id FROM " + tbl + " WHERE " + where +
		" group by " + group + ") as b").Scan(&c)

	return c.Count, nil
}

// QueryStringify is used to make a quey to string
func QueryStringify(col []interface{}, softdelete bool, c ...constant.Search) (string, error) {
	var s string
	var cs []constant.Search
	var len int

	for _, x := range c {
		if x.Params != "" {
			len++
			cs = append(cs, x)
		}
	}

	if len == 0 {
		if softdelete {
			return "deleted_at = '0000-00-00 00:00:00'", nil
		}
		return s, nil
	}

	for i, n := range cs {
		if n.ColumnName == "assigned_group" {
			slc := strings.Split(n.Params, ",")
			var c bool
			for _, as := range slc {
				if as != "" {
					if c {
						s += " AND "
					}
					s += n.ColumnName + ` LIKE '%{"cg_id":` + as + `}%'`
					c = true
				}
			}
		} else {
			ex, _ := InArray(n.ColumnName, col)
			if !ex {
				return "", errors.New("column " + n.ColumnName + " doesn't exist")
			}
			exist := strings.Contains(n.ColumnName, "date")
			if exist {
				s += "date(" + n.ColumnName + ")=" + "'" + n.Params + "'"
			} else {
				s += n.ColumnName + " LIKE '%" + n.Params + "%'"
			}
		}
		if i+1 < len {
			s += " AND "
		}
	}

	if softdelete {
		s += " AND deleted_at = '0000-00-00 00:00:00' "
	}
	return s, nil
}

func QueryToSp(col []interface{}, softdelete bool, c ...constant.Search) (string, error) {
	var s string
	var cs []constant.Search
	var len int

	for _, x := range c {
		if x.Params != "" {
			len++
			cs = append(cs, x)
		}
	}

	if len == 0 {
		if softdelete {
			return "deleted_at = '0000-00-00 00:00:00'", nil
		}
		return s, nil
	}

	for i, n := range cs {
		if n.ColumnName == "assigned_group" {
			slc := strings.Split(n.Params, ",")
			var c bool
			for _, as := range slc {
				if as != "" {
					if c {
						s += " AND "
					}
					s += n.ColumnName + ` LIKE '%{"cg_id":` + as + `}%'`
					c = true
				}
			}
		} else {
			ex, _ := InArray(n.ColumnName, col)
			if !ex {
				return "", errors.New("column " + n.ColumnName + " doesn't exist")
			}
			exist := strings.Contains(n.ColumnName, "date")
			if exist {
				s += "date(`" + n.ColumnName + "`)=" + "\"" + n.Params + "\""
			} else {
				s += "`" + n.ColumnName + "`" + " LIKE \"%" + n.Params + "%\""
			}
		}
		if i+1 < len {
			s += " AND "
		}
	}

	if softdelete {
		s += " AND deleted_at = '0000-00-00 00:00:00' "
	}
	return s, nil
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
