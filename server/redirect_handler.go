package server

import (
	"net/http"

	"github.com/erfanshekari/url-shortener/models/link"
	"github.com/labstack/echo/v4"
)

func redirectHandler(c echo.Context) error {
	slug := c.Param("slug")
	l, err := link.FindBySlug("l" + slug)
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}
	if l == nil {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}
	return c.Redirect(http.StatusPermanentRedirect, l.Target)
}
