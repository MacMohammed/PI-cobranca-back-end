package main

import (
	"fatec/router"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/handlers"
)

func main() {
	rtr := router.GerarRotas()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8001"
	}

	fmt.Printf("Inicializando servidor na porta %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(rtr)))
}
