package main

import (
	"fmt"
	"github.com/BatoBudaev/Go-Bootcamp/internal/elastic"
	"github.com/BatoBudaev/Go-Bootcamp/internal/ex02"
	"log"
	"net/http"
)

func main() {
	es := elastic.InitClient()

	store := &elastic.ElasticsearchStore{Client: es}
	http.HandleFunc("/", ex02.HandleApiRequest(store))

	fmt.Println("Сервер запущен на порту 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
