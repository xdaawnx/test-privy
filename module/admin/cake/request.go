package cake

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	CreateCakeReq struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		Rating      float32 `json:"rating"`
		CreatedBy   int     `json:"user_id"`
	}
	UpdateCakeReq struct {
		Id          int     `param:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		Rating      float32 `json:"rating"`
		UpdatedBy   int     `json:"user_id"`
	}
	DetailCakeReq struct {
		Id int `param:"id"`
	}
)

// Validate a NewRole
func (cc CreateCakeReq) Validate() error {
	return v.ValidateStruct(&cc,
		v.Field(&cc.Title, v.Required),
		v.Field(&cc.Description, v.Required),
		v.Field(&cc.Image, v.Required),
		v.Field(&cc.CreatedBy, v.Required),
	)
}
func (uc UpdateCakeReq) Validate() error {
	return v.ValidateStruct(&uc,
		v.Field(&uc.Id, v.Required),
		v.Field(&uc.Title, v.Required),
		v.Field(&uc.Description, v.Required),
		v.Field(&uc.Image, v.Required),
		v.Field(&uc.UpdatedBy, v.Required),
	)
}
func (dc DetailCakeReq) Validate() error {
	return v.ValidateStruct(&dc,
		v.Field(&dc.Id, v.Required),
	)
}
