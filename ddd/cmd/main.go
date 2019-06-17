package main

import (
	"fmt"
	"os"

	"github.com/fabiorodrigues/gostructure/ddd/domain/reservation"
	"github.com/fabiorodrigues/gostructure/ddd/infra/mongo"
	"github.com/fabiorodrigues/gostructure/ddd/infra/repositories"
)

type config struct {
	envVarOne string
	envVarTwo string
}

func main() {
	//retrieve some env vars
	config := config{
		envVarOne: os.Getenv("envOne"),
		envVarTwo: os.Getenv("envTwo"),
	}
	start(config)
}

func start(config config) {
	dbDriver := mongo.NewClient()
	bookRepo := repositories.NewBook(dbDriver)
	reservationService := reservation.NewService(bookRepo)

	if err := reservationService.Reserve("bookID"); err != nil {
		fmt.Println(err)
	}
}
