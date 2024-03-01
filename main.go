package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// 404ページを表示する関数を実行
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", gethome)
	e.Logger.Fatal(e.Start(":1323"))
}

func gethome(c echo.Context) error {
	return c.File("web/template/home.html")
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
