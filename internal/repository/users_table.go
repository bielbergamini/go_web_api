package repository

import (
	"database/sql"
	"log"
)

func CreateUsersTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela users: %v", err)
	}

	log.Println("Tabela 'users' verificada/criada com sucesso!")
}
