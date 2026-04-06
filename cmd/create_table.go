package main

import (
	"log"

	"github.com/GuilhermeKAC/go-user-api/internal/config"
	"github.com/GuilhermeKAC/go-user-api/internal/database"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.CreateUsersTable(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabela users criada com sucesso!")
}
