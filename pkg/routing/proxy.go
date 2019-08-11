package routing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/jacekk/dead-simple-proxy-server/pkg/compression"
	"github.com/jacekk/dead-simple-proxy-server/pkg/helpers"
	"github.com/jacekk/dead-simple-proxy-server/pkg/storage"
)

func isSlugValid(slug string) error {
	return validation.Validate(slug,
		validation.Required,
		validation.Length(4, 50),
		validation.Match(regexp.MustCompile("^[a-z0-9-]+$")),
	)
}

func initURLProxy(ctx *gin.Context, parsedURL *url.URL) {
	director := func(req *http.Request) {
		req.Host = "" // this is required for some unknown reasons
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.URL.Host = parsedURL.Host
		req.URL.Path = parsedURL.Path
		req.URL.RawQuery = parsedURL.RawQuery
		req.URL.Scheme = parsedURL.Scheme
	}
	responseModifier := func(resp *http.Response) error {
		isCompressed := resp.Uncompressed == false
		body, err := helpers.ReadResponseBody(resp)
		if err != nil {
			return err
		}
		if isCompressed {
			body, err = compression.UngzipBytes(body)
			if err != nil {
				return err
			}
		}
		body = bytes.Replace(body, []byte("__NEEDLE__"), []byte("__REPLACED__"), -1)
		if isCompressed {
			body, err = compression.GzipBytes(body)
			if err != nil {
				return err
			}
		}
		resp.Body = ioutil.NopCloser(bytes.NewReader(body))

		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: responseModifier}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

// GetProxyBySlug - uses slug to serve certain URL.
func GetProxyBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := isSlugValid(slug); err != nil {
		msg := fmt.Sprintf("Given 'slug' is NOT valid --> %s.", err)
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	storedItem, err := storage.GetProxyBySlug(slug)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Slug not found.")
		return
	}

	parsedURL, err := url.Parse(storedItem.URL)
	if err != nil {
		ctx.String(http.StatusBadGateway, "Wrong URL linked with this slug.")
		return
	}

	initURLProxy(ctx, parsedURL)
}
