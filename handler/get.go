package handler

import (
	"encoding/json"
	"net/http"
	"test/repository"
)

func Get(db *repository.Storage) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		key := r.URL.Query().Get("key")
		if key == "" {
			respondWithError(w, http.StatusBadRequest, "request key not specified")
			return
		}
		value, err := db.Get(key)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "can't get cell by key")
			return
		}
		json.NewEncoder(w).Encode(value)
	})
}
