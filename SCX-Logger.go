package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func log(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)
	ioutil.WriteFile("SCX_Log.txt", []byte(bodyText), 0644)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/log", log)
	http.ListenAndServe(":8000", mux)
}
