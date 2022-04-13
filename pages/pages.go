package pages

import (
	"fatec/utils"
	"net/http"
)

func CarregarTelaCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)
}

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)
}

func CarregarTelaPgInicial(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "pag_inicial.html", nil)
}

func CarregarTelaPgCadUser(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "cad_usuario.html", nil)
}
func CarregarTelaPgCadCliente(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "cad_cliente.html", nil)
}
func CarregarTelaPgBancol(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "cad_banco.html", nil)
}
