package handler

import (
	"encoding/json"
	"net/http"
	"test/repository"
)

func Set(db *repository.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var cell repository.Cell
		json.NewDecoder(r.Body).Decode(&cell)

		key := r.URL.Query().Get("key")
		if key == "" {
			respondWithError(w, http.StatusBadRequest, "request key not specified")
			return
		}
		if cell.Value == "" {
			respondWithError(w, http.StatusBadRequest, "request value not specified")
			return
		}
		if cell.TimeLife == 0 {
			respondWithError(w, http.StatusBadRequest, "request time not specified")
			return
		}
		defer r.Body.Close()
		//db.Set(key, value, time)
		db.Set(key, cell.Value, cell.TimeLife)
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
	})
}
