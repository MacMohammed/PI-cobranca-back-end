package main

import (
	"fatec/routes"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	routes.LoadRoutes(rtr)

	port := os.Getenv("PORT")

	fmt.Printf("Inicializando servidor na porta %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
