package controllers

import (
	"encoding/json"
	"fatec/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AllFile(w http.ResponseWriter, r *http.Request) {
	files := models.AllFile()
	json.NewEncoder(w).Encode(files)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	var files models.File

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &files); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetFile(files.Id_file)

	json.NewEncoder(w).Encode(ret)
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	var file models.File
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &file); err != nil {
		fmt.Print(err)
	}

	ret := models.UpdateFile(file)
	json.NewEncoder(w).Encode(ret)
}

func InsertFile(w http.ResponseWriter, r *http.Request) {
	var file models.File
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &file); err != nil {
		fmt.Print(err)
	}

	models.NewFile(file)
	json.NewEncoder(w).Encode(0)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	var file models.File
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &file); err != nil {
		fmt.Print(err)
	}

	models.DeleteFile(file)
	json.NewEncoder(w).Encode(0)
}
