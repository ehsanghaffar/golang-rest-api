package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/likexian/whois"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to The EndPointio!")
	log.Println("Api is up...")
}

func Whois(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	domainName := vars["domainName"]

	result, err := whois.Whois(domainName)

	// parsedInfo := util.WhoisParser(w, result)

	if err == nil {
		log.Println("Get Whois:", domainName)
		fmt.Fprintln(w, result)
	}
}
