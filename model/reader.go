package model

import (
	"io"
	"net/http"
	"strings"
)

var received io.Reader

func Reader(w http.ResponseWriter, r *http.Request) {
	var message string
	received = strings.NewReader("read message from sender")
	received.Read([]byte(message))
}
