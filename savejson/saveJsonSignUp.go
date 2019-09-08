package savejson

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"keikibook/database"
	"log"
	"net/http"
)

var conm *sql.DB

func ReadJsonSignUp(w http.ResponseWriter, r *http.Request) {
	conm := database.HandleDataBase()
	file, err := json.MarshalIndent(conm, "", "")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("keikibook/restapi/SignUp.json", file, 0644)
}
