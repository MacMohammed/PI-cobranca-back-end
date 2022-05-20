package router

import (
	"fatec/router/rotas"

	"github.com/gorilla/mux"
)

func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
