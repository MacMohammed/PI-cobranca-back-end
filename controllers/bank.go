package controllers

import (
	"encoding/json"
	"fatec/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AllBank(w http.ResponseWriter, r *http.Request) {
	banks := models.AllBank()
	json.NewEncoder(w).Encode(banks)
}

func GetBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &bank); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetBank(bank.Id_bank)

	json.NewEncoder(w).Encode(ret)
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &bank); err != nil {
		fmt.Print(err)
	}

	ret := models.UpdateBank(bank)
	json.NewEncoder(w).Encode(ret)
}

func InsertBank(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Errorf("ocorreu um erro %s", err))
		return
	}

	fmt.Println(body)

	// var bank models.Bank

	// if err := json.Unmarshal(body, &bank); err != nil {
	// 	json.NewEncoder(w).Encode(1)
	// }

	// models.NewBank(bank)
	// json.NewEncoder(w).Encode(0)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &bank); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	models.DeleteBank(bank)
	json.NewEncoder(w).Encode(0)
}
