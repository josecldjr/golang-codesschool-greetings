package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "8000"

	router := mux.NewRouter()

	router.HandleFunc("/", Hello)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
	}

	fmt.Printf("Server started at port %v\n", port)
	fmt.Print(srv.ListenAndServe())
}

// Hello route
func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, greeting("Code Education Rocks!"))
}

func greeting(message string) string {
	return "<b>" + message + "</b>"
}
