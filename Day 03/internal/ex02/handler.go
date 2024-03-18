package ex02

import (
	"encoding/json"
	"github.com/BatoBudaev/Go-Bootcamp/internal/elastic"
	"net/http"
	"strconv"
)

func HandleApiRequest(store elastic.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		if pageStr == "" {
			pageStr = "1"
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			http.Error(w, "Неверное значение 'page': '"+pageStr+"'", http.StatusBadRequest)
			return
		}

		limit := 100
		offset := (page - 1) * limit

		restaurants, total, err := store.GetPlaces(limit, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var prevPage, nextPage, lastPage int
		if total > limit {
			lastPage = total / limit
			if total%limit != 0 {
				lastPage++
			}
		}

		if page > 1 {
			prevPage = page - 1
		}
		if page < lastPage {
			nextPage = page + 1
		}

		responseData := map[string]interface{}{
			"name":      "Places",
			"total":     total,
			"places":    restaurants,
			"prev_page": prevPage,
			"next_page": nextPage,
			"last_page": lastPage,
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
