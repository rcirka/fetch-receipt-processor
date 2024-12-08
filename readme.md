## API Endpoints

### Process Receipt
- **POST** `/receipts/process`
- Processes a receipt and returns a unique ID
- Request Body:

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
  ],
  "total": "6.49"
}
```
- Response:

```json
{
  "id": "uuid-string"
}
```

### Get Points
- **GET** `/receipts/{id}/points`
- Returns the points awarded for the receipt
- Response:

```json
{
  "points": 32
}
```

## Points Calculation Rules

1. One point for every alphanumeric character in the retailer name
2. 50 points if the total is a round dollar amount with no cents
3. 25 points if the total is a multiple of 0.25
4. 5 points for every two items on the receipt
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer
6. 6 points if the day in the purchase date is odd
7. 10 points if the time of purchase is between 2:00pm and 4:00pm

## Running the Application

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```
3. Run the server:

```bash
go run main.go
```
The server will start on port 8080.

## Running Tests

Execute the test suite with:

```bash
go test ./...
```

## Data Storage

The application uses in-memory storage (map) to store receipts. Note that data will be lost when the server restarts.