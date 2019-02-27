package handler

import (
	"net/http"
	"strconv"
	"test/repository"
)

func Set(db *repository.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		time, timeErr := strconv.ParseInt(r.URL.Query().Get("time"), 10, 64)
		if key == "" {
			respondWithError(w, http.StatusBadRequest, "request key not specified")
			return
		}
		if value == "" {
			respondWithError(w, http.StatusBadRequest, "request value not specified")
			return
		}
		if timeErr != nil {
			respondWithError(w, http.StatusBadRequest, "request time not specified")
			return
		}
		//Expiration - key lifetime
		db.Set(key, value, time)
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
	})
}
