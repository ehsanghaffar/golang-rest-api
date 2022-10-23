package util

import (
	"fmt"
	"net/http"

	whoisparser "github.com/likexian/whois-parser"
)

func WhoisParser(w http.ResponseWriter, whoisRaw string) {
	result, err := whoisparser.Parse(whoisRaw)

	if err == nil {
		fmt.Fprintln(w, result.Domain)
		fmt.Fprintln(w, result.Technical)
	}
}
