package restapi

import (
	"database/sql"
	"encoding/json"
	"keikibook/account"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	db      *sql.DB
	SignUps []*account.SignUp
	Logins  []*account.Login
)

func GetSignUps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT id,FullName,Email,UserName,Password,RepeatPassword from SignUps")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next() {
		var signUp *account.SignUp
		err := result.Scan(&signUp.ID, &signUp.FullName, &signUp.Email,
			&signUp.Password, &signUp.RepeatPassword, &signUp.UserName)
		if err != nil {
			log.Fatal(err)
		}
		SignUps = append(SignUps, signUp)
	}
	json.NewEncoder(w).Encode(SignUps)
}

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT id,FullName,Email,UserName,Password,RepeatPassword FROM SignUps where id=?", params["id"])
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	var signUp account.SignUp
	for result.Next() {
		err := result.Scan(&signUp.ID, &signUp.FullName, &signUp.Email,
			&signUp.Password, &signUp.RepeatPassword, &signUp.UserName)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, item := range SignUps {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newSignUp *account.SignUp
	json.NewDecoder(r.Body).Decode(&newSignUp)
	newSignUp.ID = strconv.Itoa(len(SignUps) + 1)
	SignUps = append(SignUps, newSignUp)
	json.NewEncoder(w).Encode(newSignUp)
}

func UppdateSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range SignUps {
		if item.ID == params["ID"] {
			SignUps = append(SignUps[:i], SignUps[i+1])
			var newSinUp *account.SignUp
			json.NewDecoder(r.Body).Decode(&newSinUp)
			newSinUp.ID = params["ID"]
			SignUps = append(SignUps, newSinUp)
			json.NewEncoder(w).Encode(newSinUp)
			return
		}
	}
}

func DeleteSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parmass := mux.Vars(r)
	for i, item := range SignUps {
		if item.ID == parmass["ID"] {
			SignUps = append(SignUps[:1], SignUps[i+1])
			break
		}
	}
	json.NewEncoder(w).Encode(SignUps)
}

func GetLogins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT ID,Email,Password from account.Login")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next() {
		var login *account.Login
		err := result.Scan(&login.ID, &login.Email, &login.Password)
		if err != nil {
			log.Fatal(err)
		}
		Logins = append(Logins, login)
	}
	json.NewEncoder(w).Encode(Logins)
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT ID,Email,Password FROM account.Login")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	var login account.Login
	for result.Next() {
		err := result.Scan(&login.ID, &login.Email, &login.Password)
		if err != nil {

		}
	}
	for _, item := range Logins {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newLogin *account.Login
	json.NewDecoder(r.Body).Decode(&newLogin)
	newLogin.ID = strconv.Itoa(len(Logins) + 1)
	Logins = append(Logins, newLogin)
	json.NewEncoder(w).Encode(newLogin)
}

func UpdateLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parmas := mux.Vars(r)
	for i, item := range Logins {
		if item.ID == parmas["ID"] {
			Logins = append(Logins[:i], Logins[i+1])
			var newLogin *account.Login
			json.NewDecoder(r.Body).Decode(&newLogin)
			newLogin.ID = parmas["ID"]
			Logins = append(Logins, newLogin)
			json.NewEncoder(w).Encode(newLogin)
			return
		}
	}
}

func DeleteLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range Logins {
		if item.ID == params["ID"] {
			Logins = append(Logins[:1], Logins[i+1])
			break
		}
	}
	json.NewEncoder(w).Encode(Logins)
}
