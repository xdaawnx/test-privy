package cake

import (
	"errors"
	"net/http"

	"test-privy/helper/constant"

	"github.com/jinzhu/copier"

	"github.com/labstack/echo/v4"
)

func GetAll(ctx echo.Context) error {
	p := new(constant.Params)
	resp := new(constant.Response)

	ctx.Bind(p)
	p.SetDefault()
	err := p.Validate(Column...)
	if err != nil {
		resp.Status = constant.Badrequest
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	m := new(ListCakes)
	result, err := m.CakeList(p)

	if err != nil {
		resp.Status = constant.Internalerror
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = constant.Success
	resp.Data = result
	resp.Responsebuilder()

	return ctx.JSON(http.StatusOK, resp)
}

func Find(ctx echo.Context) error {
	response := new(constant.Response)
	req := &DetailCakeReq{}
	ctx.Bind(req)
	err := req.Validate()
	if err != nil {
		response.Status = constant.Badrequest
		response.Errors = err.Error()
		response.Responsebuilder()
		return ctx.JSON(http.StatusBadRequest, response)
	}
	m := new(Cake)
	err = m.CakeDetail(*req)
	if err != nil && m.Id == 0 {
		response.Status = constant.Internalerror
		response.Errors = err.Error()
		response.Responsebuilder()
		return ctx.JSON(http.StatusInternalServerError, response)
	} else if err == nil && m.Id == 0 {
		response.Status = constant.Nodata
		response.Errors = errors.New("data not found")
		response.Responsebuilder()
		return ctx.JSON(http.StatusNotFound, response)
	}
	response.Status = constant.Success
	data := &CakeDetailRes{}
	copier.Copy(&data, &m)
	response.Data = data
	response.Responsebuilder()
	return ctx.JSON(http.StatusOK, response)
}

func Create(ctx echo.Context) error {
	resp := &constant.Response{}
	req := &CreateCakeReq{}
	m := new(Cake)
	ctx.Bind(req)

	err := req.Validate()
	if err != nil {
		resp.Status = constant.Badrequest
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp.Status = constant.Internalerror
	err = m.CakeCreate(req)
	if err != nil {
		resp.Status = constant.Internalerror
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = constant.SuccessCreate
	resp.Responsebuilder()
	return ctx.JSON(http.StatusCreated, resp)
}

func Update(ctx echo.Context) error {
	resp := &constant.Response{}
	req := &UpdateCakeReq{}
	ctx.Bind(req)

	m := new(Cake)
	err := m.CakeUpdate(req)
	if err != nil && err.Error() == "not found" {
		resp.Status = constant.Nodata
		resp.Errors = errors.New("data not found")
		resp.Responsebuilder()
		return ctx.JSON(http.StatusNotFound, resp)
	} else if err != nil {
		resp.Status = constant.Internalerror
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = constant.Success
	resp.Responsebuilder()
	return ctx.JSON(http.StatusOK, resp)
}

func Delete(ctx echo.Context) error {
	resp := new(constant.Response)
	req := &DetailCakeReq{}
	ctx.Bind(req)
	err := req.Validate()
	if err != nil {
		resp.Status = constant.Validationfailed
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	m := new(Cake)
	err = m.CakeDelete(*req)
	if err != nil && err.Error() == "not found" {
		resp.Status = constant.Nodata
		resp.Errors = errors.New("data not found")
		resp.Responsebuilder()
		return ctx.JSON(http.StatusNotFound, resp)
	} else if err != nil {
		resp.Status = constant.Internalerror
		resp.Errors = err.Error()
		resp.Responsebuilder()
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = constant.Success
	resp.Responsebuilder()
	return ctx.JSON(http.StatusOK, resp)
}
