package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/ntsd/alpha-interface-backend/pkg/constants"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

// decodeAndValidate is function to decode json body and validate
func (h handlers) decodeAndValidate(body io.ReadCloser, s interface{}) error {
	err := h.decodeBody(body, s)
	if err != nil {
		return err
	}

	err = h.validateStruct(s)
	if err != nil {
		return err
	}

	return nil
}

func (h handlers) decodeBody(body io.ReadCloser, s interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(s)
	if err != nil {
		baseError := &entities.ErrorStatus{
			Code:    constants.BadJSONRequestErrorCode,
			Message: constants.BadJSONRequestErrorMessage,
			Err:     err,
		}
		if strings.HasPrefix(err.Error(), "json: unknown field") {
			baseError.Details = append(baseError.Details, err.Error())
		}
		return baseError
	}
	return nil
}

func (h handlers) validateStruct(s interface{}) error {
	err := h.Validator.Struct(s)
	if err != nil {
		baseError := &entities.ErrorStatus{
			Code:    constants.InvalidParameterErrorCode,
			Message: constants.InvalidParameterErrorMessage,
			Err:     err,
		}
		switch err.(type) {
		case validator.ValidationErrors:
			for _, e := range err.(validator.ValidationErrors) {
				baseError.Details = append(baseError.Details, fmtValidationError(e, s))
			}
		}
		return baseError
	}
	return nil
}

func fmtValidationError(e validator.FieldError, s interface{}) string {
	field, ok := reflect.TypeOf(s).Elem().FieldByName(e.Field())
	if ok {
		switch e.Tag() {
		case "required":
			return fmt.Sprintf(constants.RequireParameterErrorDetail, field.Tag.Get("json"))
		default:
			return fmt.Sprintf(constants.InvalidParameterErrorDetail, field.Tag.Get("json"))
		}
	}
	return fmt.Sprintf("unknown json field %s", e.Field())
}

func responseInternalError(c *gin.Context, err error) {
	entities.NewErrorResponse(&entities.ErrorStatus{
		Code:    constants.InternalServerErrorCode,
		Message: constants.InternalServerErrorMessage,
		Err:     err,
	}).Response(c, http.StatusInternalServerError)
}

func responseUnauthorizeError(c *gin.Context, err error) {
	entities.NewErrorResponse(&entities.ErrorStatus{
		Code:    constants.UnauthorizedErrorCode,
		Message: constants.UnauthorizedErrorMessage,
		Err:     err,
	}).Response(c, http.StatusUnauthorized)
}

func responseBadRequestError(c *gin.Context, err error) {
	entities.NewErrorResponse(&entities.ErrorStatus{
		Code:    constants.InvalidParameterErrorCode,
		Message: constants.InvalidParameterErrorMessage,
		Err:     err,
	}).Response(c, http.StatusBadRequest)
}
