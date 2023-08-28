
# SERVICE USER AND API GATEWAY

This service provide authentication for application and API GATEWAY as service discovery for an other service.
## Features

- Authentication (Sign Up, Sign In and Sign Out)
    - Sign Up
    Sign Up will create new user and send email information to user. 
    - Sign In
    Sign In will response JWT as Access Token And Refresh Token.
    This service use redis to store access with specfic duration. The duration of access storage will increase when the client accesses the API. When access deleted from redis, client will be unauthenticated.
    - Sign Out
    Sign out will delete access from redis. 
- API Gateway
API Gateway used as a service discovery. So that client can access the other service from API.

## API SPECS

https://documenter.getpostman.com/view/3675744/2s9Xy2NXU3

## Run Locally

Clone the project

```bash
  git clone https://github.com/fernandojec/h_serviceuser
```

Migrate Database

```sql
    execute sql at directory docs/dbMigration/db_Migration.sql
```

Update setting in .env file


Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run cmd/api/main.go
```

## ENVIRONTMENT VARIABLE

```env
POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_DBNAME="HUser"
POSTGRES_USER="sa"
POSTGRES_PASSWORD="P@ssw0rd"
POSTGRES_SSLMODE="disable"

BASE_URL="http://localhost"
BASE_PORT=":8182"

EXPIRE_ACCESS_MIN=30    #DURATION OF JWT ACCESS EXPIRE
EXPIRE_REFRESH_MIN=60   #DURATION OF JWT REFRESH EXPIRE
AUTO_LOGOFF_MIN=10      #DURATION OF IDLE ACCESS


REDIS_HOST="localhost"
REDIS_PORT="6379"
REDIS_PASSWORD=""

SERVICE_PARAMEDIC="localhost:8801"      #HOST OF SERVICE PARAMEDIC
SERVICE_APPOINTMENT="localhost:8801"    #HOST OF SERVICE APPOINTMENT
SERVICE_PATIENT="localhost:8801"        #HOST OF SERVICE PATIENT
SERVICE_SCHEDULE="localhost:8801"       #HOST OF SERVICE SCHEDULE
SERVICE_NOTIFICATION="localhost:8801"   #HOST OF SERVICE NOTIFICATION
SERVICE_HEALTHCARE="localhost:8801"     #HOST OF SERVICE HOSPITAL
```

## AUTHOR
- Fernando Riyo Junedy Simbolon (fernando.riyo@jec.co.id)
- Ricky Lai (ricky.lai@jec.co.id)