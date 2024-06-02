package v1

var (
	Success                = newError(200, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	ErrUsernameAlreadyUse = newError(1001, "The email is already in use.")
)
