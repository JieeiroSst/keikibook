package model

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func timer(w http.ResponseWriter, r *http.Request) {
	var inform string = "loading..."
	fmt.Println("Runnig time...")
	fmt.Fprint(w, time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(inform)
		os.Remove(inform)
		time.Sleep(10)
	}
}
