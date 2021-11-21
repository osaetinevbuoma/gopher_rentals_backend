# GOPHER RENTALS

Gopher Rentals is a small project to get familiar with the Go programming languages using the
[Gin Gonic](https://gin-gonic.com/) web framework. This repository consists of a series of REST API
endpoints for managing the services of fictitious Gopher Rentals company.

## Setup

This implementation of the Gopher Rentals project does not leverage any ORM for database
interactions. It uses age-old MySQL queries stored in repositories that are used by service
functions.

In terminal, enter the following command:

```
$ mysql -u <MYSQL_USER> -p
```

Enter your MySQL user's password. While in the MySQL prompt, enter the following commands:

```
$ mysql > source /path/to/sql/file/in/db/folder/database.sql
```

This creates the database (`gopher_rentals`) and the corresponding tables.

## REST API Endpoints

### Admin User Management

#### Register Admin User

```
POST: /api/register

Payload:
{
    "email": "john.doe@email.com",
    "password": "123",
    "confirm_password": "123"
}

Response:
{
    "id": "9941a176-ee92-4c67-b7d0-b8f9c4149a24",
    "email": "john.doe@email.com",
    "password": "$2a$16$CC3a1CYWwFWwn2uD8s6/1OG.9Yk2cHAq3L.DA4DuvbCj4D7Scvb4a"
}
```

#### Login

```
POST: /api/login

Payload:
{
    "email": "john.doe@email.com",
    "password": "123"
}

Response:
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQ2NzFkYjQ1LTEzOWMtNGQzMS05ZjcyLTQ0ZjQwZjYwNDkwOCIsImVtYWlsIjoiam9obi5kb2VAZW1haWwuY29tIiwiaXNzdWVkX2F0IjoiMjAyMS0xMS0yMFQxODoyNjozMS4wMjA1MzY2MjcrMDE6MDAiLCJleHBpcmVkX2F0IjoiMjAyMS0xMS0yMFQxODoyNjozMS4wMjA2MjMxNTkrMDE6MDAifQ.Eqw7rH8LFMKslr6TK_5geJJPv3J0caUdKvPk9VJHyqw"
```

### Customer Management

#### List Customers

```
GET: /api/customers

Response:
[
  {
    "id": "1e49d592-1a0a-4473-931c-7f828c250bed",
    "first_name": "John",
    "last_name": "Doe",
    "nationality": "Nigerian",
    "identification_number": "ABC123",
    "identification_type": "International Passport"
  }
]
```

#### Get Customer

```
GET: /api/customers/:id

Response:
{
  "id": "95a3280b-c93c-42fe-bdfe-b6107e4d83d2",
  "first_name": "John",
  "last_name": "Doe",
  "nationality": "Nigerian",
  "identification_number": "ABC123",
  "identification_type": "International Passport"
}
```

#### POST Customer

```
POST: /api/customers/create

Payload:
{
    "first_name": "John",
    "last_name": "Doe",
    "nationality": "Nigerian",
    "identification_number": "ABC123",
    "identification_type": "International Passport"
}

Response:
{
    "id": "1e49d592-1a0a-4473-931c-7f828c250bed",
    "first_name": "John",
    "last_name": "Doe",
    "nationality": "Nigerian",
    "identification_number": "ABC123",
    "identification_type": "International Passport"
}
```

#### Edit Customer

```
PUT: /api/customers/edit

Payload:
{
    "id": "e42e1776-cd31-4fae-adc4-293d74956e70",
    "first_name": "Johnson",
    "last_name": "Doe",
    "nationality": "English",
    "identification_number": "ABC123",
    "identification_type": "International Passport"
}

Response:
{
    "id": "e42e1776-cd31-4fae-adc4-293d74956e70",
    "first_name": "Johnson",
    "last_name": "Doe",
    "nationality": "English",
    "identification_number": "ABC123",
    "identification_type": "International Passport"
}
```

#### Delete Customer

```
DELETE: /api/customers/delete/:id
```

### Car Management

#### List Cars

```
GET: /api/cars

Response:
[
  {
    "id": "32093eb3-0dd2-4e8d-be42-2d3d5d59447d",
    "model": "Honda",
    "year": 2016,
    "license_plate": "ABC123ERT",
    "current_km": 560.9,
    "max_kg": 34.56,
    "fuel_type": "Petrol",
    "hire_price": 908.67,
    "hire_availability": false,
    "locations": [
      {
        "id": "3ea857a4-ab32-4fc8-a979-c7feb7ada0b3",
        "car": {
          "id": "00000000-0000-0000-0000-000000000000",
          "model": "",
          "year": 0,
          "license_plate": "",
          "current_km": 0,
          "max_kg": 0,
          "fuel_type": "",
          "hire_price": 0,
          "hire_availability": false,
          "locations": null
        },
        "latitude": 78.93,
        "longitude": 289.214,
        "current_location_datetime": "2021-11-26T09:40:01Z"
      }
    ]
  },
  {
    "id": "3eb13d0b-549d-4a12-bd87-8074504616e4",
    "model": "Toyota",
    "year": 2011,
    "license_plate": "ABC123ER",
    "current_km": 560.9,
    "max_kg": 34.56,
    "fuel_type": "Petrol",
    "hire_price": 908.67,
    "hire_availability": false,
    "locations": null
  }
]
```

#### Get Car

```
GET: /api/cars/:id

Response:
{
  "id": "09bda36b-f1dd-4290-883e-6b4a49589139",
  "model": "Mazda",
  "year": 2011,
  "license_plate": "ABC123ER",
  "current_km": 560.9,
  "max_kg": 34.56,
  "fuel_type": "Petrol",
  "hire_price": 908.67,
  "hire_availability": false,
  "locations": null
}
```

#### Create Car

```
POST: /api/cars/create

Payload:
{
	"model": "Mazda",
	"year": 2011,
	"license_plate": "ABC123ER",
	"current_km": 560.90,
	"max_kg": 34.56,
	"fuel_type": "Petrol",
	"hire_price": 908.67
}

Response: 
{
  "id": "32093eb3-0dd2-4e8d-be42-2d3d5d59447d",
  "model": "Mazda",
  "year": 2011,
  "license_plate": "ABC123ER",
  "current_km": 560.9,
  "max_kg": 34.56,
  "fuel_type": "Petrol",
  "hire_price": 908.67,
  "hire_availability": false,
  "locations": null
}
```

#### Edit Car

```
PUT: /api/cars/edit

Payload:
{
	"id": "32093eb3-0dd2-4e8d-be42-2d3d5d59447d",
	"model": "Honda",
	"year": 2016,
	"license_plate": "ABC123ERT",
	"current_km": 560.90,
	"max_kg": 34.56,
	"fuel_type": "Petrol",
	"hire_price": 908.67
}

Response:
{
  "id": "32093eb3-0dd2-4e8d-be42-2d3d5d59447d",
  "model": "Honda",
  "year": 2016,
  "license_plate": "ABC123ERT",
  "current_km": 560.9,
  "max_kg": 34.56,
  "fuel_type": "Petrol",
  "hire_price": 908.67,
  "hire_availability": false,
  "locations": null
}
```

#### Delete Car
```
DELETE: /api/cars/delete/:id
```

### Car Location Management

#### List Car Locations

```
GET: /api/cars/:carId/locations

Response:
[
  {
    "id": "2ab3f402-a8e2-485b-80bf-633dd87f0e0c",
    "car": {
      "id": "00000000-0000-0000-0000-000000000000",
      "model": "",
      "year": 0,
      "license_plate": "",
      "current_km": 0,
      "max_kg": 0,
      "fuel_type": "",
      "hire_price": 0,
      "hire_availability": false,
      "locations": null
    },
    "latitude": 78.93,
    "longitude": 289.214,
    "current_location_datetime": "2021-11-26T09:40:01Z"
  }
]
```

#### List `X` number of recent locations of car

```
GET: /api/cars/:carId/locations/:recent

Response:
[
  {
    "id": "2ab3f402-a8e2-485b-80bf-633dd87f0e0c",
    "car": {
      "id": "00000000-0000-0000-0000-000000000000",
      "model": "",
      "year": 0,
      "license_plate": "",
      "current_km": 0,
      "max_kg": 0,
      "fuel_type": "",
      "hire_price": 0,
      "hire_availability": false,
      "locations": null
    },
    "latitude": 78.93,
    "longitude": 289.214,
    "current_location_datetime": "2021-11-26T09:40:01Z"
  }
]
```

#### Create Car Location

```
POST: /api/cars/:carId/locations/create

Payload:
{
	"latitude": 78.930,
	"longitude": 289.214,
	"current_location_datetime": "2021-11-21 09:40:01"
}

Response:
{
  "id": "2ab3f402-a8e2-485b-80bf-633dd87f0e0c",
  "car": {
    "id": "e63d7240-c464-4014-8113-f6389f841e7a",
    "model": "Honda",
    "year": 2016,
    "license_plate": "ABC123ERT",
    "current_km": 560.9,
    "max_kg": 34.56,
    "fuel_type": "Petrol",
    "hire_price": 908.67,
    "hire_availability": false,
    "locations": null
  },
  "latitude": 78.93,
  "longitude": 289.214,
  "current_location_datetime": "2021-11-21T09:40:01Z"
}
```

#### Edit Car Location

```
PUT: /api/cars/:carId/locations/edit

Payload:
{
	"id": "2ab3f402-a8e2-485b-80bf-633dd87f0e0c",
	"latitude": 78.930,
	"longitude": 289.214,
	"current_location_datetime": "2021-11-26 09:40:01"
}

Response:
{
  "id": "2ab3f402-a8e2-485b-80bf-633dd87f0e0c",
  "car": {
    "id": "e63d7240-c464-4014-8113-f6389f841e7a",
    "model": "Honda",
    "year": 2016,
    "license_plate": "ABC123ERT",
    "current_km": 560.9,
    "max_kg": 34.56,
    "fuel_type": "Petrol",
    "hire_price": 908.67,
    "hire_availability": false,
    "locations": null
  },
  "latitude": 78.93,
  "longitude": 289.214,
  "current_location_datetime": "2021-11-26T09:40:01Z"
}
```

#### Delete Car Location

```
DELETE: /api/cars/:carId/locations/delete/:id
```