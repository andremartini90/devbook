package banco

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_"github.com/go-sql-driver/mysql" //driver
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		fmt.Println("Ocorreu um erro ao conectar ao banco de dados")
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
		

}
