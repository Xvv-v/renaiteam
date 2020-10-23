package main

import (
	"encoding/json"
	"mygo/main/my/mysql"
	_ "mygo/main/my/mysql"
	_ "mygo/main/my/struct"
	_struct "mygo/main/my/struct"
	"net/http"
)
func Login(w http.ResponseWriter,r *http.Request)  {
	var dl _struct.Login
	json.NewDecoder(r.Body).Decode(&dl)//获取前端的json数据\
	mysql.Init()
	bool1:=mysql.Login(dl)
	if bool1==true{
		msg, _ := json.Marshal(_struct.Result{Code: 200, Msg: "验证成功！"})
		w.Write(msg)
	} else{
		msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "验证失败账号或密码错误！"})
			w.WriteHeader(400)
			w.Write(msg)
	}
	mysql.Close()
}
func register (w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var dl _struct.Login
	json.NewDecoder(r.Body).Decode(&dl)//获取前端的json数据\
	bool1:=mysql.Register(dl)
	if bool1==true{
		msg, _ := json.Marshal(_struct.Result{Code: 200, Msg: "注册成功！"})
		w.Write(msg)
	} else{
		msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "注册失败账号已存在！"})
		w.WriteHeader(400)
		w.Write(msg)
	}
	mysql.Close()
}
func addition (w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test _struct.Text
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	bool1:=mysql.Addition(test)
	if bool1==true{
		msg, _ := json.Marshal(_struct.Result{Code: 200, Msg: "添加成功！"})
		w.Write(msg)
	}else{
		msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "添加失败！"})
		w.WriteHeader(400)
		w.Write(msg)
	}
	mysql.Close()
}
func delete(w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test _struct.Text
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	bool1:=mysql.Delete(test)
	if bool1==true{
		msg, _ := json.Marshal(_struct.Result{Code: 200, Msg: "删除成功！"})
		w.Write(msg)
	}else{
		msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "删除失败！"})
		w.WriteHeader(400)
		w.Write(msg)
	}
	mysql.Close()
	
}
func modification(w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test _struct.Text
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	bool1:=mysql.Modification(test)
	if bool1==true{
		msg, _ := json.Marshal(_struct.Result{Code: 200, Msg: "修改成功！"})
		w.Write(msg)
	}else{
		msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "修改失败！"})
		w.WriteHeader(400)
		w.Write(msg)
	}
	mysql.Close()
}
func show(w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test _struct.Text
	data:=mysql.Show()
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	msg, _ := json.Marshal(data)
	w.Header().Set("name","ifbi")
	w.WriteHeader(200)
	w.Write(msg)
}
func main()  {
	http.HandleFunc("/Login",Login)
	http.HandleFunc("/register",register)
	http.HandleFunc("/addition",addition)
	http.HandleFunc("/delete",delete)
	http.HandleFunc("/modification",modification)
	http.HandleFunc("/show",show)
	http.ListenAndServe(":8080",nil)
}