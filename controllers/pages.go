package controllers

import (
	"fatec/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)
}
