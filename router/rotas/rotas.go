package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI              string
	Metodo           string
	Funcao           func(http.ResponseWriter, *http.Request)
	RequerAutenticao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotaBancos...)
	rotas = append(rotas, rotaTransacoes...)
	rotas = append(rotas, rotaCargo...)
	rotas = append(rotas, rotaClientes...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

// func LoadRoutes(rtr *mux.Router) {

// 	//Teste de Cookie
// 	rtr.HandleFunc("/deleteCookie", cookie.DeleteCookie).Methods(http.MethodGet)

// 	//Teste de File
// 	rtr.HandleFunc("/AllFile", controllers.AllFile).Methods(http.MethodGet)
// 	rtr.HandleFunc("/GetFile", controllers.GetFile).Methods(http.MethodPost)
// 	rtr.HandleFunc("/UpdateFile", controllers.UpdateFile).Methods(http.MethodPost)
// 	rtr.HandleFunc("/InsertFile", controllers.InsertFile).Methods(http.MethodPost)
// 	rtr.HandleFunc("/DeleteFile", controllers.DeleteFile).Methods(http.MethodPost)
// }
