// nolint:gomnd
package errors

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(message string) *Error {
	return Newf(400, message)
}

// IsBadRequest determines if err is an error which indicates a BadRequest error.
// It supports wrapped errors.
func IsBadRequest(err error) bool {
	return Code(err) == 400
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(message string) *Error {
	return Newf(401, message)
}

// IsUnauthorized determines if err is an error which indicates a Unauthorized error.
// It supports wrapped errors.
func IsUnauthorized(err error) bool {
	return Code(err) == 401
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(message string) *Error {
	return Newf(403, message)
}

// IsForbidden determines if err is an error which indicates a Forbidden error.
// It supports wrapped errors.
func IsForbidden(err error) bool {
	return Code(err) == 403
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(message string) *Error {
	return Newf(404, message)
}

// IsNotFound determines if err is an error which indicates an NotFound error.
// It supports wrapped errors.
func IsNotFound(err error) bool {
	return Code(err) == 404
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(message string) *Error {
	return Newf(409, message)
}

// IsConflict determines if err is an error which indicates a Conflict error.
// It supports wrapped errors.
func IsConflict(err error) bool {
	return Code(err) == 409
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(message string) *Error {
	return Newf(500, message)
}

// IsInternalServer determines if err is an error which indicates an Internal error.
// It supports wrapped errors.
func IsInternalServer(err error) bool {
	return Code(err) == 500
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailable(message string) *Error {
	return Newf(503, message)
}

// IsServiceUnavailable determines if err is an error which indicates a Unavailable error.
// It supports wrapped errors.
func IsServiceUnavailable(err error) bool {
	return Code(err) == 503
}

// GatewayTimeout new GatewayTimeout error that is mapped to a HTTP 504 response.
func GatewayTimeout(message string) *Error {
	return Newf(504, message)
}

// IsGatewayTimeout determines if err is an error which indicates a GatewayTimeout error.
// It supports wrapped errors.
func IsGatewayTimeout(err error) bool {
	return Code(err) == 504
}

// ClientClosed new ClientClosed error that is mapped to a HTTP 499 response.
func ClientClosed(message string) *Error {
	return Newf(499, message)
}

// IsClientClosed determines if err is an error which indicates a IsClientClosed error.
// It supports wrapped errors.
func IsClientClosed(err error) bool {
	return Code(err) == 499
}

func LoginNoAuth() *Error {
	return Newf(107, "用户未登录")
}

func IsLoginNoAuth(err error) bool {
	return Code(err) == 107
}

func LoginErr() *Error {
	return Newf(108, "登录出错")
}

func IsLoginErr(err error) bool {
	return Code(err) == 108
}
