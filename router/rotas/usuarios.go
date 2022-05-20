package rotas

import (
	"fatec/controllers"
	"net/http"
)

var rotaUsuarios = []Rota{
	{
		URI:              "/usuario",
		Metodo:           http.MethodGet,
		Funcao:           controllers.GetUser,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuario",
		Metodo:           http.MethodPost,
		Funcao:           controllers.InsertUser,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios",
		Metodo:           http.MethodGet,
		Funcao:           controllers.AllUsers,
		RequerAutenticao: false,
	},
}

// 	//Teste de User
// 	rtr.HandleFunc("/user", controllers.GetUser).Methods(http.MethodGet)
// 	rtr.HandleFunc("/users", controllers.AllUsers).Methods(http.MethodGet)
// 	rtr.HandleFunc("/user/update", controllers.UpdateUser).Methods(http.MethodPost)
// 	rtr.HandleFunc("/user/insert", controllers.InsertUser).Methods(http.MethodPost)
// 	rtr.HandleFunc("/user/delete", controllers.DeletetUser).Methods(http.MethodPost)
// 	rtr.HandleFunc("/login", controllers.EnterUser).Methods(http.MethodPost)
