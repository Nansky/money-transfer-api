package usecase

import (
	"context"
	"money-api-transfer/api/entity"
	"sync"
)

type ValidateAccountImplementor interface {
	ValidateAccount(ctx context.Context, req entity.ValidateBankAccountReq) (res entity.ValidateBankAccountRes, code int, err error)
}

type TransactionImplementor interface {
	BankTransfer(ctx context.Context, req entity.TestBankTransferReq, wg *sync.WaitGroup)
	CallbackBankTransfer(ctx context.Context, req entity.UpdateTransferStatusReq) (entity.UpdateTransferStatusRes, error)
}

type Usecases struct {
	ValidateAccountUc ValidateAccountUsecase
	BankTransferUc    BankTransferUsecase
}
