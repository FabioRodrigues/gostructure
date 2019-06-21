package main

import (
	"fmt"
	"os"

	"github.com/fabiorodrigues/gostructure/ddd/domain/reservation"
	"github.com/fabiorodrigues/gostructure/ddd/infra/mongo"
	"github.com/fabiorodrigues/gostructure/ddd/infra/repositories"
	"github.com/fabiorodrigues/gostructure/ddd/infra/sendgrid"
)

type config struct {
	sendgridKey string
}

func main() {
	//retrieve some env vars
	config := config{
		sendgridKey: os.Getenv("SENDGRID_KEY"),
	}
	start(config)
}

func start(config config) {
	dbDriver := mongo.NewClient()
	bookRepo := repositories.NewBook(dbDriver)

	emailClient := sendgrid.NewClient(config.sendgridKey)

	reservationService := reservation.NewService(bookRepo, emailClient)

	if err := reservationService.Reserve("bookID"); err != nil {
		fmt.Println(err)
	}
}
