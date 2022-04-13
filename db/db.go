package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Função para Conectar ao Banco de Dados Postgres
func ConectBD() *sql.DB {
	strconnection := "user=nknpxencptienk dbname=d7l4rvnlb3opsu password=4916af49ef0d66da88c91fd3067b54ec7a17f5b974391c07346b2210070648af host=ec2-18-215-96-22.compute-1.amazonaws.com sslmode=disable"
	db, err := sql.Open("postgres", strconnection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
