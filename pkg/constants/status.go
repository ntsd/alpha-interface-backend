package constants

// Constants for error code, details, and message
const (
	SuccessCode    = 200
	SuccessMessage = "OK"

	InvalidParameterErrorCode    = 400
	InvalidParameterErrorMessage = "Invalid parameter(s)"
	InvalidParameterErrorDetail  = "%s is invalid"
	RequireParameterErrorDetail  = "%s is required"
	NotExistsOrDeletedDetail     = "%s is not exists or deleted"

	UnauthorizedErrorCode    = 400
	UnauthorizedErrorMessage = "unauthorized"

	BadJSONRequestErrorCode    = 402
	BadJSONRequestErrorMessage = "Request body contains badly-formed JSON"

	NotFoundErrorCode    = 404
	NotFoundErrorMessage = "Not Found"

	InternalServerErrorCode    = 500
	InternalServerErrorMessage = "Internal Server Error"
)
