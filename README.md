# promotions-service
Application exposes REST API for uploading and retrieving promotions

## Download
To download the application repository you should launch the next command
```
git clone git@github.com:kapustaprusta/promotions-service.git promotions-service
```

## Build
To build the application you should launch the next commands
```
cd promotions-service
make build
```

## Test
To run of all tests you should launch the next command
```
make test
```

## Lint
To run linter you should launch the next command
```
make lint
````

## Compose
To build and run docker image you should launch the next command
```
make compose
````

## How it works

### Upload Promotions

Endpoint: `POST /promotions`

Request
```sh
curl -X POST -H "Content-Type: application/csv" \
        --data-binary @promotions.csv \
        http://localhost:10100/promotions
```

Responses
```sh
200 - OK
{
  "total_promotions":3
}

400 - BAD REQUEST
{
  "slug": "invalid-input"
}

500 - INTERNAL SERVER ERROR
{
  "slug": "some error message"
}
```

### Get promotion

Endpoint: `GET /promotions/{record_id}`

Request
```sh
curl http://localhost:10100/promotions/1
```

Response
```sh
200 - OK
{
  "id": "d018ef0b-dbd9-48f1-ac1a-eb4d90e57118",
  "price": 60.683466,
  "expiration_date": "2018-08-04 05:32:31 +0200 CEST"
}

400 - BAD REQUEST
{
  "slug": "cannot-parse-record-id"
}

404 - NOT FOUND
{
  "slug": "promotion-not-found"
}

500 - INTERNAL SERVER ERROR
{
  "slug": "some error message"
}
```