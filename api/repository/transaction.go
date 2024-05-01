package repository

import (
	"context"
	"database/sql"
	"log"
	"money-api-transfer/api/entity"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(dbCli *sql.DB) TransactionRepository {
	return TransactionRepository{
		db: dbCli,
	}
}

func (tr TransactionRepository) CreateTransaction(ctx context.Context, req entity.TestBankTransferReq, transactionStatus int) error {
	_, err := tr.db.ExecContext(ctx, "INSERT INTO orders (payment_type, reference_num, source_account, destination_account, gross_amount, order_status_id) VALUES($1, $2, $3, $4, $5, $6)",
		req.PaymentType, req.RefNumber, req.SourceAccount, req.DestinationAccount, req.GrossAmount, transactionStatus)
	if err != nil {
		log.Fatal("Error Create Transaction in DB :", err.Error())
		return err
	}

	return nil
}

func (tr TransactionRepository) UpdateTransactionStatus(ctx context.Context, referenceNum string, status int) (string, string, int, error) {
	var (
		destAcc, sourceAcc string
		balance            int
	)

	err := tr.db.QueryRow("UPDATE orders SET order_status_id = $1 where reference_num = $2 AND order_status_id = 1  RETURNING source_account, destination_account, gross_amount", status, referenceNum).
		Scan(&sourceAcc, &destAcc, &balance)
	if err != nil {
		log.Println("Error Update Transaction in DB :", err.Error())
		return "", "", 0, err
	}

	log.Println(destAcc, sourceAcc, balance)
	return destAcc, sourceAcc, balance, nil
}
