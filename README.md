# Dead Simple Proxy Server

Almost dead simple.

First iteration just allowed to configure urls to redirect to.
The next to rewrite body content of proxied urls.
The current version allows to precache urls based on configured interval.
In case of errors while refreshing the cache, the last succesful response is being served.

### Requirements

* [Go](https://golang.org/doc/install) [ >=1.12 ]

### Running

1. `cp dist.env .env` -- and edit if necessary
1. `make serve &`
1. `http localhost:8080/proxy/example` or `curl localhost:8080/proxy/example`
1. `fg`
1. CTRL + C
