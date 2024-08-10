package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 将request中的header写入response header
	for i := range r.Header {
		for _, str := range r.Header[i] {
			w.Header().Set(i, str)
		}
	}

	// 2. 读取version配置
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Set("VERSION", version)
	}

	// 3.记录client IP 和http返回码
	clientIP := r.RemoteAddr
	log.Printf("ClientIP:%s, HTTP Status Code:200", clientIP)
	w.WriteHeader(http.StatusOK)
	// 如果放在writeHeader前面执行Fprintln会导致自动发送200，之后又执行writeHeader，因此会报错http: superfluous response.WriteHeader call from
	fmt.Fprintln(w, "Hello, this is the response body!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
