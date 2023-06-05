package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseApi struct {
	c       echo.Context
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseApi(c echo.Context) *ResponseApi {
	return &ResponseApi{
		c: c,
	}
}

func (r *ResponseApi) SendJSON(data interface{}) error {
	r.Status = http.StatusOK
	r.Message = "success"
	r.Data = data
	return r.c.JSON(r.Status, r)
}
func (r *ResponseApi) BadRequestJSON() error {
	r.Status = http.StatusBadRequest
	r.Message = "bad request"
	return r.c.JSON(r.Status, r)
}
func (r *ResponseApi) NotFoundJSON() error {
	r.Status = http.StatusNotFound
	r.Message = "not found"
	return r.c.JSON(r.Status, r)
}

// not implement
func (r *ResponseApi) BadCredentials() error {
	r.Status = http.StatusNotFound
	r.Message = "not found"
	return r.c.JSON(http.StatusNotFound, r)
}
