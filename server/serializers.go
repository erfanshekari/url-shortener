package server

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func serializeJson(c echo.Context) (*map[string]interface{}, error) {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}

	return &jsonBody, nil
}
