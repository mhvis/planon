# API documentation

All requests need to include an `X-API-Key` header.

All requests and responses have content type `application/json`.

## POST /bookings

### Request parameters

* date: string `yyyy-mm-dd` (required)
* query: object
  * object: string
  * floor: string
  * room: string

### Response parameters

* error: string
* bookings: array of objects
  * room: string
  * bookings: array of objects
    * start: timestamp
    * end: timestamp
    * description: string

## POST /book

### Request parameters

* query: object (required)
  * object: string
  * floor: string
  * room: string
* start: timestamp (required)
* end: timestamp (required)
* description: string

### Response parameters

* error: string

