package rotas

import (
	"fatec/controllers"
	"net/http"
)

var rotaCargo = []Rota{
	{
		URI:              "/cargo",
		Metodo:           http.MethodGet,
		Funcao:           controllers.GetCargos,
		RequerAutenticao: false,
	},
	{
		URI:              "/cargo/{id_cargo}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.GetOffice,
		RequerAutenticao: false,
	},
}

// 	rtr.HandleFunc("/cargo", controllers.GetCargos).Methods(http.MethodGet)
// 	rtr.HandleFunc("/cargo/{id_cargo}", controllers.GetOffice).Methods(http.MethodGet)
// 	rtr.HandleFunc("/cargo/atualizar/{id_cargo}", controllers.UpdateOffice).Methods(http.MethodPost)
// 	rtr.HandleFunc("/cargo/criar", controllers.CreateOffice).Methods(http.MethodPost)
// 	rtr.HandleFunc("/cargo/excluir/{id_cargo}", controllers.DeleteOffice).Methods(http.MethodDelete)
