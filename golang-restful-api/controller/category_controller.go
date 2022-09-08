package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Sebenernya bisa langsung define struct tapi ini kita buat service dulu biar ke track apa yang kita buat
type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
