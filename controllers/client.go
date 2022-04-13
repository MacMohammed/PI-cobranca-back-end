package controllers

import (
	"encoding/json"
	"fatec/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Insere um novo cliente
func InsertClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := ioutil.ReadAll(r.Body)                     //captura o conteudo do body (retorno é um lista de bytes)
	if err := json.Unmarshal(body, &client); err != nil { //encapsula em banco o que vem no body
		fmt.Print(err)
	}
	retur := models.NewClient(client) //altera no banco de dados
	json.NewEncoder(w).Encode(retur)
}

//Atualiza cliente
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := ioutil.ReadAll(r.Body)                     //captura o conteudo do body (retorno é um lista de bytes)
	if err := json.Unmarshal(body, &client); err != nil { //encapsula em banco o que vem no body
		fmt.Print(err)
	}
	retur := models.UpdateClient(client) //altera no banco de dados
	json.NewEncoder(w).Encode(retur)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &client); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.DeleteClient(client.Id_client)
	json.NewEncoder(w).Encode(ret)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &client); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetClient(client.Id_client)

	json.NewEncoder(w).Encode(ret)
}

func AllClient(w http.ResponseWriter, r *http.Request) {
	ret := models.AllClient()
	json.NewEncoder(w).Encode(ret)
}
