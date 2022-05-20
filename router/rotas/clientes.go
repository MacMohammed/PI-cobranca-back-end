package rotas

import (
	"fatec/controllers"
	"net/http"
)

var rotaClientes = []Rota{
	{
		URI:              "/clientes",
		Metodo:           http.MethodGet,
		Funcao:           controllers.AllClient,
		RequerAutenticao: false,
	},
}

// 	//Teste de Client
// 	rtr.HandleFunc("/clientes", controllers.AllClient).Methods(http.MethodGet)
// 	rtr.HandleFunc("/GetClient", controllers.GetClient).Methods(http.MethodPost)
// 	rtr.HandleFunc("/UpdateClient", controllers.UpdateClient).Methods(http.MethodPost)
// 	rtr.HandleFunc("/InsertClient", controllers.InsertClient).Methods(http.MethodPost)
// 	rtr.HandleFunc("/DeleteClient", controllers.DeleteClient).Methods(http.MethodPost)
