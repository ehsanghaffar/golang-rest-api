package main

import (
	"fmt"
	"go-api/src/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var setupServer = func(appPort string) {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomePage)

	router.HandleFunc("/api/whois/{domainName}", routes.DomainDetails).Methods("GET")

	err := http.ListenAndServe(":"+appPort, router)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	appPort := os.Getenv("appPort")
	if appPort == "" {
		appPort = "3100"
	}

	setupServer(appPort)

}
