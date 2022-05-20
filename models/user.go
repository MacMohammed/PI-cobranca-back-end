package models

import (
	"errors"
	"fatec/db"
	"log"
	"strconv"
	"strings"
)

//Struct criado para refletir a tabela do BD
type User struct {
	Id_user             int    `json:"id_usuario,omitempty"`
	Name                string `json:"nome,omitempty"`
	Password            string `json:"senha,omitempty"`
	Active              bool   `json:"ativo,omitempty"`
	Fk_office_id_office int    `json:"id_cargo,omitempty"`
}

func AllUser() map[int]User {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	// selectAllUser, err := db.Query("select * from user_system order by id_user_system asc")
	selectAllUser, err := db.Query("select * from usuario order by id_usuario")
	if err != nil {
		panic(err.Error())
	}

	//Lê a query e armazena num slice
	u := User{}
	users := make(map[int]User)
	for selectAllUser.Next() {
		var id_user, name, password, active, fk_office_id_office string

		err := selectAllUser.Scan(&id_user, &name, &password, &active, &fk_office_id_office)
		if err != nil {
			panic(err.Error())
		}

		id_userConvertedForInt, err := strconv.Atoi(id_user)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		fk_office_id_officeConvertedForInt, err := strconv.Atoi(fk_office_id_office)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		activeConvertedForBool, err := strconv.ParseBool(active)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		u.Id_user = id_userConvertedForInt
		u.Name = name
		u.Password = password
		u.Active = activeConvertedForBool
		u.Fk_office_id_office = fk_office_id_officeConvertedForInt

		users[id_userConvertedForInt] = u
	}
	defer db.Close()
	return users
}

func NewUser(user User) (int, error) {
	db := db.ConectBD()

	query := `insert into usuario (nome, senha, cargo) values($1, $2, $3) on conflict on constraint unique_name do nothing returning id_usuario`

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer db.Close()

	var userID int
	err = stmt.QueryRow(user.Name, user.Password, strconv.Itoa(user.Fk_office_id_office)).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}

func DeleteUser(user User) int {
	db := db.ConectBD()

	deletUser, err := db.Prepare("delete from user_system where id_user_system=$1")
	if err != nil {
		return 1
	}

	deletUser.Exec(user.Id_user)
	defer db.Close()
	return 0
}

func GetUser(id_user int) User {
	db := db.ConectBD()

	userBank, err := db.Query("select * from user_system where id_user_system=$1", id_user)
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	for userBank.Next() {
		var id_user, name, password, active, fk_office_id_office string

		err = userBank.Scan(&id_user, &name, &password, &active, &fk_office_id_office)
		if err != nil {
			panic(err.Error())
		}
		id_userConvertedForInt, err := strconv.Atoi(id_user)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		fk_office_id_officeConvertedForInt, err := strconv.Atoi(fk_office_id_office)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		activeConvertedForBool, err := strconv.ParseBool(active)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		user.Id_user = id_userConvertedForInt
		user.Name = name
		user.Password = password
		user.Active = activeConvertedForBool
		user.Fk_office_id_office = fk_office_id_officeConvertedForInt
	}
	defer db.Close()
	return user
}

func UpdateUser(user User) int {
	db := db.ConectBD()

	UpdateUser, err := db.Prepare("update user_system set name_user=$1, secret=$2, status=$3, fk_office_id_office=$4 where id_user_system=$5")
	if err != nil {
		return 1
	}

	UpdateUser.Exec(user.Name, user.Password, user.Active, user.Fk_office_id_office, user.Id_user)
	defer db.Close()
	return 0
}

func GetUserByName(name_user string) User {
	db := db.ConectBD()

	// userBank, err := db.Query("select * from user_system where name_user=$1", name_user)
	userBank, err := db.Query("select id_usuario, nome, senha, ativo, cargo from usuario where nome=$1", name_user)
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	for userBank.Next() {
		var id_user, name, password, active, fk_office_id_office string

		err = userBank.Scan(&id_user, &name, &password, &active, &fk_office_id_office)
		if err != nil {
			panic(err.Error())
		}
		id_userConvertedForInt, err := strconv.Atoi(id_user)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		fk_office_id_officeConvertedForInt, err := strconv.Atoi(fk_office_id_office)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		activeConvertedForBool, err := strconv.ParseBool(active)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		user.Id_user = id_userConvertedForInt
		user.Name = name
		user.Password = password
		user.Active = activeConvertedForBool
		user.Fk_office_id_office = fk_office_id_officeConvertedForInt
	}
	defer db.Close()
	return user
}

//Preparar chama os métodos de validar e formatar o usuário recebido
func (user *User) Preparar() error {
	if err := user.validar(); err != nil {
		return err
	}

	user.formatar()
	return nil
}

func (user *User) formatar() {
	user.Name = strings.TrimSpace(user.Name)
}

func (user *User) validar() error {
	if user.Name == "" {
		return errors.New("o nome do usuário é obrigatório e não pode estar em branco")
	}
	return nil
}
