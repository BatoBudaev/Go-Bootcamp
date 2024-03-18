package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/BatoBudaev/Go-Bootcamp/internal/model"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"net/http"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]model.Restaurant, int, error)
	SearchClosestRestaurants(lat float64, lon float64, size int) ([]model.Restaurant, error)
}

type ElasticsearchStore struct {
	Client *elasticsearch.Client
}

func (es *ElasticsearchStore) GetPlaces(limit int, offset int) ([]model.Restaurant, int, error) {
	query := map[string]interface{}{
		"size": limit,
		"from": offset,
		"sort": []map[string]interface{}{
			{
				"id": map[string]string{
					"order": "asc",
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, 0, fmt.Errorf("ошибка при кодировании запроса: %v", err)
	}

	res, err := es.Client.Search(
		es.Client.Search.WithContext(context.Background()),
		es.Client.Search.WithIndex("places"),
		es.Client.Search.WithBody(&buf),
		es.Client.Search.WithSize(100),
		es.Client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка выполнения поиска: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("неожиданный код состояния: %s", res.Status())
	}

	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source model.Restaurant `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	var restaurants []model.Restaurant
	for _, hit := range result.Hits.Hits {
		restaurants = append(restaurants, hit.Source)
	}

	return restaurants, result.Hits.Total.Value, nil
}

func (es *ElasticsearchStore) SearchClosestRestaurants(lat float64, lon float64, size int) ([]model.Restaurant, error) {
	query := map[string]interface{}{
		"size": size,
		"sort": []map[string]interface{}{
			{
				"_geo_distance": map[string]interface{}{
					"location": map[string]interface{}{
						"lat": lat,
						"lon": lon,
					},
					"order":           "asc",
					"unit":            "km",
					"mode":            "min",
					"distance_type":   "arc",
					"ignore_unmapped": true,
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("ошибка при кодировании запроса: %v", err)
	}

	res, err := es.Client.Search(
		es.Client.Search.WithContext(context.Background()),
		es.Client.Search.WithIndex("places"),
		es.Client.Search.WithBody(&buf),
		es.Client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения поиска: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.IsError() {
		return nil, fmt.Errorf("неожиданный код состояния: %s", res.Status())
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source model.Restaurant `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	var restaurants []model.Restaurant
	for _, hit := range result.Hits.Hits {
		restaurants = append(restaurants, hit.Source)
	}

	return restaurants, nil
}
