package main

import (
	"log"
	"net/http"
	db "github.com/tusharhow/go-api/db"
	handlers "github.com/tusharhow/go-api/handlers"
)

func main() {

	log.Println("Starting the application...")
	router := http.NewServeMux()
	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/signup", handlers.Signup)
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to DB")
		log.Fatal(http.ListenAndServe(":8080", router))
	}

}
