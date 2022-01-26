package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/catatan-dapur/model/web"
	"github.com/zenklot/catatan-dapur/service"
)

type ResepController interface {
	Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ResepControllerImpl struct {
	ResepService service.ResepService
}

func NewResepController(resepService service.ResepService) *ResepControllerImpl {
	return &ResepControllerImpl{
		ResepService: resepService,
	}
}

func (controller *ResepControllerImpl) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	t := template.Must(template.ParseFiles("./views/resep.gohtml"))

	resepResponse := controller.ResepService.FindAll()

	t.ExecuteTemplate(writer, "resep.gohtml", resepResponse)
}

func (controller *ResepControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	resepCreateRequest := web.ResepCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&resepCreateRequest); err != nil {
		panic(err)
	}

	resepResponse := controller.ResepService.Create(resepCreateRequest)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   resepResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *ResepControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	resepResponse := controller.ResepService.FindAll()

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resepResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *ResepControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	resepId := params.ByName("resepId")
	resepResponse := controller.ResepService.FindById(resepId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resepResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(webResponse); err != nil {
		panic(err)
	}
}
