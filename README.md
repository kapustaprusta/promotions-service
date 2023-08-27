### Application exposes an REST API for uploading and retrieving promotions

#### Upload Promotions

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

#### Get promotion

Endpoint: `GET /promotions/{record_id}`

Request
```sh
curl http://localhost:10100/promotions/1
```

#### Response
```sh
200 - OK
{
  "ID": "d018ef0b-dbd9-48f1-ac1a-eb4d90e57118",
  "Price": 60.683466,
  "ExpirationDate": "2018-08-04T05:32:31+02:00"
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