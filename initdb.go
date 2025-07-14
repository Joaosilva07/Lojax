package main

import (
	"log"
	db "lojaX/backend"
)

func InitializeDatabase() {
	err := db.InitDB("loja.db")
	if err != nil {
		log.Fatal("Erro ao inicializar o banco:", err)
	}
	log.Println("Banco de dados criado com sucesso!")
}
