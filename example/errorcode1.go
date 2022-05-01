package errorcode

// 用户模块错误码
const (
	// ignore
	userBaseCode = 0x000100000 + iota

	// 用户不存在
	UserNotExist = userBaseCode + iota
)
