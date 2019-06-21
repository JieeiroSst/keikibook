package model

import (
	"time"

	"github.com/facebookgo/clock"
)

type Application struct {
	Clock clock.Clock
}

func createClock() {
	var app Application
	app.Clock = clock.New()
	go func() {
		app.Clock.Now()
		time.Sleep(2 * time.Millisecond)
	}()
}
