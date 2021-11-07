package entities

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/constants"
	"github.com/ntsd/alpha-interface-backend/pkg/utils"
	"go.uber.org/zap"
)

// errorResponse is struct for success json response
type errorResponse struct {
	Status ErrorStatus `json:"status"`
}

// NewErrorResponse is for create error response
func NewErrorResponse(err error) errorResponse {
	errorStatus, ok := err.(*ErrorStatus)
	if !ok {
		return errorResponse{
			Status: ErrorStatus{
				Code:    constants.InternalServerErrorCode,
				Message: constants.InternalServerErrorMessage,
				Err:     err,
			},
		}
	}
	return errorResponse{
		Status: *errorStatus,
	}
}

// Response make a http writer to write json body and http status code
func (r errorResponse) Response(c *gin.Context, httpStatus int) {
	zap.L().Error(r.Status.Error())
	c.AbortWithStatusJSON(httpStatus, r)
}

// Redirect make a http writer to write json body and http status code
func (r errorResponse) Redirect(c *gin.Context, uri string) {
	zap.L().Error(r.Status.Error())
	u, err := url.Parse(uri)
	if err != nil {
		NewErrorResponse(&ErrorStatus{
			Code:    constants.InternalServerErrorCode,
			Message: constants.InternalServerErrorMessage,
			Err:     err,
		}).Response(c, http.StatusInternalServerError)
		return
	}

	err = utils.InterfaceToURLQuery(u, r.Status)
	if err != nil {
		NewErrorResponse(&ErrorStatus{
			Code:    constants.InternalServerErrorCode,
			Message: constants.InternalServerErrorMessage,
			Err:     err,
		}).Response(c, http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, u.String())
}
