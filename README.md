# Dead Simple Proxy Server

Almost dead simple.

First iteration just allowed to configure urls to serve content from.
The next one to rewrite body content of proxied urls.
The current version allows to precache urls based on configured interval.
In case of errors while refreshing the cache, the last successful response is being served.

### Requirements

* [Go](https://golang.org/doc/install) [ >=1.12 ]
* [pm2](https://pm2.keymetrics.io/) [ >=3.5 ]

### Development

1. `cp dist.env .env` -- and edit if necessary
1. `make run-all &`
1. `http localhost:8080/proxy/example` or `curl localhost:8080/proxy/example`
1. `fg`
1. CTRL + C

### First release

1. `git clone https://github.com/jacekk/dead-simple-proxy-server`
1. `cd dead-simple-proxy-server`
1. `make build`
1. `pm2 start`

### Update

1. `git pull`
1. `make build`
1. `pm2 restart dead-simple-proxy-server`

### License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.
