package repository

import (
	"context"
	"money-api-transfer/api/entity"
)

type AccountRepositoryImplementor interface {
	GetAccountId(accNumber, bankCode string) (string, error)
	UpdateAccountBalance(sourceAcc, destAcc string, balance int) error
}

type TransactionRepositoryImplementor interface {
	GetTransactionAccount(refNum string) (string, string)
	CreateTransaction(ctx context.Context, req entity.TestBankTransferReq, transactionStatus int) error
	UpdateTransactionStatus(ctx context.Context, referenceNum string, status int) (string, string, int, error)
}

type Repository struct {
	AccountRepo     AccountRepository
	TransactionRepo TransactionRepository
}
