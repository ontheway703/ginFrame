package gerror

var (
	// 系统错误
	ServiceUnavailable   = buildError(10001, "service unavailable")
	ServiceInternalError = buildError(10002, "service internal error")

	// 业务错误
	RequestParamsError = buildError(10003, "request params error")
)
