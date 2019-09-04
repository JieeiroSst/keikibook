package account

import (
	"errors"
	"log"

	"github.com/bwmarrin/snowflake"
)

func distributed() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		errors.New("err is error")
	}

	for i := 0; i < 100; i++ {
		id := n.Generate()
		log.Print(id)
		log.Print(id.Node())
		log.Print(id.Step())
		log.Print(id.Time())
	}
}
