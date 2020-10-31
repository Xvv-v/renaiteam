package main

import (
	"encoding/json"

	"net/http"
)

func Delu_01(w http.ResponseWriter,r *http.Request)  {
	len:=r.ContentLength
	body:=make([]byte,len)
	r.Body.Read(body)
	dl:= main.Delu{}
	json.Unmarshal(body,&dl)
	err:=dl.Denglu()
	if err!=nil{
		return
	}
	w.WriteHeader(200)
	return
}
