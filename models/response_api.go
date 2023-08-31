package models

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResponseApi struct {
	c       *fiber.Ctx
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseApi(c *fiber.Ctx) *ResponseApi {
	return &ResponseApi{
		c: c,
	}
}

func (r *ResponseApi) SendJSON(data interface{}) error {
	r.Status = http.StatusOK
	r.Message = "success"
	r.Data = data
	return r.c.JSON(r)
}
func (r *ResponseApi) BadRequestJSON() error {
	r.Status = http.StatusBadRequest
	r.Message = "bad request"
	return r.c.JSON(r)
}
func (r *ResponseApi) BadRequestDataJSON(data interface{}) error {
	r.Status = http.StatusBadRequest
	r.Message = "bad request"
	r.Data = data
	return r.c.JSON(r)
}
func (r *ResponseApi) NotFoundJSON() error {
	r.Status = http.StatusNotFound
	r.Message = "not found"
	return r.c.JSON(r)
}
func (r *ResponseApi) NotFoundWithDataJSON(data interface{}) error {
	r.Status = http.StatusNotFound
	r.Message = "not found"
	r.Data = data
	return r.c.JSON(r)
}

func (r *ResponseApi) CreatedJSON() error {
	r.Status = http.StatusCreated
	r.Message = "success"
	return r.c.JSON(r)
}

// not implement
func (r *ResponseApi) BadCredentials() error {
	r.Status = http.StatusNotFound
	r.Message = "not found"
	return r.c.JSON(r)
}
