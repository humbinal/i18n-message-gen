package errorcode

// 测试模块错误码
const (
	// ignore
	testBaseCode = 0x000300000 + iota

	// BCD 测试十六进制
	BCD = testBaseCode + iota
	//测试十六进制1
	BCD1 = testBaseCode + iota
	//测试十六进制2
	BCD2 = testBaseCode + iota
	//测试十六进制3
	BCD3 = testBaseCode + iota
	//测试十六进制4
	BCD4 = testBaseCode + iota
	//测试十六进制5
	BCD5 = testBaseCode + iota
	//测试十六进制6
	BCD6 = testBaseCode + iota
	//测试十六进制7
	BCD7 = testBaseCode + iota
	//测试十六进制8
	BCD8 = testBaseCode + iota
	//测试十六进制9
	BCD9 = testBaseCode + iota
	//测试十六进制10
	BCD10 = testBaseCode + iota
	//测试十六进制11
	BCD11 = testBaseCode + iota
	//测试十六进制12
	BCD12 = testBaseCode + iota
	//测试十六进制13
	BCD13 = testBaseCode + iota
	//测试十六进制14
	BCD14 = testBaseCode + iota
	//测试十六进制15
	BCD15 = testBaseCode + iota
	//测试十六进制16
	BCD16 = testBaseCode + iota
	//测试十六进制17
	BCD17 = testBaseCode + iota
	//测试十六进制18
	BCD18 = testBaseCode + iota

	//测试占位符,你好{{.who}}
	PING1_TEXT = testBaseCode + iota
)
