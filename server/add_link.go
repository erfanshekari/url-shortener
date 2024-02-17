package server

import (
	"net/http"
	"net/url"

	"github.com/erfanshekari/url-shortener/models/link"
	"github.com/labstack/echo/v4"
)

func addLink(c echo.Context) error {

	data, err := serializeJson(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data.")
	}

	if data != nil {
		data_ := *data
		if value, ok := data_["target"]; ok {
			switch val := value.(type) {
			case string:
				url, err := url.ParseRequestURI(val)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, "target's value must be a valid URL.")
				}
				if url.Scheme != "http" && url.Scheme != "https" {
					return echo.NewHTTPError(http.StatusBadRequest, "Invalid URL schema.")
				}
				linkObj, err := link.FindByTarget(val)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Unable to handle request.")
				}
				if linkObj != nil {
					return c.JSON(http.StatusOK, linkObj.ToJson())
				} else {
					newLink := link.Link{
						Target: val,
						Slug:   genUniqueSlug(10),
					}
					err := newLink.Save()
					if err != nil {
						return echo.NewHTTPError(http.StatusInternalServerError, "Unable to handle request.")
					}
					return c.JSON(http.StatusOK, newLink.ToJson())
				}
			default:
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid value type for \"target\".")
			}
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data.")
		}
	}

	return echo.NewHTTPError(http.StatusBadRequest, "Bad request.")
}
