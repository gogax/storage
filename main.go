package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"test/handler"
	"test/repository"
)

func main() {
	db := repository.GetInstance()

	db.Set("1", "hi", 10) //example
	db.Set("2", "hi1", 43) //example

	router := mux.NewRouter()
	router.Handle("/get", handler.Get(db))
	router.Handle("/set", handler.Set(db))
	router.Handle("/delete", handler.Delete(db))
	http.Handle("/", router)

	go db.StartClean(1) //start loop clean

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}