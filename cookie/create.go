package cookie

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func writeCookie(c echo.Context) error {
	Cookie := new(http.Cookie)
	Cookie.Name = "keikibook"
	Cookie.Value = "Sst"
	Cookie.Expires = time.Now().Add(25 * time.Hour)
	Cookie.Domain = "www.keikibook.com"
	c.SetCookie(Cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func readCookie(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookie")
}
