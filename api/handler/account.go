package handler

import (
	"encoding/json"
	"money-api-transfer/api/commons"
	"money-api-transfer/api/entity"
	"money-api-transfer/api/usecase"
	"net/http"
	"strings"
)

type AccountValidationHandler struct {
	UseCase   usecase.Usecases
	AuthToken string
}

func (av AccountValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	reqToken := strings.TrimPrefix(authHeader, "Bearer ")
	if reqToken != av.AuthToken {
		e := commons.ErrorResponse{
			Message: "Invalid Token",
			Error:   "Invalid Auth Token",
		}

		e.SetJSONError(w, http.StatusUnauthorized)
		return
	}

	if r.URL.Query().Get("account") == "" || r.URL.Query().Get("bank_name") == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	reqValidate := entity.ValidateBankAccountReq{
		Name:    r.URL.Query().Get("bank_name"),
		Account: r.URL.Query().Get("account"),
	}

	resValidate, statusCode, err := av.UseCase.ValidateAccountUc.ValidateAccount(r.Context(), reqValidate)
	if err != nil {
		e := commons.ErrorResponse{
			Message: err.Error(),
			Error:   err.Error(),
		}

		e.SetJSONError(w, statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resValidate)
}
