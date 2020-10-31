package main

import (
	"io/ioutil"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/index", index) // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}