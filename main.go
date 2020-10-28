package main

import (
	"encoding/json"
	"log"
	"mygo/main/my/mysql"
	_ "mygo/main/my/mysql"
	"mygo/main/my/structural_text"
	_ "mygo/main/my/structural_text"
	"net/http"
	"strconv"
)


func Login(w http.ResponseWriter,r *http.Request)  {
	var dl structural_text.Login
	json.NewDecoder(r.Body).Decode(&dl)//获取前端的json数据\
	mysql.Init()
	bool1:=mysql.Login(dl)
	if bool1==true{
		r.Cookie("user");
		data,err:=json.Marshal(structural_text.Result{Code: 200, Msg: "验证成功！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		//msg, _ := json.Marshal(structural_text.Result{Code: 200, Msg: "验证成功！"})
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Write(b)
	} else{
		data,err:=json.Marshal(structural_text.Result{Code: 400, Msg: "验证失败！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		//msg, _ := json.Marshal(structural_text.Result{Code: 400, Msg: "验证失败账号或密码错误！"})
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")

		w.Write(b)
	}
	mysql.Close()
}
func register (w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var dl structural_text.Login
	json.NewDecoder(r.Body).Decode(&dl)//获取前端的json数据\
	bool1:=mysql.Register(dl)
	if bool1==true{
		data,err:=json.Marshal(structural_text.Result{Code: 200, Msg: "验证成功！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Write(b)
	} else{
		data,err:=json.Marshal(structural_text.Result{Code: 400, Msg: "验证失败！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		//msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "验证失败账号或密码错误！"})
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")

		w.Write(b)
	}
	mysql.Close()
}
func addition (w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test structural_text.Text
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	bool1:=mysql.Addition(test)
	if bool1==true{
		data,err:=json.Marshal(structural_text.Result{Code: 200, Msg: "验证成功！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Write(b)
	}else{
		data,err:=json.Marshal(structural_text.Result{Code: 400, Msg: "验证失败！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		//msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "验证失败账号或密码错误！"})
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")

		w.Write(b)
	}
	mysql.Close()
}
func delete(w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test structural_text.Text
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	bool1:=mysql.Delete(test)
	if bool1==true{
		data,err:=json.Marshal(structural_text.Result{Code: 200, Msg: "验证成功！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Write(b)
	}else{
		data,err:=json.Marshal(structural_text.Result{Code: 400, Msg: "验证失败！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		//msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "验证失败账号或密码错误！"})
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")

		w.Write(b)
	}
	mysql.Close()
}
func modification(w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test structural_text.Text
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	bool1:=mysql.Modification(test)
	if bool1==true{
		data,err:=json.Marshal(structural_text.Result{Code: 200, Msg: "验证成功！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Write(b)
	}else{
		data,err:=json.Marshal(structural_text.Result{Code: 400, Msg: "验证失败！"})
		if err!=nil{
			log.Println(err,"测试")
			return
		}
		var a json.RawMessage
		a = data
		b,_:=a.MarshalJSON()
		//msg, _ := json.Marshal(_struct.Result{Code: 400, Msg: "验证失败账号或密码错误！"})
		w.Header().Set("Content-Length",strconv.Itoa(len(b)))
		w.Header().Set("Access-Control-Allow-Origin","*")

		w.Write(b)
	}
	mysql.Close()
}
func show(w http.ResponseWriter,r *http.Request)  {
	mysql.Init()
	var test structural_text.Text
	data:=mysql.Show()
	json.NewDecoder(r.Body).Decode(&test)//获取前端的json数据\
	data1,err:=json.Marshal(data)
	if err!=nil{
		log.Println(err,"测试")
		return
	}
	var a json.RawMessage
	a = data1
	b,_:=a.MarshalJSON()
	w.Header().Set("Content-Length",strconv.Itoa(len(b)))
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Write(b)
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