package main

import (
	"log"
	"net/http"

	"github.com/GuilhermeKAC/go-user-api/internal/config"
	"github.com/GuilhermeKAC/go-user-api/internal/handlers"
	"github.com/GuilhermeKAC/go-user-api/internal/repository"
)

func main() {
	// Conectar ao banco
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar repository e handler
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// Rotas
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.HandleFunc("GET /users", userHandler.GetAllUsers)
	http.HandleFunc("GET /users/", userHandler.GetUser)
	http.HandleFunc("PUT /users/", userHandler.UpdateUser)
	http.HandleFunc("DELETE /users/", userHandler.DeleteUser)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
