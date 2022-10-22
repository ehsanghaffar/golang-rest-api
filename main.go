package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to The EndPointio!")
}

var setupServer = func(appPort string) {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)

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
