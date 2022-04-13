package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Função para Conectar ao Banco de Dados Postgres
func ConectBD() *sql.DB {
	connectino := "user=gebeewwzjtnuar dbname=dap7h89poiespl password=d49af27e6b6de17aaf323d7839547b58fe4b8a9296faa8f5f9d149931b229068 host=ec2-44-194-4-127.compute-1.amazonaws.com sslmode=disable"
	db, err := sql.Open("postgres", connectino)
	if err != nil {
		panic(err.Error())
	}
	return db
}
