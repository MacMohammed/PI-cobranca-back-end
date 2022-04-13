package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Office struct {
	Id_office   int
	Description string
}

func AllOffice() map[int]Office {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	selectAllOffice, err := db.Query("select * from office order by id_office asc")
	if err != nil {
		panic(err.Error())
	}

	//Lê a query e armazena num slice
	o := Office{}
	mapOffice := make(map[int]Office)
	for selectAllOffice.Next() {
		var id_office, description string

		err := selectAllOffice.Scan(&id_office, &description)
		if err != nil {
			panic(err.Error())
		}

		id_OfficeConvertedForInt, err := strconv.Atoi(id_office)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		o.Id_office = id_OfficeConvertedForInt
		o.Description = description

		mapOffice[o.Id_office] = o
	}
	defer db.Close()
	return mapOffice
}

func NewOffice(office Office) int {
	db := db.ConectBD()

	insertDataOffice, err := db.Prepare("insert into office(description) values($1)")
	if err != nil {
		log.Println("Error inserting into office table")
		return 1
	}

	insertDataOffice.Exec(office.Description)
	defer db.Close()
	return 0
}

func DeleteOffice(office Office) int {
	db := db.ConectBD()

	deletBank, err := db.Prepare("delete from office where id_office=$1")
	if err != nil {
		return 1
	}

	_, err = deletBank.Exec(office.Id_office)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetOffice(id_office int) Office {
	db := db.ConectBD()

	offices, err := db.Query("select * from office where id_office=$1", id_office)
	if err != nil {
		panic(err.Error())
	}

	office := Office{}
	for offices.Next() {
		var id_office, description string

		err = offices.Scan(&id_office, &description)
		if err != nil {
			panic(err.Error())
		}

		id_officeConvertedForInt, err := strconv.Atoi(id_office)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		office.Id_office = id_officeConvertedForInt
		office.Description = description
	}
	defer db.Close()
	return office
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateOffice(office Office) int {
	db := db.ConectBD()

	UpdateOffices, err := db.Prepare("update office set description = $1 where id_office=$2")
	if err != nil {
		fmt.Println("Error in updating the bank table")
		return 1
	}

	UpdateOffices.Exec(office.Description, office.Id_office)
	defer db.Close()
	return 0
}
