package model

type ParamModel struct {
	Ap       int    `json:"ap"`        //WebApiServer启动端口
	TestAddr string `json:"test_addr"` //测试延迟地址
}
