package errors

var (
	Success       = NewError(200, "Success")
	BadParameters = NewError(400, "Bad Parameters")
	Unauthorized  = NewError(401, "Unauthorized")
	InternalError = NewError(500, "Internal Error")
)
