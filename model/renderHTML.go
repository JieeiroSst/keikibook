package model

import (
	"html/template"
	"net/http"
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

func RenderHome(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("/src/templates/home/index.html")

	tpl.Execute(w, nil)
}

func RenderWall(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("/src/templates/wallPage/index.html")

	tpl.Execute(w, nil)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("/src/templates/Login/index.html")
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
	}
	details := Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	_ = details
	tpl.Execute(w, struct{ Success bool }{true})
}

func RenderSignUp(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("/src/templates/sign-up/index.html")
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
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
