# Golang rest API

A rest API writen with Golang programming language.

## EndPoints

- [http://localhost:3000](http://localhost:3000)

return:

```bash
Welcome to The EndPointio!
```

- [http://localhost:3000/api/whois/google.com](http://localhost:3000/api/whois/google.com)

return:

```bash
{
  "domain": {
    "id": "2138514_DOMAIN_COM-VRSN",
    "domain": "google.com",
    "punycode": "google.com",
    "name": "google",
    "extension": "com",
    "whois_server": "whois.markmonitor.com",
    "status": [
      "clientdeleteprohibited",
      "clienttransferprohibited",
      "clientupdateprohibited",
      "serverdeleteprohibited",
      "servertransferprohibited",
      "serverupdateprohibited"
    ],
    "name_servers": [
      "ns1.google.com",
      "ns2.google.com",
      "ns3.google.com",
      "ns4.google.com"
    ],
    "created_date": "1997-09-15T04:00:00Z",
    "updated_date": "2019-09-09T15:39:04Z",
    "expiration_date": "2028-09-14T04:00:00Z"
  },
  "registrar": {
    "id": "292",
    "name": "MarkMonitor Inc.",
    "phone": "+1.2086851750",
    "email": "abusecomplaints@markmonitor.com",
    "referral_url": "http://www.markmonitor.com"
  }
}
```
