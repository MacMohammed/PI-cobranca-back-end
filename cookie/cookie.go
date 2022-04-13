package cookie

import (
	"fmt"
	"net/http"
)

func SetCokieUserAutetication(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		fmt.Println("favicon")
		return
	}

	//Verfica se existe algum cookie de autenticação para o site
	cookie, err := r.Cookie("fatec.Autentication.Cookie")
	if err == http.ErrNoCookie {
		fmt.Println(1)
	} else {
		if cookie.Value != "" {
			//Retorna os dados do usuario já autenticado
			fmt.Println(0)
		} else {
			fmt.Println(2)
		}
	}
}

//Apaga o valor do cookie de autenticação para o site
func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	//Gravando Cookie o valor vazio
	cookie := &http.Cookie{
		Name:  "fatec.Autentication.Cookie",
		Value: "",
	}
	http.SetCookie(w, cookie)
	fmt.Println(cookie)
	fmt.Println("Cookie Apagado")
	return
}
