# go-http-server

## Commands

### Run application

- `Make run`

### Run tests

- `Make run_tests`

## Example

### user renaming

**Request**

```
curl --location --request POST 'localhost:8090/user' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": 123, "name": "Gabriel"
}'
```

**Response**

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 27 Mar 2022 20:16:22 GMT
Content-Length: 35

{"id":123,"name":"Gabriel_renamed"}%
```