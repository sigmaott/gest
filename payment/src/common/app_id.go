package common

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func GetAppId(c echo.Context) string {
	appId := c.Request().Header.Get("X-App-Id")
	if appId == "" {
		return "default-app"
	}
	return appId
}
func GetAcceptLanguage(c echo.Context) string {
	language := c.Request().Header.Get("Accept-Language")
	if language == "" {
		return "en"
	}
	languageSplit := strings.Split(language, "-")
	if len(languageSplit) == 2 {
		return languageSplit[0]
	}
	if len(languageSplit) == 1 {
		return languageSplit[0]
	}
	return ""
}
