# Customer Service API

REST API's for Customer Service CRUD & Auth

## API Documentation

### GET /api/v1/customers

#### Response

```json
{
  "status_code": 200,
  "data": [
    {
      "id": 1,
      "first_name": "Jeremiah",
      "last_name": "Ferdinand",
      "email": "jeremylombogia7@gmail.com",
      "password": "$2a$10$Za.lUsMCkJ8SbTi43x03GuyGcE/9Bdjm8Mx9ZWBzXV4lo6TpVXkNa",
      "age": 29
    }
  ]
}
```

### GET /api/v1/customers/{id}

#### Response

```json
{
  "status_code": 200,
  "data": {
    "id": 1,
    "first_name": "Jeremiah",
    "last_name": "Ferdinand",
    "email": "",
    "password": "$2a$10$Il9Uynb0Y/knXrnJhg6EO.h5dnS3XdNBwgK31l4J8QJ/GVFsukWUK",
    "age": 29
  }
}
```

### POST /api/v1/customers

#### Payload

```json
{
  "first_name": "Jeremiah",
  "last_name": "Ferdinand",
  "email": "jeremylombogia7@gmail.com",
  "password": "123",
  "age": 29
}
```

#### Response

```json
{
  "status_code": 201,
  "message": "Data created",
  "data": {
    "id": 1,
    "first_name": "Jeremiah",
    "last_name": "Ferdinand",
    "email": "jeremylombogia7@gmail.com",
    "password": "$2a$10$fSSeuSnTrt6Oo1m9VVS4Y.5cWSSUtUmtBiu7uZJvg8Z8Boj.ysk1i",
    "age": 29
  }
}
```

### PATCH /api/v1/customers/{id}

#### Payload

```json
{
  "first_name": "Jeremiah",
  "last_name": "Ferdinand",
  "email": "jeremylombogia7@gmail.com",
  "password": "123",
  "age": 29
}
```

#### Response

```json
{
  "status_code": 201,
  "message": "Data updated"
}
```

### DELETE /api/v1/customers/{id}

#### Response

```json
{
  "status_code": 200,
  "message": "Data deleted"
}
```