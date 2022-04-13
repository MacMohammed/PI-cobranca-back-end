package models

import (
	"fatec/db"
	"fmt"
)

//Struct criado para refletir a tabela do BD
type File struct {
	Id_file           int
	Date              string
	Types             string
	Fk_id_transaction int
}

func AllFile() map[int]File {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	selectAllFile, err := db.Query("select * from file order by id_file asc")
	if err != nil {
		panic(err.Error())
	}

	//Lê a query e armazena num slice
	//t := Transaction{}
	mapFile := make(map[int]File)
	for selectAllFile.Next() {
		var f File

		err := selectAllFile.Scan(&f.Id_file, &f.Date, &f.Types, &f.Fk_id_transaction)
		if err != nil {
			panic(err.Error())
		}

		// id_transactionConvertedForInt, err := strconv.Atoi(id_transaction)
		// if err != nil {
		// 	log.Println("Conversion error: ", err)
		// }

		// t.Id_transaction = id_transactionConvertedForInt
		// t.Maturity_date = maturity_date
		// t.Invoice = invoice
		// t.Issue_date = issue_date
		mapFile[f.Id_file] = f
	}
	defer db.Close()
	return mapFile
}

func NewFile(file File) {
	db := db.ConectBD()

	insertDataFile, err := db.Prepare("insert into file(date, types, fk_id_transaction) values($1, $2,$3)")
	if err != nil {
		panic(err.Error())
	}

	insertDataFile.Exec(file.Date, file.Types, file.Fk_id_transaction)
	defer db.Close()
}

func DeleteFile(file File) int {
	db := db.ConectBD()

	deletFile, err := db.Prepare("delete from file where id_file=$1")
	if err != nil {
		return 1
	}

	_, err = deletFile.Exec(file.Id_file)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetFile(id_file int) File {
	db := db.ConectBD()

	file, err := db.Query("select * from file where id_file=$1", id_file)
	if err != nil {
		panic(err.Error())
	}

	f := File{}
	for file.Next() {
		//var f File

		err = file.Scan(&f.Id_file, &f.Date, &f.Types, &f.Fk_id_transaction)
		if err != nil {
			panic(err.Error())
		}

		// id_transactionConvertedForInt, err := strconv.Atoi(id_transaction)
		// if err != nil {
		// 	log.Println("Conversion error: ", err)
		// }

		// transaction.Id_transaction = id_transactionConvertedForInt
		// transaction.Maturity_date = maturity_date
		// transaction.Invoice = invoice
		// transaction.Issue_date = issue_date
	}
	defer db.Close()
	return f
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateFile(file File) int {
	db := db.ConectBD()

	UpdateFiles, err := db.Prepare("update file set date = $1, types=$2, fk_id_transaction = $3 where id_file=$4")
	if err != nil {
		fmt.Println("Error in updating the bank table: ", err)
		return 1
	}

	UpdateFiles.Exec(file.Date, file.Types, file.Fk_id_transaction)
	defer db.Close()
	return 0
}
