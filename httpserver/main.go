package main

import (
	"os"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	fmt.Println("request:", r.URL.Query())
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		w.Header().Set(k, fmt.Sprintf("%s", v))
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Println("res header: ", w.Header().Clone())
}
