package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func log(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)

	logfile, _ := os.OpenFile("SCX_Log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)

	var now = time.Now()

	logfile.WriteString(now.String() + ": " + bodyText + "\r\n")

	logfile.Close()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/log", log)
	http.ListenAndServe(":8000", mux)
}
