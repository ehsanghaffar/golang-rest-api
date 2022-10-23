package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to The EndPointio!")
	log.Println("Api is up...")
}

func DomainDetails(w http.ResponseWriter, r *http.Request) {
	// get domain name from client
	vars := mux.Vars(r)
	domainName := vars["domainName"]

	// initial environment and get variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	url := os.Getenv("whoisApi")
	token := os.Getenv("token")

	// get domain whois from external API
	req, _ := http.NewRequest("GET", url+domainName, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// return response
	log.Println("Get whois for:", domainName)
	fmt.Fprintf(w, string(body))
}
