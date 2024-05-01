# money-transfer-api

Simple Money API Transfer using Go

### What's this all about
Money Transfer APIs using Golang programming consist of 3 endpoints :
1. `/api/v1/account_validation` : Account Validation URL endpoint that used to validate account number and name of the bank account owner.
2. `/api/v1/transfer` : Transfer/Disbursement URL endpoint that used to transfer money to the
destination account. To transfer the money, you can create a mock endpoint
similar to point 1 that acts as a bank.
3. `/api/v1/transfer/update_status` : Transfer/Disbursement Callback URL endpoint that used to receive callback
status of transfer from the bank.