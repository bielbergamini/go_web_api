package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/bielbergamini/go_web_api/internal/handler"
	"github.com/bielbergamini/go_web_api/internal/repository"
)

func main() {
	// 1. Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis do sistema...")
	}

	// 2. Conectar ao banco
	db := repository.ConnectPostgres()
	defer db.Close()

	// 3. Criar tabela se não existir
	repository.CreateUsersTable(db)

	// 4. Rotas
	http.HandleFunc("/status", handler.StatusHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
