package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/catatan-dapur/app/exception"
	"github.com/zenklot/catatan-dapur/controller"
)

func NewRouter(bahanController controller.BahanController, kategoriController controller.KategoriController, resepController controller.ResepController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/bahans", bahanController.FindAll)
	router.GET("/api/bahan/:bahanId", bahanController.FindById)
	router.DELETE("/api/bahan/:bahanId", bahanController.Delete)
	router.POST("/api/bahan", bahanController.Create)
	router.PUT("/api/bahan", bahanController.Update)

	router.GET("/api/kategories", kategoriController.FindAll)
	router.GET("/api/kategori/:kategoriId", kategoriController.FindById)
	router.DELETE("/api/kategori/:kategoriId", kategoriController.Delete)
	router.POST("/api/kategori", kategoriController.Create)
	router.PUT("/api/kategori", kategoriController.Update)

	router.POST("/api/resep", resepController.Create)
	router.GET("/api/reseps", resepController.FindAll)
	router.GET("/api/resep/:resepId", resepController.FindById)
	router.PanicHandler = exception.ErrorHandler

	return router
}
