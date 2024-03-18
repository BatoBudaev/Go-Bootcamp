package main

import (
	"fmt"
	"github.com/BatoBudaev/Go-Bootcamp/internal/ex04"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/get_token", ex04.GetTokenHandler)
	http.HandleFunc("/api/recommend", ex04.RecommendHandler)

	fmt.Println("Сервер запущен на порту 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
