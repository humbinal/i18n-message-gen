package errorcode

// API通用错误码
const (
	// ignore
	commonBaseCode = 0x000000000 + iota

	// 服务器内部错误
	InternalServerError = commonBaseCode + iota
)

// 用户模块错误码
const (
	// ignore
	userBaseCode = 0x000100000 + iota

	// 用户不存在
	UserNotExist = userBaseCode + iota
)

// 订单模块错误码
const (
	// ignore
	orderBaseCode = 0x000200000 + iota

	// 订单不存在
	OrderNotExist = orderBaseCode + iota
)

// 测试模块错误码
const (
	// ignore
	testBaseCode = 0x000300000 + iota

	//你好
	BCD = testBaseCode + iota
	//你好
	BCD1 = testBaseCode + iota
	//你好
	BCD2 = testBaseCode + iota
	//你好
	BCD3 = testBaseCode + iota
	//你好
	BCD4 = testBaseCode + iota
	//你好
	BCD5 = testBaseCode + iota
	//你好
	BCD6 = testBaseCode + iota
	//你好
	BCD7 = testBaseCode + iota
	//你好
	BCD8 = testBaseCode + iota
	//你好
	BCD9 = testBaseCode + iota
	//你好
	BCD10 = testBaseCode + iota
	//你好
	BCD11 = testBaseCode + iota
	//你好
	BCD12 = testBaseCode + iota
	//你好
	BCD13 = testBaseCode + iota
	//你好
	BCD14 = testBaseCode + iota
	//你好
	BCD15 = testBaseCode + iota
	//你好
	BCD16 = testBaseCode + iota
	//你好
	BCD17 = testBaseCode + iota
	//你好
	BCD18 = testBaseCode + iota

	//你好,{{.who}}
	PING1_TEXT = testBaseCode + iota

	//我是{{.who}}
	PING2_TEXT = testBaseCode + iota

	//测试
	PING3_TEXT = testBaseCode + iota
)
