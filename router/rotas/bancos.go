package rotas

import (
	"fatec/controllers"
	"net/http"
)

var rotaBancos = []Rota{
	{
		URI:              "/bancos",
		Metodo:           http.MethodGet,
		Funcao:           controllers.AllBank,
		RequerAutenticao: false,
	},
}

// 	//Teste de Bank
// 	rtr.HandleFunc("/bancos", controllers.AllBank).Methods(http.MethodGet)
// 	rtr.HandleFunc("/GetBank", controllers.GetBank).Methods(http.MethodPost)
// 	rtr.HandleFunc("/UpdateBank", controllers.UpdateBank).Methods(http.MethodPost)
// 	rtr.HandleFunc("/InsertBank", controllers.InsertBank).Methods(http.MethodPost)
// 	rtr.HandleFunc("/DeleteBank", controllers.DeleteBank).Methods(http.MethodPost)
