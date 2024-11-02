package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

// ConnectDB cria uma conexão com o banco de dados usando GORM e retorna um ponteiro para *gorm.DB
func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *gorm.DB
	var err error

	// Tentativa de conexão com reconexão automática
	for attempts := 1; attempts <= 5; attempts++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// Verifica a conexão com o banco de dados
			sqlDB, _ := db.DB()
			err = sqlDB.Ping()
			if err == nil {
				fmt.Println("Conectado ao banco de dados!")
				return db, nil
			}
		}

		log.Printf("Tentativa %d de conexão ao banco de dados falhou: %v. Retentando em 5 segundos...", attempts, err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("não foi possível conectar ao banco de dados: %w", err)
}
