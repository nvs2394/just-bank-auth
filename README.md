# just-bank-auth

### How to run local

> SERVER_ADDRESS=localhost SERVER_PORT=8001 DB_USER=just_bank_user DB_PASSWORD=password DB_NAME=just_bank_db go run main.go

If you want to run on release mode then put the cmd: `GIN_MODE=release`


![HLD](./high-level-design.png)

### Authentication flow

![HLD](./authen-flow.png)