package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func response(path, message string) (int, string) {
	switch path {
	case "/":
		return 200, message + "\n"
	case "/error":
		return 500, "controlled error\n"
	default:
		return 404, "not found\n"
	}
}

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}
	message, exists := os.LookupEnv("SERVER_MESSAGE")
	if !exists {
		message = "Hello, production!"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status, body := response(r.URL.Path, message)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(status)
		fmt.Fprint(w, body)
		fmt.Printf("%s %s %d\n", r.Method, r.URL.Path, status)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
