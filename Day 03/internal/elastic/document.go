package elastic

import (
	"context"
	"encoding/json"
	"github.com/BatoBudaev/Go-Bootcamp/internal/model"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func IndexDocuments(es *elasticsearch.Client, indexName string, restaurants []model.Restaurant) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,
		Client:        es,
		NumWorkers:    4,
		FlushBytes:    5e6,
		FlushInterval: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Ошибка создания BulkIndexer: %s", err)
	}

	for _, restaurant := range restaurants {
		data, err := json.Marshal(restaurant)
		if err != nil {
			log.Fatalf("Ошибка преобразования в JSON: %s", err)
		}

		err = bi.Add(context.Background(), esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: strconv.Itoa(restaurant.ID),
			Body:       strings.NewReader(string(data)),
		})
		if err != nil {
			log.Fatalf("Ошибка добавления документа: %s", err)
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Ошибка закрытия BulkIndexer: %s", err)
	}

	biStats := bi.Stats()
	log.Printf("Успешно индексировано %d документов", biStats.NumFlushed)
}
