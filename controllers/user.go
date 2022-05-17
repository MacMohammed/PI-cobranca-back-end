package controllers

import (
	"encoding/json"
	"fatec/autentication"
	"fatec/models"
	"fatec/respostas"
	"fatec/securit"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Insere um novo usuario
func InsertUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, _ := ioutil.ReadAll(r.Body)                   //captura o conteudo do body (retorno é um lista de bytes)
	if err := json.Unmarshal(body, &user); err != nil { //encapsula em banco o que vem no body
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	retur := models.NewUser(user)
	respostas.JSON(w, http.StatusCreated, retur)
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

//Metodo para entrar na area do usuario
func EnterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		respostas.JSON(w, http.StatusBadRequest, err)
		return
	}

	//Pega as informações do usuario no BD
	u := models.GetUserByName(user.Name)

	//Verifica se os hash da senha são iguais
	err := securit.VerificarSenha(u.Password, user.Password)

	if err != nil {
		respostas.JSON(w, http.StatusForbidden, "A senhas não são iguais...")
		return
	}

	//Criando Token
	token, err := autentication.CreateToken(uint64(u.Id_user))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, err)
		return
	}

	var dataAutentication models.DataAutentication
	dataAutentication.Name = user.Name
	dataAutentication.AccountID = fmt.Sprintf("%d", user.Id_user)
	dataAutentication.Token = token

	// Gravando tokie no Cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	w.Header().Set("Authorization", token)
	respostas.JSON(w, http.StatusCreated, dataAutentication)
}
