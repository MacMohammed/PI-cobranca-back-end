package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Bank struct {
	Id_bank int    `json:"id,omitempty"`
	DtHrReg string `json:"dt_hr_reg,omitempty"`
	Name    string `json:"nome,omitempty"`
	Cod     int    `json:"codigo_febraban,omitempty"`
	Ativo   bool   `json:"ativo,omitempty"`
}

func AllBank() []Bank {
	//Conecta com Postgres
	db := db.ConectBD()
	defer db.Close()

	//Executa uma  query ono postgres
	rows, err := db.Query("select id_banco, nome, codigo from banco order by id_banco;")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var bancos []Bank

	for rows.Next() {
		var banco Bank

		if err := rows.Scan(
			&banco.Id_bank,
			&banco.Name,
			&banco.Cod,
		); err != nil {
			panic(err.Error())
		}

		bancos = append(bancos, banco)
	}

	return bancos
}

func NewBank(bank Bank) (int, error) {
	db := db.ConectBD()

	query := `insert into banco(nome, codigo) values($1, $2) returning id_banco;`

	smt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id_banco int
	if err = smt.QueryRow(bank.Name, bank.Cod).Scan(&id_banco); err != nil {
		return 0, err
	}

	return id_banco, nil

}

func DeleteBank(bank Bank) int {
	db := db.ConectBD()

	deletBank, err := db.Prepare("delete from bank where id_bank=$1")
	if err != nil {
		return 1
	}

	_, err = deletBank.Exec(bank.Id_bank)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetBank(id_bank int) Bank {
	db := db.ConectBD()

	banks, err := db.Query("select * from bank where id_bank=$1", id_bank)
	if err != nil {
		panic(err.Error())
	}

	bank := Bank{}
	for banks.Next() {
		var name, cod, id_bank string

		err = banks.Scan(&cod, &name, &id_bank)
		if err != nil {
			panic(err.Error())
		}

		id_bankConvertedForInt, err := strconv.Atoi(id_bank)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		codConvertedForInt, err := strconv.Atoi(cod)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		bank.Id_bank = id_bankConvertedForInt
		bank.Name = name
		bank.Cod = codConvertedForInt
	}
	defer db.Close()
	return bank
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateBank(bank Bank) int {
	db := db.ConectBD()

	UpdateBanks, err := db.Prepare("update bank set name=$1, cod=$2 where id_bank=$3")
	if err != nil {
		fmt.Println("Error in updating the bank table")
		return 1
	}

	UpdateBanks.Exec(bank.Name, bank.Cod, bank.Id_bank)
	defer db.Close()
	return 0
}
