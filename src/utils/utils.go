// #

package utils

import (
	"encoding/json"
	"models"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, status int, error models.Error ){

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}