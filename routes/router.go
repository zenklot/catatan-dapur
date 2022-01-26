package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/catatan-dapur/app/exception"
	"github.com/zenklot/catatan-dapur/controller"
)

func BahanRouter(bahanController controller.BahanController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/bahans", bahanController.FindAll)
	router.GET("/api/bahan/:bahanId", bahanController.FindById)
	router.DELETE("/api/bahan/:bahanId", bahanController.Delete)
	router.POST("/api/bahan", bahanController.Create)
	router.PUT("/api/bahan", bahanController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func KategoriRouter(kategoriController controller.KategoriController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/kategories", kategoriController.FindAll)
	router.GET("/api/kategori/:kategoriId", kategoriController.FindById)
	router.DELETE("/api/kategori/:kategoriId", kategoriController.Delete)
	router.POST("/api/kategori", kategoriController.Create)
	router.PUT("/api/kategori", kategoriController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}
