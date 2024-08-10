package web

import (
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", GeneralHandler)
	http.HandleFunc("/healthz", HelloHandler)
	log.Println("Start Web Server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
