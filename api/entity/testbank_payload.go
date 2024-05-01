package entity

type TestBankTransferReq struct {
	PaymentType        string `json:"payment_type"`
	SourceAccount      string `json:"source_account"`
	GrossAmount        int    `json:"gross_amount"`
	RefNumber          string `json:"reference_number"`
	CallbackUrl        string `json:"callback_url"`
	DestinationAccount string `json:"destination_account"`
	BankName           string `json:"bank_name"`
}

type TestBankTransferRes struct {
	PaymentType       string `json:"payment_type"`
	StatusMessage     string `json:"status_message"`
	Amount            string `json:"amount"`
	RefNumber         string `json:"reference_num"`
	TransactionStatus string `json:"transaction_status"`
	Currency          string `json:"currency"`
	BankName          string `json:"bank_name"`
}
