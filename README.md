# Golang rest API

A rest API writen with Golang programming language.

## Getting Start

At this time, you have a RESTful API server running at http://127.0.0.1:3000. It provides the following endpoints:

* ```GET /api/whois/{domainName}``` returns the detailed information of a domain

### How To Use

```bash
git clone https://github.com/ehsanghaffar/golang-rest-api.git
cd golang-rest-api

go run ./main.go
```

#### Tools

* Routing [Mux](https://github.com/gorilla/mux)
* Whois query [Whois](github.com/likexian/whois)
* Whois information parser [WhoisParser](github.com/likexian/whois-parser)

#### TODO

- [ ] API Documentation - Lot of potential to improve.
- [ ] Add error handling.
