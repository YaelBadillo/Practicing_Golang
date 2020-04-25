package main

import (
	"log"

	"github.com/YaelJBS/Practicing_Golang/db"
	"github.com/YaelJBS/Practicing_Golang/handlers"
)

func main() {
	if db.CheckConnection() == false {
		log.Fatal("Could not connected to MongoDB.")
		return
	}

	handlers.Handlers()
}
