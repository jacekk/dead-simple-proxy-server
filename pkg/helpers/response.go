package helpers

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// IsResponseCompressed -
func IsResponseCompressed(resp *http.Response) bool {
	return strings.Contains(resp.Header.Get("Content-Encoding"), "gzip")
}

// ReadResponseBody -
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
