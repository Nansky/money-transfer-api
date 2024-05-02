# money-transfer-api

Simple Money API Transfer using Go. 

### What's this all about
Money Transfer APIs using Golang programming consist of 3 endpoints :
1. `/api/v1/account_validation` : Account Validation URL endpoint that used to validate account number and name of the bank account owner.
2. `/api/v1/transfer` : Transfer/Disbursement URL endpoint that used to transfer money to the destination account. To transfer the money, you can create a mock endpoint similar to point 1 that acts as a bank.
3. `/api/v1/transfer/update_status` : Callback URL endpoint that used to receive callback status of transfer from the bank. (Note that this will trigger from our side for testing only, assuming sender already get callback URL when doing transfer).
4. For authentication, i'm using hardcoded token that is stored in environment variables for simplicity and will be checked accordingly for each incoming request. (I know this will trigger vulnerability secret leak issue but in real project hardcoded custom token will not be applied).

### How the code works
In this simulation, i'm using `https://mockapi.io` to mock Bank API named `testbank`. At first i want to create 4 bank API but in free version it only allows me to create 2 mock resources and limited customization. `testbank` will act as a bank to do check validation and process the transfer for following point 1 and 2
1. Account Validation handler will validate account to destination bank owner. it uses http `GET` method and accpet 2 url parameters `bank_name` and `account`. Instead of send request directly to mock bank API, I do the tricky part when get  response from `mock testbank`, the response is in array of detail bank account (due to mockAPI free version), so first i get the account `ID` from accounts that stored in local Database (the id from the response same as id in mock bank response), the ID is used as an URL Param to get the account detail response.
2. Transfer endpoint handler will process all transfer from source account to destination account. Destination account to be transfered can be multiple accounts (depending on the request), so the process will be concurrent for each transfer and get the response from `mock testbank`. For simplicity, i do not check the status of every concurrent process, for example if there are 3 transfer process (we hope all transaction are OK), we do not really know how much transactions that is success. The workaround for this is we can using channel or errgroup to ensure that we can track every concurrent process.
3. Callback handler  check and update the transfer status and balance that sent from the bank based on transaction id or reference number in this simulation (in this case we send it manually as a bank). 

### Payloads Request
Account Validation
```cURL
curl --location 'http://localhost:6969/api/v1/account_validation?account=034101056895506&bank_name=testbank' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer VlQtc2VydmVyLUNwbzAza1lET2MwY05VS2d0NmhuTGtLZzo' \
--data ''
```

Money Transfer 
```cURL
curl --location 'http://localhost:6969/api/v1/transfer' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer VlQtc2VydmVyLUNwbzAza1lET2MwY05VS2d0NmhuTGtLZzo' \
--data '{
    "payment_type" : "bank_transfer",
    "source_account" : "0000412421",
    "transfer_recipients" : [
        {
            "destination_account" : "111001110963",
            "bank_name" : "testbank",
            "transfer_amount" : 10000,
            "reference_number" : "598379835798"
        },
        {
            "destination_account" : "034101056895506",
            "bank_name" : "testbank",
            "transfer_amount" : 20000,
            "reference_number" : "169023111100"
        }
    ],
    "callback_url" : "http://localhost:6060/api/v1/update_status"
}'
```


Update status Endpoint (Assuming this sends from the bank)
```cURL
curl --location --request PATCH 'http://localhost:6969/api/v1/transfer/update_status' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer VlQtc2VydmVyLUNwbzAza1lET2MwY05VS2d0NmhuTGtLZzo' \
--data '{
    
    "reference_num" : "598379835798",
    "status" : "SUCCESS"
        
}'
```

### Sample Response
- Account Validation
```json
{
    "id": 1,
    "account_no": "034101056895506",
    "account_name": "Si Budi",
    "bank_name": "testbank"
}
```
- Transfer
```json
{
    "status_code": 201,
    "status_message": "Success, Bank Transfer transaction is created",
    "transaction_status": "PENDING",
    "currency": "IDR"
}
```
- Update Status (Callback)
```json
{
    "status_code": 202,
    "reference_num": "169023111100",
    "message": "Success update transaction"
}
```

### Database 
I am using database migration to make it simple that can be run using make command. There are 3 tables that i'm using :
- accounts : stored account number, name and balance (for internal data purpose)
- order_status : stored list of order status 
- orders : stored order transaction for transfer and transfer status

### Run this in your local environment using Docker Compose
1. Clone this repository
2. in your CLI, Execute `make run` with your docker installed
```sh
make run
```
3. Then execute database migration using `make migrate-up` command. this will create all tables stored under `migration/sql` directory
```sh
make migrate-up
```
4. All set and you can test the money-transfer API.