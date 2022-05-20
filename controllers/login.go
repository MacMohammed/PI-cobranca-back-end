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

func Login(w http.ResponseWriter, r *http.Request) {
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
