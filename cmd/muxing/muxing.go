package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GolangUnited/helloweb/internal/webserver"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	handler := webserver.NewHandler()
	handler.Register(router)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}

	Start(host, port)

}
