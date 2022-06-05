package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Transaction struct {
	Id_transaction int
	Maturity_date  string
	Issue_date     string
	Invoice        string
}

type Transacao struct {
	IDTransacao     string  `json:"id,omitempty"`
	DT_HR_REG       string  `json:"dt-hr-reg,omitempty"`
	Data_Emissao    string  `json:"emissao,omitempty"`
	Data_Vencimento string  `json:"vencimento,omitempty"`
	NF_Servico      string  `json:"nota,omitempty"`
	Valor           float32 `json:"valor,omitempty"`
	Banco           string  `json:"banco,omitempty"`
	Cliente         string  `json:"cliente,omitempty"`
	Liquidado       string  `json:"liquidado,omitempty"`
	Liquidado_EM    string  `json:"liquidado-em,omitempty"`
	Cancelado       string  `json:"cancelado,omitempty"`
	Cancelado_EM    string  `json:"cancelado-em,omitempty"`
}

func AllTransaction() []Transacao {
	//Conecta com Postgres
	db := db.ConectBD()
	defer db.Close()

	//Executa uma  query no postgres
	rows, err := db.Query(`select 
								t.id_transacao,
								to_char(t.dt_hr_reg, 'DD/MM/YYYY HH24:MI') as dt_hr_reg,
								to_char(t.data_emissao, 'DD/MM/YYYY') as data_emissao,
								to_char(t.data_vencimento,'DD/MM/YYYY') as data_vencimento,
								t.nf_servico,
								t.valor,
								b.nome as banco,
								c.nome as cliente,
								t.liquidado::text as liquidado,
								to_char(coalesce(t.liquidado_em, '1111-11-11')::date, 'DD/MM/YYYY') as liquidado_em,
								t.cancelado::text as cancelado,
								to_char(coalesce(t.cancelado_em, '1111-11-11')::date, 'DD/MM/YYYY') as cancelado_em
						from transacao t
							inner join banco b using (id_banco)
							inner join cliente c using (id_cliente)
					order by t.id_transacao;`)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var trasacoes []Transacao

	for rows.Next() {
		var transacao Transacao

		if err := rows.Scan(
			&transacao.IDTransacao,
			&transacao.DT_HR_REG,
			&transacao.Data_Emissao,
			&transacao.Data_Vencimento,
			&transacao.NF_Servico,
			&transacao.Valor,
			&transacao.Banco,
			&transacao.Cliente,
			&transacao.Liquidado,
			&transacao.Liquidado_EM,
			&transacao.Cancelado,
			&transacao.Cancelado_EM,
		); err != nil {
			fmt.Println(err)
		}
		trasacoes = append(trasacoes, transacao)
	}
	return trasacoes
}

func AllTransactionForPeriod(dtini, dtfin string) ([]Transacao, error) {
	//Conecta com Postgres
	db := db.ConectBD()
	defer db.Close()

	//Executa uma  query no postgres
	rows, err := db.Query(`select 
								t.id_transacao,
								to_char(t.dt_hr_reg, 'DD/MM/YYYY HH24:MI') as dt_hr_reg,
								to_char(t.data_emissao, 'DD/MM/YYYY') as data_emissao,
								to_char(t.data_vencimento,'DD/MM/YYYY') as data_vencimento,
								t.nf_servico,
								t.valor,
								b.nome as banco,
								c.nome as cliente,
								t.liquidado::text as liquidado,
								to_char(coalesce(t.liquidado_em, '1111-11-11')::date, 'DD/MM/YYYY') as liquidado_em,
								t.cancelado::text as cancelado,
								to_char(coalesce(t.cancelado_em, '1111-11-11')::date, 'DD/MM/YYYY') as cancelado_em
						from transacao t
							inner join banco b using (id_banco)
							inner join cliente c using (id_cliente)
					where 
						t.dt_hr_reg::date between $1 and $2
					order by t.id_transacao;`, dtini, dtfin)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var trasacoes []Transacao

	for rows.Next() {
		var transacao Transacao

		if err := rows.Scan(
			&transacao.IDTransacao,
			&transacao.DT_HR_REG,
			&transacao.Data_Emissao,
			&transacao.Data_Vencimento,
			&transacao.NF_Servico,
			&transacao.Valor,
			&transacao.Banco,
			&transacao.Cliente,
			&transacao.Liquidado,
			&transacao.Liquidado_EM,
			&transacao.Cancelado,
			&transacao.Cancelado_EM,
		); err != nil {
			fmt.Println(err)
		}
		trasacoes = append(trasacoes, transacao)
	}
	return trasacoes, nil
}

