package usecase

import (
	"context"
	"encoding/json"
	"log"
	"money-api-transfer/api/commons"
	"money-api-transfer/api/entity"
	"money-api-transfer/api/repository"
	"net/http"
	"sync"
)

var mapStatus = map[string]int{
	"PENDING": 1,
	"SUCCESS": 2,
	"FAILED":  3,
}

type BankTransferUsecase struct {
	repo repository.Repository
}

func NewBankTransfertUsecase(rp repository.Repository) BankTransferUsecase {
	return BankTransferUsecase{
		repo: rp,
	}
}

func (tu BankTransferUsecase) BankTransfer(ctx context.Context, req entity.TestBankTransferReq, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(commons.BankTransferUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	res := entity.TestBankTransferRes{}
	resArr := []entity.TestBankTransferRes{}

	if err = json.NewDecoder(resp.Body).Decode(&resArr); err != nil {
		return
	}

	for _, response := range resArr {
		if response.RefNumber == req.RefNumber {
			res = response
		}
	}

	statusTrx := mapStatus[res.TransactionStatus]
	tu.repo.TransactionRepo.CreateTransaction(ctx, req, statusTrx)
}

func (tu BankTransferUsecase) CallbackBankTransfer(ctx context.Context, req entity.UpdateTransferStatusReq) (entity.UpdateTransferStatusRes, error) {
	// update status and get account balance
	sourceAcc, destinationAcc, balance, err := tu.repo.TransactionRepo.UpdateTransactionStatus(ctx, req.ReferenceNumber, mapStatus[req.TransactionStatus])
	if err != nil {
		log.Println("Data not found")
		return entity.UpdateTransferStatusRes{
			StatusCode: http.StatusBadRequest,
			Message:    "Data not found",
		}, err
	}

	err = tu.repo.AccountRepo.UpdateAccountBalance(sourceAcc, destinationAcc, balance)
	if err != nil {
		log.Println("Failed updating balance")
		return entity.UpdateTransferStatusRes{
			StatusCode: http.StatusNotFound,
			Message:    "Failed updating balance accounts",
		}, err
	}

	return entity.UpdateTransferStatusRes{
		StatusCode:      http.StatusAccepted,
		ReferenceNumber: req.ReferenceNumber,
		Message:         "Success update transaction",
	}, nil

}
