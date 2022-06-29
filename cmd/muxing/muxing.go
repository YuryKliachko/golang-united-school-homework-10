package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "I got message:\n", string(body))
	}).Methods("POST")

	router.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		a, _ := strconv.Atoi(r.Header.Get("a"))
		b, _ := strconv.Atoi(r.Header.Get("b"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "\"a+b\": ", "\"", a+b, "\"")
	}).Methods("GET")

	router.HandleFunc("/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello,", vars["name"])
	}).Methods("GET")

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods("GET")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
