package main

import (
	"fmt"
	"go_server/api/routes"
	"go_server/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	err = db.Exec("SELECT 1").Error
	if err != nil {
		log.Fatal("Erro ao testar a conexão com o banco de dados:", err)
	}

	r := routes.Setup(db)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
