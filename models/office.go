package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Cargo struct {
	IDCargo   int    `json:"id,omitempty"`
	Descricao string `json:"descricao"`
}

func GetCargos() ([]Cargo, error) {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	rows, err := db.Query("select id_cargo, descricao from cargo order by id_cargo;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cargos []Cargo

	for rows.Next() {
		var cargo Cargo
		if err = rows.Scan(
			&cargo.IDCargo,
			&cargo.Descricao); err != nil {
			return nil, err
		}

		cargos = append(cargos, cargo)
	}

	defer db.Close()

	return cargos, nil
}

func NewOffice(office Cargo) int {
	db := db.ConectBD()

	insertDataOffice, err := db.Prepare("insert into office(description) values($1)")
	if err != nil {
		log.Println("Error inserting into office table")
		return 1
	}

	insertDataOffice.Exec(office.Descricao)
	defer db.Close()
	return 0
}

func DeleteOffice(office Cargo) int {
	db := db.ConectBD()

	deletBank, err := db.Prepare("delete from office where id_office=$1")
	if err != nil {
		return 1
	}

	_, err = deletBank.Exec(office.IDCargo)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetOffice(id_office int) Cargo {
	db := db.ConectBD()

	offices, err := db.Query("select * from cargo where id_cargo=$1", id_office)
	if err != nil {
		panic(err.Error())
	}

	office := Cargo{}
	for offices.Next() {
		var id_cargo, descricao string

		err = offices.Scan(&id_cargo, &descricao)
		if err != nil {
			panic(err.Error())
		}

		id_officeConvertedForInt, err := strconv.Atoi(id_cargo)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		office.IDCargo = id_officeConvertedForInt
		office.Descricao = descricao
	}
	defer db.Close()
	return office
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateOffice(office Cargo) int {
	db := db.ConectBD()

	UpdateOffices, err := db.Prepare("update office set description = $1 where id_office=$2")
	if err != nil {
		fmt.Println("Error in updating the bank table")
		return 1
	}

	UpdateOffices.Exec(office.Descricao, office.IDCargo)
	defer db.Close()
	return 0
}
