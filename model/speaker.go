package model

import (
	"net/http"
	"strings"
)

func sayer(w http.ResponseWriter, r *http.Request) {
	messages := r.URL.Path
	messages = strings.TrimPrefix(messages, "/")
	messages = "hello" + messages

	w.Write([]byte(messages))
}
