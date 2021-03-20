package main

import (
	"fmt"
	"github.com/AlaaAttya/hatly/http/createsession"
	"github.com/AlaaAttya/hatly/mongo"
	"github.com/AlaaAttya/hatly/session"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting server at port 8080...")

	fmt.Println("connecting to mongodb...")
	mongoClient, err := mongo.ConnectToDB(
		"db", "27017", "", "root", "root")
	if err != nil {
		fmt.Printf("failed to connect to mongo db with error %s", err)
	}

	fmt.Println("connected successfully to mongodb...")

	// setup create session handler
	sessionMongoDBRepo := session.NewMongoDBRepository(mongoClient)
	createSessionHandler := createsession.NewHandler(validator.New(), sessionMongoDBRepo)

	router := mux.NewRouter()
	router.HandleFunc("/session", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("hello session")
		err := createSessionHandler.HandleRequest(req)

		if err != nil {
			fmt.Println("failed to handle create session request with error: " + err.Error())
		}
	}).Methods("POST")

	router.HandleFunc("/api/session/id/", func(res http.ResponseWriter, req *http.Request) {

	})
	router.HandleFunc("/api/patch/id/", func(res http.ResponseWriter, req *http.Request) {

	})
	mux := http.NewServeMux()
	mux.Handle("/", router)
	var handler http.Handler = mux
	http.ListenAndServe(":8080", handler)

}
