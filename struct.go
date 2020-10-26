package _struct
type Login struct {
	Num string `json:"num"`
	Password string `json:"password"`
}
type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
type Text struct {
	Num int `json:"num"`
	Text string `json:"text"`
}