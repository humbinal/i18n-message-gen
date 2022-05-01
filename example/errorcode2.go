package errorcode

// 订单模块错误码
const (
	// ignore
	orderBaseCode = 0x000200000 + iota

	// 订单不存在
	OrderNotExist = orderBaseCode + iota
)
