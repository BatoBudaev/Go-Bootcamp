package ex03

import (
	"encoding/json"
	"github.com/BatoBudaev/Go-Bootcamp/internal/elastic"
	"net/http"
	"strconv"
)

func HandleRecommendationRequest(store elastic.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")

		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			http.Error(w, `{"Неверное значение 'lat': '`+latStr+`"}`, http.StatusBadRequest)
			return
		}

		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			http.Error(w, `{"Неверное значение 'lon': '`+lonStr+`"}`, http.StatusBadRequest)
			return
		}

		restaurants, err := store.SearchClosestRestaurants(lat, lon, 3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseData := map[string]interface{}{
			"name":   "Recommendation",
			"places": restaurants,
		}

		w.Header().Set("Content-Type", "application/json")

		jsonResponse, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)
	}
}
