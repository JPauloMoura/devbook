package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("failed to read body", err)
	}

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal("failed to unmarshal user", err)
	}

	db := database.ConnectDb()
	id, err := repositories.NewUserRepository(db).Create(user)
	if err != nil {
		log.Fatal("failed to create user", err)
	}

	w.Write([]byte(fmt.Sprintf("created: %d", id)))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário encontrado"))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuários listados"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário atualizado"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário deletado"))
}
