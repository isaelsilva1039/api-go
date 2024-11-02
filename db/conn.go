package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error

	for attempts := 1; attempts <= 5; attempts++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Tentativa %d: Erro ao abrir conexão com o banco de dados: %v", attempts, err)
		} else {
			err = db.Ping()
			if err == nil {
				fmt.Println("Conectado ao banco de dados!")
				return db, nil
			}
		}

		log.Printf("Tentativa %d falhou. Retentando em 5 segundos...", attempts)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("não foi possível conectar ao banco de dados: %w", err)
}
