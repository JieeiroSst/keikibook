package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SignUp struct {
	ID             string
	FullName       string
	Email          string
	UserName       string
	Password       string
	RepeatPassword string
}

type Login struct {
	ID       string
	Email    string
	Password string
}

var SignUps []SignUp
var Logins []Login

//SignUp
//index
func getSignUps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	json.NewEncoder(w).Encode(SignUps)
}

//show
func getSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	params := mux.Vars(r)
	for _, item := range SignUps {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//create
func createSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	var newSignUp SignUp
	json.NewDecoder(r.Body).Decode(&newSignUp)
	newSignUp.ID = strconv.Itoa(len(SignUps) + 1)
	SignUps = append(SignUps, newSignUp)
	json.NewEncoder(w).Encode(newSignUp)
}

//update
func uppdateSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	params := mux.Vars(r)
	for i, item := range SignUps {
		if item.ID == params["id"] {
			SignUps = append(SignUps[:i], SignUps[i+1])
			var newSignUp SignUp
			json.NewDecoder(r.Body).Decode(&newSignUp)
			newSignUp.ID = params["id"]
			SignUps = append(SignUps, newSignUp)
			json.NewEncoder(w).Encode(newSignUp)
			return
		}
	}
}

//delete
func deleteSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	params := mux.Vars(r)
	for i, item := range SignUps {
		if item.ID == params["id"] {
			SignUps = append(SignUps[:1], SignUps[i+1])
			break
		}
	}
	json.NewEncoder(w).Encode(SignUps)
}

//Login
//index
func getLogins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	json.NewEncoder(w).Encode(Logins)
}

//show
func getLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	params := mux.Vars(r)
	for _, item := range Logins {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//create
func createLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	var newLogin Login
	json.NewDecoder(r.Body).Decode(&newLogin)
	newLogin.ID = strconv.Itoa(len(Logins) + 1)
	Logins = append(Logins, newLogin)
	json.NewEncoder(w).Encode(newLogin)
}

//update
func uppdateLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	params := mux.Vars(r)
	for i, item := range Logins {
		if item.ID == params["id"] {
			Logins = append(Logins[:i], Logins[i+1])
			var newLogin Login
			json.NewDecoder(r.Body).Decode(&newLogin)
			newLogin.ID = params["id"]
			Logins = append(Logins, newLogin)
			json.NewEncoder(w).Encode(newLogin)
			return
		}
	}
}

//delete
func deleteLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	params := mux.Vars(r)
	for i, item := range Logins {
		if item.ID == params["id"] {
			Logins = append(Logins[:1], Logins[i+1])
			break
		}
	}
	json.NewEncoder(w).Encode(Logins)
}
