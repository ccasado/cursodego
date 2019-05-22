package repo

import (
	"github.com/jmoiron/sqlx"
	/*
		"github.com/go-sql-driver/mysql" não é usado diretamente na aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL função que abre conexão com banco de dados
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return

}
