package api

import (
	"errors"
	"go-apiserver-template/internal/ctx"
	"net/http"
)

const (
	CodeError   = -1
	CodeSuccess = 0

	MsgBadRequest   = "Bad Request"
	MsgUnauthorized = "Unauthorized"
	MsgForbidden    = "Forbidden"
	MsgNotFound     = "Not Found"
)

var (
	ErrBadRequest   = errors.New(MsgBadRequest)
	ErrUnauthorized = errors.New(MsgUnauthorized)
	ErrForbidden    = errors.New(MsgForbidden)
	ErrNotFound     = errors.New(MsgNotFound)
)

// Response is the common response struct
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Success returns a success response
func Success(c *ctx.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Data: data,
	})
}

// Error returns an error response
func Error(c *ctx.Context, err error) {
	c.JSON(http.StatusOK, Response{
		Code: CodeError,
		Msg:  err.Error(),
		Data: nil,
	})
}

// ErrorWithStatus returns an error response with specified http status code
func ErrorWithStatus(c *ctx.Context, status int, err error) {
	c.JSON(status, Response{
		Code: CodeError,
		Msg:  err.Error(),
		Data: nil,
	})
}

// BadRequest returns a bad request response
func BadRequest(c *ctx.Context) {
	ErrorWithStatus(c, http.StatusBadRequest, ErrBadRequest)
}

// Unauthorized returns an unauthorized response
func Unauthorized(c *ctx.Context) {
	ErrorWithStatus(c, http.StatusUnauthorized, ErrUnauthorized)
}

// Forbidden returns a forbidden response
func Forbidden(c *ctx.Context) {
	ErrorWithStatus(c, http.StatusForbidden, ErrForbidden)
}

// NotFound returns a not found response
func NotFound(c *ctx.Context) {
	ErrorWithStatus(c, http.StatusNotFound, ErrNotFound)
}
