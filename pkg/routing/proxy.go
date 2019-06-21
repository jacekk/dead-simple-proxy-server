package routing

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func isSlugValid(slug string) error {
	return validation.Validate(slug,
		validation.Required,
		validation.Length(4, 50),
		validation.Match(regexp.MustCompile("^[a-z0-9-]+$")),
	)
}

// GetProxyBySlug - uses slug to serve certain URL.
func GetProxyBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := isSlugValid(slug); err != nil {
		msg := fmt.Sprintf("Given 'slug' is NOT valid --> %s.", err)
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	ctx.String(http.StatusOK, slug)
}
