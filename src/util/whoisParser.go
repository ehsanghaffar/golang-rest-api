package util

import (
	"log"
	"net/http"

	whoisparser "github.com/likexian/whois-parser"
)

func WhoisParser(w http.ResponseWriter, whoisRaw string) whoisparser.WhoisInfo {
	result, err := whoisparser.Parse(whoisRaw)

	if err != nil {
		log.Fatalln("err", err)
	}
	return result
}
