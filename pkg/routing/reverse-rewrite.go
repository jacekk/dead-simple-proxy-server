package routing

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jacekk/dead-simple-proxy-server/pkg/storage"
)

func isParamValid(param string) error {
	return validation.Validate(param,
		validation.Required,
		validation.Length(4, 50),
		validation.Match(regexp.MustCompile("^[a-z0-9-]+$")),
	)
}

func initURLRewrite(ctx *gin.Context, parsedURL *url.URL) {
	director := func(req *http.Request) {
		req.Host = "" // this is required for some unknown reasons
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.URL.Host = parsedURL.Host
		req.URL.Path = parsedURL.Path
		req.URL.RawQuery = parsedURL.RawQuery
		req.URL.Scheme = parsedURL.Scheme
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func getReverseRewrite(ctx *gin.Context) {
	id := ctx.Param("id")
	suffix := ctx.Param("suffix")
	if err := isParamValid(id); err != nil {
		msg := fmt.Sprintf("Given 'id' is NOT valid --> %s.", err)
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	rewriteKey, err := storage.GetBodyRewriteSource(id)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Requested rewrite key not found.")
		return
	}

	rewriteURL := fmt.Sprintf("%s%s", rewriteKey, suffix)
	parsedURL, err := url.Parse(rewriteURL)
	if err != nil {
		ctx.String(http.StatusBadGateway, "Cannot parse rewrite URL.")
		return
	}

	initURLRewrite(ctx, parsedURL)
}
