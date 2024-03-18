package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func InitClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Ошибка создания клиента: %s", err)
	}

	return es
}
