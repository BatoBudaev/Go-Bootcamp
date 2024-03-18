package elastic

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"log"
	"strings"
)

func CreateIndex(es *elasticsearch.Client, indexName string) {
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body: strings.NewReader(`{
			"mappings": {
				"properties": {
					"name": {
						"type": "text"
					},
					"address": {
						"type": "text"
					},
					"phone": {
						"type": "text"
					},
					"location": {
						"type": "geo_point"
					}
				}
			}
		}`),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Ошибка создания индекса: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.IsError() {
		log.Printf("Ошибка создания индекса: %s", res.String())
	} else {
		log.Printf("Индекс %s создан успешно", indexName)
	}
}
