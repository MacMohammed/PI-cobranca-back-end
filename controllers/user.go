package controllers

import (
	"encoding/json"
	"fatec/models"
	"fatec/respostas"
	"fatec/securit"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Insere um novo usuario
func InsertUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body) //captura o conteudo do body (retorno é um lista de bytes)

	if err != nil { //encapsula em banco o que vem no body
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil { //encapsula em banco o que vem no body
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	passwordString, err := securit.HashPassword(user.Password)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = string(passwordString)

	user.Id_user, err = models.NewUser(user)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, user)
}

func DeletetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, _ := ioutil.ReadAll(r.Body)                   //captura o conteudo do body (retorno é um lista de bytes)
	if err := json.Unmarshal(body, &user); err != nil { //encapsula em banco o que vem no body
		fmt.Print(err)
	}
	retur := models.DeleteUser(user)
	json.NewEncoder(w).Encode(retur)
}

//Altera os dados do usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Print(err)
	}

	retur := models.UpdateUser(user)
	json.NewEncoder(w).Encode(retur)
}

//Retorna os dados do usuario
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	userBD := models.GetUser(user.Id_user)

	json.NewEncoder(w).Encode(userBD)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.AllUser()

	json.NewEncoder(w).Encode(users)
}
