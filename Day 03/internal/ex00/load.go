package ex00

import (
	"encoding/csv"
	"github.com/BatoBudaev/Go-Bootcamp/internal/model"
	"io"
	"log"
	"os"
	"strconv"
)

func LoadRestaurants() []model.Restaurant {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	if _, err := reader.Read(); err != nil {
		log.Fatalf("Ошибка чтения заголовка: %v", err)
	}

	var restaurants []model.Restaurant

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Ошибка чтения строки: %v", err)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("Ошибка при преобразовании ID: %v", err)
		}

		lat, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatalf("Ошибка при преобразовании широты: %v", err)
		}

		lon, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatalf("Ошибка при преобразовании долготы: %v", err)
		}

		restaurant := model.Restaurant{
			ID:      id + 1,
			Name:    record[1],
			Address: record[2],
			Phone:   record[3],
			Location: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}{Lat: lat, Lon: lon},
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants
}
