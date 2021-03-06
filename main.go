package main

import (
	"fmt"
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

	kategoriRepository := repository.NewKategoriRepository()
	kategoriService := service.NewKategoriService(kategoriRepository, db, validate)
	kategoriController := controller.NewKategoriController(kategoriService)

	resepRepository := repository.NewResepRepository()
	resepService := service.NewResepService(resepRepository, db, validate)
	resepController := controller.NewResepController(resepService)

	router := routes.NewRouter(bahanController, kategoriController, resepController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	fmt.Println("Server Started at http://localhost:3000")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
