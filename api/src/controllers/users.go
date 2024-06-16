package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário criado"))
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
