package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/catatan-dapur/model/web"
	"github.com/zenklot/catatan-dapur/service"
)

type BahanController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type BahanControllerImpl struct {
	BahanService service.BahanService
}

func NewBahanController(bahanService service.BahanService) *BahanControllerImpl {
	return &BahanControllerImpl{
		BahanService: bahanService,
	}
}
func (controller *BahanControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bahanCreateRequest := web.BahanCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&bahanCreateRequest); err != nil {
		panic(err)
	}

	bahanResponse := controller.BahanService.Create(bahanCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   bahanResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}

}

func (controller *BahanControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bahanUpdateRequest := web.BahanUpdateRequest{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&bahanUpdateRequest); err != nil {
		panic(err)
	}

	bahanResponse := controller.BahanService.Update(bahanUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bahanResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *BahanControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bahanId := params.ByName("bahanId")
	id, err := strconv.Atoi(bahanId)
	if err != nil {
		panic(err)
	}

	controller.BahanService.Delete(id)
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

func (controller *BahanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bahanId := params.ByName("bahanId")
	id, err := strconv.Atoi(bahanId)
	if err != nil {
		panic(err)
	}

	bahanResponse := controller.BahanService.FindById(id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bahanResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *BahanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bahanResponse := controller.BahanService.FindAll()
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bahanResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}
