# Wex Challenge

This project was developed for the selection process at Wex. It includes an API REST made in golang, containing 2 endpoints that allows you to store and retrieve transactions

**Note**: The application was delayed by 60 seconds to allow for database initialization.


## Getting Started
To get started with this project, clone the repository and install the dependencies…

### Running the Project
To run this project, use Docker Compose with the following command:

```sh
    docker-compose up -d
```

## Endpoints


### Store transactions
#### POST /transactions/insert

This endpoint allows you to store a transaction. 

#### Request

The request body should include the following fields:

- `description`: A string that describes the transaction.
- `date`: The date of the transaction in the format "YYYY/MM/DD".
- `amount`: The amount of the transaction as a decimal number.

Here is an example of a request:

```json
{
    "description": "foo",
    "date": "2023/12/02",
    "amount": 1.20
}
```
#### Response
The response will include the following fields:

- `Description`: A string that describes the transaction.
- `Date`: The date of the transaction in ISO 8601 format.
- `Amount`: The amount of the transaction as a decimal number.
- `Transaction_id`: A unique identifier for the transaction.


Here is an example of a response:

```json
{
    "Description": "foo",
    "Date": "2023-12-02T00:00:00Z",
    "Amount": 1.2,
    "Transaction_id": "b38bb6eb-74ef-4d46-a66f-2dfa9d23a7da"
}
```
Curl example: 
```
curl --location --request POST 'http://localhost:3000/transactions/insert' \
--header 'Content-Type: application/json' \
--data-raw '{
    "description": "foo",
    "date": "2023/12/02",
    "amount": 18.20003
}'
```

### Get Global transactions
#### GET /transactions/{transaction_id}/country/{country_name}
This endpoint retrieves the details of a specific transaction by its ID and the country name.

#### Path Parameters:
- `transaction_id`: The unique identifier of the transaction.
- `country_name`: The name of the country where the transaction took place.

#### Response:
The response is a JSON object that contains the following fields:

- `USDTransaction`: An object that contains details about the transaction.
  - `Description`: A string that describes the transaction.
  - `Date`: A string that represents the date and time of the transaction in ISO 8601 format (e.g., “2023-12-02T00:00:00Z”).
  - `Amount`: An integer that represents the amount of the transaction in USD.
  - `Transaction_id`: A string that represents the unique identifier of the transaction.
- `ExchangeRate`: An integer that represents the exchange rate from USD to the local currency of the country specified in the request.
- `ConvertedAmount`: An integer that represents the amount of the transaction in the local currency of the country specified in the request.

```json 
{
    "USDTransaction": {
        "Description": "foo",
        "Date": "2023-12-02T00:00:00Z",
        "Amount": 18.2,
        "Transaction_id": "127a2f70-c6d8-47b6-9bb1-4849929bddb4"
    },
    "ExchangeRate": 266,
    "ConvertedAmount": 4841.2
}
```

Curl example: 
```
curl --location --request GET 'http://localhost:3000/transactions/127a2f70-c6d8-47b6-9bb1-4849929bddb4/country/Argentina'
```