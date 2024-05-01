package repository

import (
	"database/sql"
	"log"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(dbCli *sql.DB) AccountRepository {
	return AccountRepository{
		db: dbCli,
	}
}

func (ar AccountRepository) GetAccountId(accNumber, bankCode string) (string, error) {
	var id string

	if err := ar.db.QueryRow("select id from accounts where account_number = $1 and bank_code = $2", accNumber, bankCode).Scan(&id); err != nil {
		log.Println(err)
		return "", sql.ErrNoRows
	}

	return id, nil
}

func (ar AccountRepository) UpdateAccountBalance(sourceAcc, destAcc string, balance int) error {
	// Get a Tx for making transaction requests.
	tx, err := ar.db.Begin()
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// deduct source account
	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 where account_number = $2", balance, sourceAcc)
	if err != nil {
		log.Println("Failed to deduct account balance")
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 where account_number = $2", balance, destAcc)
	if err != nil {
		log.Println("Failed to update account balance")
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println("Failed commit transaction")
		return err
	}

	return nil
}
