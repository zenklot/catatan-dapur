package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zenklot/catatan-dapur/app"
	"github.com/zenklot/catatan-dapur/controller"
	"github.com/zenklot/catatan-dapur/repository"
	"github.com/zenklot/catatan-dapur/routes"
	"github.com/zenklot/catatan-dapur/service"
)

func main() {

	db := app.NewDB()

	validate := validator.New()
	bahanRepository := repository.NewBahanRepository()
	bahanService := service.NewBahanService(bahanRepository, db, validate)
	bahanController := controller.NewBahanController(bahanService)

	router := routes.BahanRouter(bahanController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
