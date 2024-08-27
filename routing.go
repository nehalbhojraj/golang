package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/employees", CreateEmployees).Methods("POST")
	r.HandleFunc("/employees", GetEmployee).Methods("GET")
	r.HandleFunc("/employee/{eid}", GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
