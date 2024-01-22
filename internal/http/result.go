package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Data: data,
	})
}

// Error returns an error response
func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{
		Code: CodeError,
		Msg:  err.Error(),
		Data: nil,
	})
}

// ErrorWithStatus returns an error response with specified http status code
func ErrorWithStatus(c *gin.Context, status int, err error) {
	c.JSON(status, Response{
		Code: CodeError,
		Msg:  err.Error(),
		Data: nil,
	})
}

// BadRequest returns a bad request response
func BadRequest(c *gin.Context) {
	ErrorWithStatus(c, http.StatusBadRequest, ErrBadRequest)
}

// Unauthorized returns an unauthorized response
func Unauthorized(c *gin.Context) {
	ErrorWithStatus(c, http.StatusUnauthorized, ErrUnauthorized)
}

// Forbidden returns a forbidden response
func Forbidden(c *gin.Context) {
	ErrorWithStatus(c, http.StatusForbidden, ErrForbidden)
}

// NotFound returns a not found response
func NotFound(c *gin.Context) {
	ErrorWithStatus(c, http.StatusNotFound, ErrNotFound)
}
