package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Hello world")

	router := mux.NewRouter()
	router.HandleFunc("/api/session", func(res http.ResponseWriter, req *http.Request) {

	})
	router.HandleFunc("/api/session/id/", func(res http.ResponseWriter, req *http.Request) {

	})
	router.HandleFunc("/api/patch/id/", func(res http.ResponseWriter, req *http.Request) {

	})
	mux := http.NewServeMux()
	mux.Handle("/", router)
	var handler http.Handler = mux
	http.ListenAndServe(":8080", handler)

}
