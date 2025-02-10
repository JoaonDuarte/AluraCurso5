package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func ConectaDB() {
	stringDeConexao := "host=localhost user=dicomvix password=system98 dbname=clinux_ris_old_old port=8986 sslmode=disable"
	DB, err = sql.Open("postgres", stringDeConexao)
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados.")
	}

	if err := DB.Ping(); err == nil {
		fmt.Println("Conectado")
	}

}
