package main

import (
	"fmt"
	"keikibook/api"
	"keikibook/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func noFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1><b>keikibook no found</b></h1>")
	fmt.Fprint(w, "<h2>We apologize for this incident</h2>")
	fmt.Fprint(w, "<h3>This problem will be fixed in the shortest time</h3>")
	fmt.Fprint(w, "<p>You can contact our company to fix the problem</p>")
	fmt.Fprint(w, "<p>email:luumanhquan.91@gmail.com</p>")
}

func main() {
	port := ":8080"
	r := mux.NewRouter()
	//render front end
	r.HandleFunc("/", noFound)
	r.HandleFunc("/keikibook", model.RenderHome)
	r.HandleFunc("/keikibook/login", model.RenderLogin)
	r.HandleFunc("/keikibook/sign-up", model.RenderSignUp)
	r.HandleFunc("/{id}", model.RenderWall)
	//api database
	//sign up
	r.HandleFunc("/keikibook", api.getSignUps).Methods("GET")
	r.HandleFunc("/keikibook/{id}", api.getSignUp).Methods("GET")
	r.HandleFunc("/keikibook", api.createSignUp).Methods("POST")
	r.HandleFunc("/keikibook/{id}", api.uppdateSignUp).Methods("POST")
	r.HandleFunc("/keikibook/{id}", api.deleteSignUp).Methods("DELETE")

	//Login
	r.HandleFunc("/keikibook", api.getLogins).Methods("GET")
	r.HandleFunc("/keikibook/{id}", api.getLogin).Methods("GET")
	r.HandleFunc("/keikibook", api.createLogin).Methods("POST")
	r.HandleFunc("/keikibook/{id}", api.uppdateLogin).Methods("POST")
	r.HandleFunc("/keikibook/{id}", api.deleteLogin).Methods("DELETE")

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}

}
