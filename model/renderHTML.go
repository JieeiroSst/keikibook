package model //renderHTML.go

import (
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type SignUp struct {
	FullName       string
	Email          string
	UserName       string
	Password       string
	RepeatPassword string
}

type Login struct {
	Email    string
	Password string
}

func handleSignUp(c echo.Context, err error) error {
	u := new(SignUp)
	if err = c.Bind(u); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, u)
}

func handleLogin(c echo.Context, err error) error {
	u := new(Login)
	if err = c.Bind(u); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, u)
}

func RenderHome(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("keikibook/templates/home/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, nil)
}

func RenderWall(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("keikibook/templates/wallPage/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("keikibook/templates/Login/index.html")
	if err != nil {
		log.Fatal(err)
	}
	details := Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	_ = details
	tpl.Execute(w, struct{ Success bool }{true})
}

func RenderSignUp(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("keikibook/templates/sign-up/index.html")
	if err != nil {
		log.Fatal(err)
	}
	details := SignUp{
		FullName:       r.FormValue("fullname"),
		Email:          r.FormValue("email"),
		UserName:       r.FormValue("username"),
		Password:       r.FormValue("password"),
		RepeatPassword: r.FormValue("repeat-password"),
	}
	_ = details
	tpl.Execute(w, struct{ Success bool }{true})
}
