package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	fmt.Println("rodando api na porta :3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
