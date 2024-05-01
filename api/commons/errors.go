package commons

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (e ErrorResponse) SetJSONError(w http.ResponseWriter, errCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(e)
}

func MapErrorMessage(errStr string) string {
	mapErr := make(map[string]string)
	mapErr[SqlRowsNotFound] = "Error Rows not found in db"

	return mapErr[errStr]
}
