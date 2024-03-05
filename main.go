package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Basic認証
func checkAuth(r *http.Request) bool {
	// 認証情報取得
	clientID, clientSecret, ok := r.BasicAuth()
	if ok == false {
		return false
	}
	return clientID == "client_id" && clientSecret == "client_secret"
}

func main() {

	http.HandleFunc("/basic_auth",
		func(w http.ResponseWriter, r *http.Request) {
			// 認証失敗時
			if checkAuth(r) == false {
				w.Header().Add("WWW-Authenticate", `Basic realm="SECRET AREA"`)
				w.WriteHeader(http.StatusUnauthorized) // 401
				http.Error(w, "Unauthorized", 401)
				return
			}
			_, err := fmt.Fprintf(w, "Successful Basic Authentication   \n")
			if err != nil {
				log.Fatal(err)
			}
		},
	)
	log.Println(http.ListenAndServe(":1323", nil))

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

func client() {
	// クライアント生成
	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	// リクエスト定義
	req, err := http.NewRequest("GET", "http://localhost:1323/basic_auth", nil)
	if err != nil {
		log.Fatal(err)
	}
	// 認証情報をセット
	req.SetBasicAuth("client_id", "client_secret")
	// リクエスト実行
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
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
