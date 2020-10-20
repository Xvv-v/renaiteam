package main

import (
	"encoding/json"
	"mygo/main/my_object/struct"
	"net/http"
)

//noinspection ALL
func Delu_01(w http.ResponseWriter,r *http.Request)  {
	len:=r.ContentLength
	body:=make([]byte,len)
	r.Body.Read(body)
	dl:= _struct.Delu{}
	json.Unmarshal(body,&dl)
	err:=dl.Denglu()
	if err!=nil{
		return
	}
	w.WriteHeader(200)
	return
}
func main()  {
	   http.HandleFunc("/delu_01",Delu_01)
	   http.ListenAndServe(":8080",nil)
}