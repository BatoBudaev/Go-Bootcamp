package main

import (
	"context"
	"github.com/BatoBudaev/Go-Bootcamp/internal/elastic"
	"github.com/BatoBudaev/Go-Bootcamp/internal/ex00"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

func main() {
	es := elastic.InitClient()

	indexName := "places"

	res, err := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Ошибка проверки существования индекса: %s", err)
	}

	if res.StatusCode == 404 {
		elastic.CreateIndex(es, indexName)
	} else {
		log.Printf("Индекс %s уже существует", indexName)
	}

	restaurants := ex00.LoadRestaurants()

	elastic.IndexDocuments(es, indexName, restaurants)
}
