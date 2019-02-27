package handler

import (
	"net/http"
	"test/repository"
)

func Delete(db *repository.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			respondWithError(w, http.StatusBadRequest, "request key not specified")
			return
		}
		err := db.Delete(key)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "can't delete cell")
			return
		}
		respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
	})
}
