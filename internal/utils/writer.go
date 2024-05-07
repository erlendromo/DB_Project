package utils

import (
	"DB_Project/internal/constants"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func setHeaders(w http.ResponseWriter, statuscode int) error {
	if statuscode == http.StatusNoContent {
		return errors.New("invalid statuscode -> use http.StatusNoContent directly")
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.APP_JSON)
	w.WriteHeader(statuscode)
	return nil
}

func JSON(w http.ResponseWriter, statuscode int, data any) {
	if err := setHeaders(w, statuscode); err != nil {
		log.Print(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}
