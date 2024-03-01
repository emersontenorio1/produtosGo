package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "cvdptfiymxxm90p3:i2cx9ee0f8lgr38y@tcp(y5s2h87f6ur56vae.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/c6anclt2pe0a493o"
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())

	}
	return db
}
