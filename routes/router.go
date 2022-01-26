package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/catatan-dapur/controller"
)

func BahanRouter(bahanController controller.BahanController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/bahans", bahanController.FindAll)
	router.POST("/api/bahan", bahanController.Create)

	return router
}
