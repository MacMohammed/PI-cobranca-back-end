package main

import (
	"fatec/routes"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	routes.LoadRoutes(rtr)

	port := os.Getenv("PORT")

	fmt.Printf("Inicializando servidor na porta %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(rtr)))
}
