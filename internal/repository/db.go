package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgres() *sql.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com PostgreSQL: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar no PostgreSQL: %v", err)
	}

	log.Println("Conexão com PostgreSQL estabelecida com sucesso!")
	return db
}
