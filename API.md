# API documentation

All requests need to include an `X-API-Key` header, otherwise `401 Unauthorized` should be returned.

All responses have content type `application/json`.

SSL everywhere. Also no redirection from http to https to prevent unencrypted requests which are silently redirected ([source](http://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api#ssl)).

## /bookings

### GET

#### Request

* date: `yyyy-mm-dd` (required)
* object
* floor
* room

Example: `/bookings?date=2017-03-03&object=3100`

#### Response

`HTTP 200 OK`

```json
[  
   {  
      "room":"-2.380",
      "bookings":[  
         {  
            "start":12345,
            "end":12345,
            "description":"Hallo"
         }
      ]
   }
]

```

`HTTP 500`

```json
{
 "message": "Error message"
}
```

### POST

#### Request

* 

#### Response


HTTP 201 Created




### DELETE

#### Request

* start: timestamp
* object
* room


`HTTP 204 No Content` when successfully deleted

`HTTP 500` on error


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

