package cake

import (
	"errors"
	"fmt"

	"test-privy/core/db"
	helper "test-privy/helper"
	"test-privy/helper/constant"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source model.go -package mock -destination .\mock\cake.go
type Model interface {
	CakeList(*constant.Params) (*constant.DataPaging, error)
	CakeDetail(DetailCakeReq) error
	CakeCreate(*CreateCakeReq) error
	CakeUpdate(*UpdateCakeReq) error
	CakeDelete(DetailCakeReq) error
}

type ModelRepository struct {
	Model
}

func (m *ListCakes) CakeList(p *constant.Params) (*constant.DataPaging, error) {
	pg := new(constant.DataPaging)
	pgr := new(helper.Pager)
	var count int

	db, err := db.MainDB()
	if err != nil {
		db.Close()
		return nil, errors.New("failed to connect database")
	}
	defer db.Close()

	qs, err := helper.QueryStringify(Column, false, p.Search...)
	if err != nil {
		return nil, errors.New("error filter list")
	}

	pgr.MaxRows = p.MaxRows
	pgr.Page = p.Page
	pg.PageActive = p.Page

	q := db.Model(&Cake{}).
		Select(`
		id, 
		title, 
		description, 
		rating, 
		image, 
		DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') created_at, 
		DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') updated_at
		`).
		Where(qs).
		Order(p.OrderBy + " " + p.OrderType)
	err = q.Count(&count).Error

	if err != nil {
		return pg, err
	}
	if count == 0 {
		return pg, nil
	}

	pgr.Total = count
	pg.Total = count
	pg.Showing = pgr.GetOffset()

	err = q.Offset(pgr.Offset).
		Limit(p.MaxRows).
		Scan(&m.Cakes).Error
	if err != nil {
		return pg, err
	}

	pg.PageList = pgr.Paging()

	res := []CakeListRes{}
	if count == 0 {
		return pg, nil
	}
	copier.Copy(&res, &m.Cakes)
	pg.Rows = res
	fmt.Printf(`%+v`, m.Cakes)

	return pg, nil

}

func (m *Cake) CakeDetail(cake DetailCakeReq) error {
	db, err := db.MainDB()
	if err != nil {
		db.Close()
		return errors.New("failed to connect database")
	}
	defer db.Close()

	err = db.Model(&Cake{}).
		Select(`
		id, 
		title, 
		description, 
		rating, 
		image, 
		created_by, 
		DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, 
		updated_by, 
		DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at
		`).
		Where("id = ?", cake.Id).
		Scan(&m).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (m *Cake) CakeCreate(c *CreateCakeReq) error {
	e := errors.New("error create Cake")

	db, err := db.MainDB()
	if err != nil {
		db.Close()
		return errors.New("failed to connect database")
	}
	defer db.Close()

	m.CreatedAt = helper.GetDateNow()
	copier.Copy(&m, &c)
	if err = db.Create(&m).Error; err != nil {
		return e
	}

	return nil
}

func (m *Cake) CakeUpdate(c *UpdateCakeReq) error {
	e := errors.New("error update cake")
	db, err := db.MainDB()
	if err != nil {
		db.Close()
		return errors.New("failed to connect database")
	}
	defer db.Close()

	err = c.Validate()
	if err != nil {
		return err
	}

	err = db.Model(&Cake{}).
		Where("id = ?", c.Id).
		Find(&m).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("not found")
	}
	copier.Copy(&m, &c)
	if err = db.Model(&Cake{}).Update(&m).Error; err != nil {
		return e
	}

	return nil
}

func (m *Cake) CakeDelete(c DetailCakeReq) error {
	e := errors.New("error delete cake")
	db, err := db.MainDB()
	if err != nil {
		db.Close()
		return errors.New("failed to connect database")
	}

	err = db.Model(&Cake{}).
		Where("id = ?", c.Id).
		Find(&m).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("not found")
	}

	defer db.Close()
	copier.CopyWithOption(&m, &c, copier.Option{IgnoreEmpty: true})
	if err = db.Delete(&m).Error; err != nil {
		return e
	}
	return nil
}
