package rotas

import (
	"fatec/controllers"
	"net/http"
)

var rotaTransacoes = []Rota{
	{
		URI:              "/transacoes",
		Metodo:           http.MethodGet,
		Funcao:           controllers.AllTransaction,
		RequerAutenticao: false,
	},
	{
		URI:              "/transacao",
		Metodo:           http.MethodPost,
		Funcao:           controllers.InsertTransaction,
		RequerAutenticao: false,
	},
	{
		URI:              "/trasancao/baixar/{id_transacao}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.TransacaoBaixar,
		RequerAutenticao: false,
	},
	{
		URI:              "/trasancao/cancelar/{id_transacao}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.TransacaoCancelar,
		RequerAutenticao: false,
	},
	{
		URI:              "/trasancao/extornar/{id_transacao}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.TransacaoExtornar,
		RequerAutenticao: false,
	},
}

// 	rtr.HandleFunc("/transacoes", controllers.AllTransaction).Methods(http.MethodGet)
// 	rtr.HandleFunc("/GetTransaction", controllers.GetTransaction).Methods(http.MethodPost)
// 	rtr.HandleFunc("/UpdateTransaction", controllers.UpdateTransaction).Methods(http.MethodPost)
// 	rtr.HandleFunc("/transacao", controllers.InsertTransaction).Methods(http.MethodPost)
// 	rtr.HandleFunc("/DeleteTransaction", controllers.DeleteTransaction).Methods(http.MethodPost)
// 	rtr.HandleFunc("/trasancao/baixar/{id_transacao}", controllers.TransacaoBaixar).Methods(http.MethodPut)
// 	rtr.HandleFunc("/trasancao/cancelar/{id_transacao}", controllers.TransacaoCancelar).Methods(http.MethodPut)
