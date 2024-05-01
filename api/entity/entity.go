package entity

type ValidateBankAccountReq struct {
	Account string
	Name    string
}

type ValidateBankAccountRes struct {
	Id          int    `json:"id"`
	Account     string `json:"account_no"`
	AccountName string `json:"account_name"`
	BankName    string `json:"bank_name"`
}

type TransferRecipients struct {
	DestinationAccount string `json:"destination_account"`
	BankName           string `json:"bank_name"`
	TransferAmount     int    `json:"transfer_amount"`
	RefNumber          string `json:"reference_number"`
}

type TransferReq struct {
	PaymentType        string               `json:"payment_type"`
	SourceAccount      string               `json:"source_account"`
	TransferRecipients []TransferRecipients `json:"transfer_recipients"`
	CallbackUrl        string               `json:"callback_url"`
}

type TransferRes struct {
	StatusCode        int    `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionStatus string `json:"transaction_status"`
	Currency          string `json:"currency"`
}

type UpdateTransferStatusReq struct {
	ReferenceNumber   string `json:"reference_num"`
	TransactionStatus string `json:"status"`
}

type UpdateTransferStatusRes struct {
	StatusCode      int    `json:"status_code"`
	ReferenceNumber string `json:"reference_num"`
	Message         string `json:"message"`
}
