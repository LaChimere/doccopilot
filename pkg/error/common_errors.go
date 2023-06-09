package error

var (
	Success         = NewError(0, "success")
	ServerError     = NewError(10000000, "server internal error")
	InvalidParams   = NewError(10000001, "invalid parameters")
	NotFound        = NewError(10000002, "not found")
	Unauthorized    = NewError(10000003, "unauthorized, invalid token")
	TooManyRequests = NewError(10000004, "too many requests")
)
