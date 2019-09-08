package savejson

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"keikibook/database"
	"log"
	"net/http"
)

var conn *sql.DB

func ReadJsonLogin(w http.ResponseWriter, r *http.Request) {
	conn := database.HandleDataBase()
	file, err := json.MarshalIndent(conn, "", "")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("Login.json", file, 0644)
}
