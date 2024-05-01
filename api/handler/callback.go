package handler

import (
	"encoding/json"
	"money-api-transfer/api/commons"
	"money-api-transfer/api/entity"
	"money-api-transfer/api/usecase"
	"net/http"
	"strings"
)

type TransactionCallbackHandler struct {
	UseCase   usecase.Usecases
	AuthToken string
}

func (th TransactionCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	reqToken := strings.TrimPrefix(authHeader, "Bearer ")
	if reqToken != th.AuthToken {
		e := commons.ErrorResponse{
			Message: "Invalid Token",
			Error:   "Invalid Auth Token",
		}

		e.SetJSONError(w, http.StatusUnauthorized)
		return
	}

	trfReq := entity.UpdateTransferStatusReq{}

	err := json.NewDecoder(r.Body).Decode(&trfReq)
	if err != nil {
		e := commons.ErrorResponse{
			Message: "Invalid Request for Transfer",
			Error:   err.Error(),
		}

		e.SetJSONError(w, http.StatusBadRequest)
		return
	}

	res, err := th.UseCase.BankTransferUc.CallbackBankTransfer(r.Context(), trfReq)
	if err != nil {
		e := commons.ErrorResponse{
			Message: res.Message,
			Error:   err.Error(),
		}

		e.SetJSONError(w, res.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(res)
}
