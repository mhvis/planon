# API documentation

(I recommend Insomnia for testing REST calls :) )

### Authentication
Requests are validated using a `X-API-Key` header with a valid key.

### HTTP response status codes
* `2XX`: when the request succeeded.
* `400 Bad request`: when a request parameter is incorrect (e.g. incorrect timestamp format).
* `401 Unauthorized`: when the API key is incorrect.
* `418 I'm a teapot`: when Planon returned an error. The body includes the error message.
* `500 Internal server error`: when something else went wrong. May or may not include a helpful message.

## /reservations

### GET

#### Request

* room_id
* start: RFC3339 timestamp
* end: RFC3339 timestamp

Example: `/reservations?room_id=3100%0a-2%0a-2.380&start=2017-11-01T00:00:00%2b01:00&end=2017-11-30T00:00:00%2b01:00`

#### Response example

`HTTP 200 OK`

```json
[
    {
        "id": "2027944.00",
        "room": {
            "id": "3100\n-2\n-2.386",
            "status": "FREE",
            "reserved": false,
            "people": 0
        },
        "person": {
            "id": "0000207847",
            "last_name": "Rodriguez Gómez",
            "first_name": "Luis",
            "name_prefix": "",
            "email": "l.rodriguez.gomez@student.tue.nl",
            "avail": "UNDEFINED",
            "department": {
                "id": "36",
                "name": "Faculteit Electrical Engineering"
            }
        },
        "mine": false,
        "start": "2017-11-05T16:45:00+01:00",
        "end": "2017-11-05T19:00:00+01:00",
        "type": "OtherReservation",
        "name": "Room reservation"
    },
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
            "last_name": "Kock",
            "first_name": "Bor",
            "name_prefix": "de",
            "email": "b.b.d.kock@student.tue.nl",
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

### POST

#### Request

* room_id
* start: RFC3339
* end: RFC3339
* name (not actually used at the moment)

#### Response


HTTP 201 Created




### DELETE

#### Request

* reservation_id

#### Response

`HTTP 204 No Content` when successfully deleted