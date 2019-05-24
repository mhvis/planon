# API documentation

### Requests
For simplicity, all requests must be POST requests with Content-Type `application/x-www-form-urlencoded` for the parameters.

### Authentication
Requests are validated using a `X-API-Key` header with a valid key.

### Responses
* `200 OK`: the action succeeded, `application/json` data is returned.
* `204 No content`: the action succeeded, no data is returned.
* `400 Bad request`: when a request parameter is incorrect (e.g. incorrect timestamp format).
* `401 Unauthorized`: when the API key is incorrect.
* `418 I'm a teapot`: when Planon returned an error. The body includes the error message.
* `500 Internal server error`: when something else went wrong. May or may not include a helpful message.

## /getReservations

* room_id
* start: RFC3339 timestamp
* end: RFC3339 timestamp

Response: `200 OK`.

Response example:
```json
[
    {
        "id": "1997966.00",
        "room": {
            "id": "3100\n-2\n-2.386",
            "status": "FREE",
            "reserved": false,
            "people": 0
        },
        "person": {
            "id": "0000108538",
            "last_name": "Raap",
            "first_name": "Henk",
            "name_prefix": "van",
            "email": "some@email.nl",
            "avail": "PRIVATE",
            "department": {
                "id": "32",
                "name": "Faculteit Wiskunde \u0026 Informatica"
            }
        },
        "mine": false,
        "start": "2017-11-07T18:30:00+01:00",
        "end": "2017-11-07T23:00:00+01:00",
        "type": "OtherReservation",
        "name": "Room reservation"
    }
]
```

## /reserveRoom

* room_id
* start: RFC3339
* end: RFC3339
* name (not actually used at the moment)

Response: `204 No content`.

## /endReservation

* room_id
* start: RFC3339
* end: RFC3339

Response: `204 No content`.

## /extendReservation

* reservation_id
* end: RFC3339

Response: `204 No content`.

## /getPropertyList

Response: `200 OK`.

## /getFloorsOfProperty

* property_id

Response: `200 OK`.

## /getFreeRooms

* property_id

Response: `200 OK`.

## /getMe

Response: `200 OK`.