func NewTransaction(t Transacao) error {
	db := db.ConectBD()
	defer db.Close()

	insertDataTransaction, err := db.Prepare(`insert into transacao(data_emissao, data_vencimento, nf_servico, id_banco, id_cliente, valor) 
		values($1, $2, $3, $4, $5, $6);`)
	// insertDataTransaction, err := db.Prepare(`insert into transacao(data_emissao, data_vencimento, nf_servico, id_banco, id_cliente, valor)
	// 	values($1, $2, $3, $4, $5, $6) on conflict on constraint unq_transacao do nothing`)

	if err != nil {
		return err
	}

	_, err = insertDataTransaction.Exec(t.Data_Emissao, t.Data_Vencimento, t.NF_Servico, t.Banco, t.Cliente, t.Valor)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTransaction(transaction Transaction) int {
	db := db.ConectBD()

	deletTransaction, err := db.Prepare("delete from Transaction where id_transaction=$1")
	if err != nil {
		return 1
	}

	_, err = deletTransaction.Exec(transaction.Id_transaction)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetTransaction(id_transaction int) Transaction {
	db := db.ConectBD()

	transactions, err := db.Query("select * from Transaction where id_transaction=$1", id_transaction)
	if err != nil {
		panic(err.Error())
	}

	transaction := Transaction{}
	for transactions.Next() {
		var id_transaction, maturity_date, invoice, issue_date string

		err = transactions.Scan(&id_transaction, &maturity_date, &invoice, &issue_date)
		if err != nil {
			panic(err.Error())
		}

		id_transactionConvertedForInt, err := strconv.Atoi(id_transaction)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		transaction.Id_transaction = id_transactionConvertedForInt
		transaction.Maturity_date = maturity_date
		transaction.Invoice = invoice
		transaction.Issue_date = issue_date
	}
	defer db.Close()
	return transaction
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateTransaction(transaction Transaction) int {
	db := db.ConectBD()

	UpdateBanks, err := db.Prepare("update Transaction set maturity_date = $1, invoice=$2, issue_date = $3 where id_transaction=$4")
	if err != nil {
		fmt.Println("Error in updating the bank table: ", err)
		return 1
	}

	UpdateBanks.Exec(transaction.Maturity_date, transaction.Invoice, transaction.Issue_date, transaction.Id_transaction)
	defer db.Close()
	return 0
}

func BaixarTitulo(id_transacao uint64) error {
	db := db.ConectBD()
	defer db.Close()

	prepare, err := db.Prepare("update transacao set liquidado = true, liquidado_em = now() where id_transacao = $1")
	if err != nil {
		return fmt.Errorf("erro ao tentar baixar o título %d\n%s", id_transacao, err)
	}

	_, err = prepare.Exec(id_transacao)
	if err != nil {
		return fmt.Errorf("erro ao tentar baixar o título %d\n%s", id_transacao, err)
	}

	return nil
}

func CancelarTitulo(id_transacao uint64) error {
	db := db.ConectBD()
	defer db.Close()

	prepare, err := db.Prepare("update transacao set cancelado = true, cancelado_em = now() where id_transacao = $1")
	if err != nil {
		return fmt.Errorf("erro ao tentar cancelar o título %d\n%s", id_transacao, err)
	}

	_, err = prepare.Exec(id_transacao)
	if err != nil {
		return fmt.Errorf("erro ao tentar cancelar o título %d\n%s", id_transacao, err)
	}

	return nil
}

func ExtornarTitulo(id_transacao uint64) error {
	db := db.ConectBD()
	defer db.Close()

	prepare, err := db.Prepare(`update transacao 
		set extornado = true, 
		extornado_em = now(), 
		liquidado = false 
	where 
		id_transacao = $1`)

	if err != nil {
		return fmt.Errorf("erro ao tentar extornar o título %d\n%s", id_transacao, err)
	}

	_, err = prepare.Exec(id_transacao)
	if err != nil {
		return fmt.Errorf("erro ao tentar extornar o título %d\n%s", id_transacao, err)
	}

	return nil
}
