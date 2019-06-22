# Dead Simple Proxy Server

### Requirements

* [Go](https://golang.org/doc/install) [ >=1.12 ]

### Running

1. `cp dist.env .env` -- and edit if necessary
1. `make serve &`
1. `http localhost:8080/proxy/example` or `curl localhost:8080/proxy/example`
1. `fg`
1. CTRL + C
