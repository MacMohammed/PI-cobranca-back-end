package models

import (
	"fatec/db"
	"fmt"
	"log"
)

//Struct criado para refletir a tabela do BD
type Client struct {
	Id_client    int    `json:"id,omitempty"`
	Name         string `json:"nome,omitempty"`
	Doc          string `json:"documento"`
	Name_fantasy string `json:"nome_fantasy"`
}

func AllClient() []Client {
	//Conecta com Postgres
	db := db.ConectBD()

	defer db.Close()

	//Executa uma  query ono postgres
	rows, err := db.Query("select id_cliente, nome, documento, nome_fantasia from cliente where ativo is true order by id_cliente;")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var clients []Client

	for rows.Next() {
		var client Client

		if err := rows.Scan(
			&client.Id_client,
			&client.Name,
			&client.Doc,
			&client.Name_fantasy,
		); err != nil {
			panic(err.Error())
		}

		clients = append(clients, client)
	}
	defer db.Close()
	return clients
}

func NewClient(client Client) int {
	db := db.ConectBD()

	insertDataBank, err := db.Prepare("insert into client(nome, doc, fantasy_name) values($1, $2, $3)")
	if err != nil {
		log.Println("Erro in client table exemption: ", err)
		return 1
	}

	insertDataBank.Exec(client.Name, client.Doc, client.Name_fantasy)
	defer db.Close()
	return 0
}

func DeleteClient(id_client int) int {
	db := db.ConectBD()

	deletClient, err := db.Prepare("delete from client where id_client=$1")
	if err != nil {
		fmt.Println("Error deleting value from table client: ", err)
		return 1
	}

	deletClient.Exec(id_client)
	defer db.Close()
	return 0
}

func GetClient(id_client int) Client {
	db := db.ConectBD()

	clientBank, err := db.Query("select * from client where id_client=$1", id_client)
	if err != nil {
		panic(err.Error())
	}

	client := Client{}
	for clientBank.Next() {
		var id_client int
		var name, name_fantasy, doc string

		err = clientBank.Scan(&id_client, &name, &doc, &name_fantasy)
		if err != nil {
			panic(err.Error())
		}
		// docConvertedForInt, err := strconv.Atoi(doc)
		// if err != nil {
		// 	log.Println("Conversion error: ", err)
		// }

		client.Id_client = id_client
		client.Name = name
		client.Doc = doc
		client.Name_fantasy = name_fantasy
	}
	defer db.Close()
	return client
}

func UpdateClient(client Client) int {
	db := db.ConectBD()

	UpdateClient, err := db.Prepare("update client set name=$1, doc=$2, fantasy_name=$3 where id_client=$4")
	if err != nil {
		return 1
	}

	UpdateClient.Exec(client.Name, client.Doc, client.Name_fantasy, client.Id_client)
	defer db.Close()
	return 0
}
