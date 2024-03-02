package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/static", "web/static")
	// 404ページを表示する関数を実行
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/home", getHome)
	e.GET("/calendar", getCalendar)
	e.GET("/entrance", getEntrance)
	e.GET("/serviceInfo", getServiceInfo)
	e.GET("/latestRelease", getLatestRelease)
	e.GET("/termsOfService", getTermsOfService)
	e.GET("/inquiry", getInquiry)
	e.Logger.Fatal(e.Start(":1323"))
}

func getHome(c echo.Context) error {
	return c.File("web/template/home.html")
}

func getCalendar(c echo.Context) error {
	return c.File("web/template/calendar.html")
}

func getEntrance(c echo.Context) error {
	return c.File("web/template/entrance.html")
}

func getServiceInfo(c echo.Context) error {
	return c.File("web/template/serviceInfo.html")
}

func getLatestRelease(c echo.Context) error {
	return c.File("web/template/latestRelease.html")
}

func getTermsOfService(c echo.Context) error {
	return c.File("web/template/termsOfService.html")
}

func getInquiry(c echo.Context) error {
	return c.File("web/template/inquiry.html")
}

// 404ページを表示する関数
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	if code == http.StatusNotFound {
		// 404エラーの場合は404.htmlを返す
		c.File("404.html")
		return
	}

	c.JSON(code, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	})
}
