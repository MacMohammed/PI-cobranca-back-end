package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

//Função para Conectar ao Banco de Dados Postgres
func ConectBD() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
