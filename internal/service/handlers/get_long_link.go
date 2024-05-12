package handlers

import (
    "strconv"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zepif/Test-service/internal/data"
)

type GetShortLinkResponse struct {
	FullURL  string `json:"full_url"`
	ShortURL string `json:"short_url"`
}

func GetShortLink(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
    id, err := strconv.Atoi(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := r.Context().Value(1).(data.URLStorage).New()
	/*if !ok {
		http.Error(w, "Failed to get data.URLStorage from context", http.StatusInternalServerError)
		return
	}*/

	linkQ := db.Link()
	fullURL, shortURL, err := linkQ.Get(int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Link not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := GetShortLinkResponse{FullURL: fullURL, ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
