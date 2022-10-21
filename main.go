package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to The EndPointio!")
}

func getWhois(w http.ResponseWriter, r *http.Request) {
	// get domain from query
	vars := mux.Vars(r)
	domainName := vars["domainName"]
	var token = os.Getenv("token")
	var url = os.Getenv("whoisApi")
	// request Whois api
	req, _ := http.NewRequest("GET", url+domainName, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)

	// get Whois route
	router.HandleFunc("/api/whois/{domainName}", getWhois).Methods("GET")

	err := http.ListenAndServe(":3300", router)
	if err != nil {
		fmt.Println(err)
	}
}
