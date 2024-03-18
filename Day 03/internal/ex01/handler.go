package ex01

import (
	"github.com/BatoBudaev/Go-Bootcamp/internal/elastic"
	"github.com/BatoBudaev/Go-Bootcamp/internal/model"
	"html/template"
	"net/http"
	"strconv"
)

func HandleRequest(store elastic.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		if pageStr == "" {
			pageStr = "1"
		}

		page, err := strconv.Atoi(pageStr)
		limit := 100
		offset := (page - 1) * limit

		places, total, err := store.GetPlaces(limit, offset)
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

		if err != nil || page <= 0 || page > lastPage {
			http.Error(w, "Неверное значение 'page': '"+pageStr+"'", http.StatusBadRequest)
			return
		}

		if page > 1 {
			prevPage = page - 1
		}
		if page < lastPage {
			nextPage = page + 1
		}

		tmpl, err := template.ParseFiles("cmd/ex01/places.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Places   []model.Restaurant
			Total    int
			Page     int
			PrevPage int
			NextPage int
			LastPage int
		}{
			Places:   places,
			Total:    total,
			Page:     page,
			PrevPage: prevPage,
			NextPage: nextPage,
			LastPage: lastPage,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
