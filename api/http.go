package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		res := []byte("ok")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
	addr := "localhost:9090"
	server := &http.Server{Addr: addr, Handler: router.Handler()}
	log.Printf("Start api server %s", addr)
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Println(err)
			os.Exit(2)
		}
	}
	os.Exit(0)
}
