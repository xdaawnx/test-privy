package constant

import (
	"errors"
	"strings"

	"test-privy/core/db"

	"github.com/go-ozzo/ozzo-validation/is"
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

type (
	// Response is used a structure of response
	Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}

	// Search is used to crate a search params
	Search struct {
		ColumnName string `json:"column"`
		Params     string `json:"value"`
	}

	// Params is a struct params in url to search
	Params struct {
		Page      int      `json:"page" query:"page"`
		MaxRows   int      `json:"max_rows" query:"max_rows"`
		OrderBy   string   `json:"order_by" query:"order_by"`
		OrderType string   `json:"order_type" query:"order_type"`
		Search    []Search `json:"search" query:"search"`
	}

	// Deletes is a struct to save a delete list
	Deletes struct {
		Type      string        `json:"type"`
		DeleteIDs []interface{} `json:"delete_ids"`
	}
	// Download is a struct to download
	Download struct {
		Type        string        `json:"type"`
		DownloadIDs []interface{} `json:"download_ids"`
	}

	// Approval is used to request approval
	Approval struct {
		Status string `json:"status"`
	}

	// DownloadFileError is a struct to download error when bulk upload
	DownloadFileError struct {
		Filename string `json:"filename"`
	}
)

// Validate is used to validating a record
func (d Deletes) Validate() error {
	return v.ValidateStruct(&d,
		v.Field(&d.Type, v.In("all", "selected")),
		v.Field(&d.DeleteIDs, v.Required, v.Each(is.Int)),
	)
}

// Validate in download error when bulk upload
func (d DownloadFileError) Validate() error {
	return v.ValidateStruct(&d,
		v.Field(&d.Filename, v.Required),
	)
}

// Validate in download struct
func (d Download) Validate() error {
	return v.ValidateStruct(&d,
		v.Field(&d.Type, v.Required, v.In("all", "selected")),
		v.Field(&d.DownloadIDs, v.Each(is.Int)),
	)
}

// Validate in search struct
func (s Search) Validate() error {
	return v.ValidateStruct(&s,
		v.Field(&s.ColumnName, v.Required, v.Match(AlphUnds)),
	)
}

// Validate in Params struct
func (p Params) Validate(ord ...interface{}) error {
	return v.ValidateStruct(&p,
		v.Field(&p.OrderBy, v.Required, v.In(ord...)),
		v.Field(&p.OrderType, v.Required, v.In("asc", "desc")),
		v.Field(&p.Search),
	)
}

// SetDefault is used to fill default value when it blank
func (p *Params) SetDefault() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.MaxRows == 0 {
		p.MaxRows = 10
	}
	p.OrderType = strings.ToLower(p.OrderType)

}

// Responsebuilder function is for getting response code from database
func (r *Response) Responsebuilder() error {
	e := errors.New("can't get rc data")
	db, err := db.MainDB()
	rc := new(RC)
	err = db.Table(rc.TableName()).Where("code = ?", r.Status).First(&rc).Error
	if gorm.IsRecordNotFoundError(err) {
		r.Status = Internalerror
		r.Message = rc.Message
		return e
	}
	if err != nil {
		r.Status = Internalerror
		r.Message = rc.Message
		return e
	}
	r.Message = rc.Message

	return nil
}
