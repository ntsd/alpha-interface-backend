package entities

import (
	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/constants"
)

type BaseStatus struct {
	Code    uint   `json:"code"`
	Message string `json:"message,omitempty"`
}

// JSONResponse is interface for json response
type JSONResponse interface {
	Response(c *gin.Context, code int)
}

// response is struct for success json response
type response struct {
	Status BaseStatus  `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// NewResponse is for create success response
func NewResponse(data interface{}) JSONResponse {
	return response{
		Status: BaseStatus{
			Code:    constants.SuccessCode,
			Message: constants.SuccessMessage,
		},
		Data: data,
	}
}

// Response make a http writer to write json body and http status code
func (r response) Response(c *gin.Context, httpStatus int) {
	c.AbortWithStatusJSON(httpStatus, r)
}
