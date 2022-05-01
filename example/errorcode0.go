package errorcode

// API通用错误码
const (
	// ignore
	commonBaseCode = 0x000000000 + iota

	// 服务器内部错误
	InternalServerError = commonBaseCode + iota
)
