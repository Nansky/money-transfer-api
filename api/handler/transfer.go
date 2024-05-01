package handler

import (
	"encoding/json"
	"money-api-transfer/api/commons"
	"money-api-transfer/api/entity"
	"money-api-transfer/api/usecase"
	"net/http"
	"strings"
	"sync"
)

type TransferHandler struct {
	UseCase   usecase.Usecases
	AuthToken string
}

func (th TransferHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	var trfReq entity.TransferReq

	err := json.NewDecoder(r.Body).Decode(&trfReq)
	if err != nil {
		e := commons.ErrorResponse{
			Message: "Invalid Request for Transfer",
			Error:   err.Error(),
		}

		e.SetJSONError(w, http.StatusBadRequest)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(trfReq.TransferRecipients))

	for _, recipient := range trfReq.TransferRecipients {
		testBankReq := entity.TestBankTransferReq{
			PaymentType:        trfReq.PaymentType,
			SourceAccount:      trfReq.SourceAccount,
			GrossAmount:        recipient.TransferAmount,
			RefNumber:          recipient.RefNumber,
			CallbackUrl:        commons.CallbackUrl,
			DestinationAccount: recipient.DestinationAccount,
			BankName:           recipient.BankName,
		}

		go th.UseCase.BankTransferUc.BankTransfer(r.Context(), testBankReq, wg)
	}

	wg.Wait() // wait until all transfer complete

	resTrf := entity.TransferRes{
		StatusCode:        http.StatusCreated,
		StatusMessage:     "Success, Bank Transfer transaction is created",
		TransactionStatus: "PENDING",
		Currency:          "IDR",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resTrf.StatusCode)
	json.NewEncoder(w).Encode(resTrf)
}
