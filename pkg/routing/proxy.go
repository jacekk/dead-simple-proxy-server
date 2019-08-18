package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"

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

func initURLProxy(
	ctx *gin.Context,
	config storage.Item,
	parsedURL *url.URL,
) {
	director := func(req *http.Request) {
		req.Host = "" // this is required for some unknown reasons
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.URL.Host = parsedURL.Host
		req.URL.Path = parsedURL.Path
		req.URL.RawQuery = parsedURL.RawQuery
		req.URL.Scheme = parsedURL.Scheme
	}
	responseModifier := func(resp *http.Response) error {
		isCompressed := helpers.IsResponseCompressed(resp)
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
		for from, to := range config.BodyRewrite {
			body = bytes.Replace(body, []byte(from), []byte(to), -1)
		}
		if isCompressed {
			body, err = compression.GzipBytes(body)
			if err != nil {
				return err
			}
		}
		resp.Body = ioutil.NopCloser(bytes.NewReader(body))
		if !isCompressed {
			resp.ContentLength = int64(len(body))
			resp.Header.Set("Content-Length", strconv.Itoa(len(body)))
		}

		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: responseModifier}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func getProxyBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := isSlugValid(slug); err != nil {
		msg := fmt.Sprintf("Given 'slug' is NOT valid --> %s.", err)
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	storedItem, err := storage.GetProxyConfigBySlug(slug)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Slug not found.")
		return
	}

	parsedURL, err := url.Parse(storedItem.URL)
	if err != nil {
		ctx.String(http.StatusBadGateway, "Wrong URL linked with this slug.")
		return
	}

	if storedItem.IsCacheEnabled {
		readCachedResponse(ctx, storedItem, parsedURL)
		return
	}

	initURLProxy(ctx, storedItem, parsedURL)
}

func readCachedResponse(
	ctx *gin.Context,
	item storage.Item,
	parsedURL *url.URL,
) {
	bodyPath := storage.SlugCachePath(item.ID, "body")
	headersPath := storage.SlugCachePath(item.ID, "headers")

	body, err := ioutil.ReadFile(bodyPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read content from cache.")
		return
	}

	gzipped, err := compression.GzipBytes(body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to gzip cached content.")
		return
	}

	headers, err := ioutil.ReadFile(headersPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read cached headers.")
		return
	}

	var parsedHeaders http.Header
	err = json.Unmarshal(headers, &parsedHeaders)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse headers from cache --> %s.", err)
		ctx.String(http.StatusInternalServerError, msg)
		return
	}

	for key, value := range parsedHeaders {
		ctx.Writer.Header().Set(key, value[0])
	}

	ctx.Writer.Header().Set("Content-Encoding", "gzip")
	ctx.String(http.StatusOK, string(gzipped))
}
