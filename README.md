# account-mgmt

account-mgmt is made up of an account transactions UI and and Rest API.

## Prerequisites
- docker compose
- node 12
- go 13

# Run
`docker-compose up --build`
or you can eventually run each application
- account-mgmt-front: `npm run local`
- account-mgmt-front: `go run main.go` 
 

### Use it
UI runs at `http://localhost:8080 `

##### API reference
Get account balance
```
curl --location --request GET 'http://localhost:3000/api/account-mgmt/account/balance' \
--header 'Content-Type: application/json'
```
List account transactions
```
curl --location --request GET 'http://localhost:3000/api/account-mgmt/account/transactions' \
--header 'Content-Type: application/json' 
```
Get transaction by id
```
curl --location --request GET 'http://localhost:3000/api/account-mgmt/account/transaction/:id' \
--header 'Content-Type: application/json' \
```
Save a transaction
```
curl --location --request POST 'http://localhost:3000/api/account-mgmt/account/transaction' \
--header 'Content-Type: application/json' \
--data-raw '{
	"type": "credit|debit",
	"amount": 5000
}'
```