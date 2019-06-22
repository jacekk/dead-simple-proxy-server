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

func isSlugValid(slug string) error {
	return validation.Validate(slug,
		validation.Required,
		validation.Length(4, 50),
		validation.Match(regexp.MustCompile("^[a-z0-9-]+$")),
	)
}

func initURLProxy(parsedURL *url.URL, ctx *gin.Context) {
	req := ctx.Request

	req.URL.Host = parsedURL.Host
	req.URL.Scheme = parsedURL.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = parsedURL.Host

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)
	proxy.ServeHTTP(ctx.Writer, req)
}

// GetProxyBySlug - uses slug to serve certain URL.
func GetProxyBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := isSlugValid(slug); err != nil {
		msg := fmt.Sprintf("Given 'slug' is NOT valid --> %s.", err)
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	storedURL := storage.GetProxyBySlug(slug)
	if storedURL == "" {
		ctx.String(http.StatusBadRequest, "Slug not found.")
		return
	}

	parsedURL, err := url.Parse(storedURL)
	if err != nil {
		ctx.String(http.StatusBadGateway, "Wrong URL linked with this slug.")
		return
	}

	initURLProxy(parsedURL, ctx)
}
