package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/catatan-dapur/model/web"
	"github.com/zenklot/catatan-dapur/service"
)

type KategoriController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type KategoriControllerImpl struct {
	KategoriService service.KategoriService
}

func NewKategoriController(kategoriService service.KategoriService) *KategoriControllerImpl {
	return &KategoriControllerImpl{
		KategoriService: kategoriService,
	}
}

func (controller *KategoriControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kategoriCreateRequest := web.KategoriCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&kategoriCreateRequest); err != nil {
		panic(err)
	}

	kategoriResponse := controller.KategoriService.Create(kategoriCreateRequest)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   kategoriResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}

}

func (controller *KategoriControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kategoriUpdateRequest := web.KategoriUpdateRequest{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&kategoriUpdateRequest); err != nil {
		panic(err)
	}

	kategoriResponse := controller.KategoriService.Update(kategoriUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kategoriResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *KategoriControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kategoriId := params.ByName("kategoriId")
	id, err := strconv.Atoi(kategoriId)
	if err != nil {
		panic(err)
	}

	controller.KategoriService.Delete(id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *KategoriControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	kategoriResponse := controller.KategoriService.FindAll()

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kategoriResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *KategoriControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kategoriId := params.ByName("kategoriId")
	id, err := strconv.Atoi(kategoriId)
	if err != nil {
		panic(err)
	}

	kategoriResponse := controller.KategoriService.FindById(id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kategoriResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}
