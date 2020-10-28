package structural_text//结构体，文本·

type Login struct {//登陆结构体
	Accout string `json:"accout"`//账号
	Password string `json:"password"`//密码
}
type Result struct {//返回结构结构体
	Code int `json:"code"`//状态码
	Msg string `json:"msg"`//信息
}
type Text struct {//记事本文本文本结构体
	Num string `json:"num"`//编号
	Text string `json:"text"`//文本
}