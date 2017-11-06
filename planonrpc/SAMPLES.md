# /RoomInformation getReservation

## Request

```json
{
   "start":"RFC3339",
   "end":"RFC3339",
   "oRoom":{
      "ldpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "dpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "fac":[

      ],
      "isRes":true,
      "oms":0,
      "people":1,
      "res":"3100.-2.380",
      "status":"FREE",
      "name":null,
      "id":"3100\n-2\n-2.380"
   }
}
```

## Response

```json
{
   "list":[
      {
         "cEnd":false,
         "fEnd":false,
         "fExt":false,
         "canShow":false,
         "myRes":false,
         "maxExtTime":null,
         "beg":"2017-11-01T15:07:09+0100",
         "end":"2017-11-01T17:07:09+0100",
         "wsCode":null,
         "pers":{
            "mgr":null,
            "dep":{
               "name":"Faculteit Wiskunde \u0026 Informatica",
               "id":"32"
            },
            "tel":{

            },
            "avail":"UNDEFINED",
            "ema":"g.w.v.d.heijden@student.tue.nl",
            "fnm":"Wessel",
            "func":null,
            "prefix":"van der",
            "registrationID":null,
            "room":null,
            "ldpv":{
               "icon":{
                  "upload":"NO",
                  "type":null,
                  "name":"BOType_Reservation",
                  "id":"BOType_Reservation"
               },
               "image":{
                  "upload":"NO",
                  "type":"Person",
                  "name":null,
                  "id":"0000110179"
               },
               "extradata":{
                  "address":"Groene Loper 5, Eindhoven",
                  "workmail":"g.w.v.d.heijden@student.tue.nl"
               },
               "desc":"Heijden, G.W. van der Wessel",
               "subtitle":"",
               "title":"Heijden, G.W. van der Wessel"
            },
            "dpv":{
               "icon":{
                  "upload":"NO",
                  "type":null,
                  "name":"BOType_Reservation",
                  "id":"BOType_Reservation"
               },
               "image":{
                  "upload":"NO",
                  "type":"Person",
                  "name":null,
                  "id":"0000110179"
               },
               "extradata":{
                  "address":"Groene Loper 5, Eindhoven",
                  "workmail":"g.w.v.d.heijden@student.tue.nl"
               },
               "desc":"Heijden, G.W. van der Wessel",
               "subtitle":"",
               "title":"Heijden, G.W. van der Wessel"
            },
            "name":"Heijden",
            "id":"0000110179"
         },
         "resType":"OtherReservation",
         "showStat":"UNKNOWN",
         "locCode":{
            "isRes":false,
            "oms":0,
            "people":0,
            "fac":null,
            "status":"FREE",
            "res":null,
            "ldpv":{
               "icon":null,
               "image":null,
               "extradata":null,
               "desc":null,
               "subtitle":null,
               "title":null
            },
            "dpv":{
               "icon":null,
               "image":null,
               "extradata":null,
               "desc":null,
               "subtitle":null,
               "title":null
            },
            "name":null,
            "id":"3100\n-2\n-2.386"
         },
         "canExt":false,
         "ldpv":{
            "icon":{
               "upload":"NO",
               "type":null,
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "image":{
               "upload":"NO",
               "type":null,
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "extradata":{

            },
            "desc":"Room reservation",
            "subtitle":"",
            "title":"Room reservation"
         },
         "dpv":{
            "icon":{
               "upload":"NO",
               "type":null,
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "image":{
               "upload":"NO",
               "type":null,
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "extradata":{

            },
            "desc":"Room reservation",
            "subtitle":"",
            "title":"Room reservation"
         },
         "name":"Room reservation",
         "id":"2004442.00"
      },
      {
         "cEnd":false,
         "fEnd":false,
         "fExt":false,
         "canShow":false,
         "myRes":false,
         "maxExtTime":null,
         "beg":"2017-11-03T00:52:00+0100",
         "end":"2017-11-03T03:00:00+0100",
         "wsCode":null,
         "pers":null,
         "resType":"ReserveNowFreeTime",
         "showStat":null,
         "locCode":null,
         "canExt":false,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":"",
            "subtitle":"",
            "title":"Meet Now!"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":"",
            "subtitle":"",
            "title":"Meet Now!"
         },
         "name":null,
         "id":null
      }
   ],
   "searchQuery":null
}
```

# /MobileSettings getSettings

## Request

`{}`

## Response

```json
{
   "autoSearchCount":3,
   "buttonDetails":[
      {
         "ipu":false,
         "key":"MY_MEETING",
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":"",
         "id":""
      },
      {
         "ipu":false,
         "key":"MY_CAD_MEETING",
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":"",
         "id":""
      },
      {
         "ipu":false,
         "key":"MY_PEOPLE_SEARCH",
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":"",
         "id":""
      },
      {
         "ipu":false,
         "key":"MY_CAD_FLEXSPACE",
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":"",
         "id":""
      },
      {
         "ipu":false,
         "key":"MY_WEBPAGE",
         "ldpv":{
            "icon":{
               "upload":"HONORMAIN",
               "type":"Button",
               "name":"3",
               "id":"3"
            },
            "image":{
               "upload":"HONORMAIN",
               "type":"Button",
               "name":"3",
               "id":"3"
            },
            "extradata":{
               "color":"33cc00",
               "url":"https://agnes2.campus.tue.nl:18443/case/tue/MOB_002"
            },
            "desc":"",
            "subtitle":"Reserveren",
            "title":"Reserveren"
         },
         "dpv":{
            "icon":{
               "upload":"HONORMAIN",
               "type":"Button",
               "name":"3",
               "id":"3"
            },
            "image":{
               "upload":"HONORMAIN",
               "type":"Button",
               "name":"3",
               "id":"3"
            },
            "extradata":{
               "color":"33cc00",
               "url":"https://agnes2.campus.tue.nl:18443/case/tue/MOB_002"
            },
            "desc":"",
            "subtitle":"Reserveren",
            "title":"Reserveren"
         },
         "name":"MY_WEBPAGE",
         "id":"3"
      }
   ],
   "listOfDisabledUsecases":[
      "REPORT_NOW",
      "MY_ACTIONS",
      "MY_ASSETS",
      "FORCE_RELOGIN"
   ]
}
```

# /UserInformation getMe

## Request
`{}`

## Response

```json
{
   "mgr":null,
   "dep":{
      "name":"Faculteit Wiskunde \u0026 Informatica",
      "id":"32"
   },
   "tel":{

   },
   "avail":"PRIVATE",
   "ema":"m.h.visscher@student.tue.nl",
   "fnm":"Maarten",
   "func":null,
   "prefix":null,
   "registrationID":null,
   "room":null,
   "ldpv":{
      "icon":{
         "upload":"NO",
         "type":null,
         "name":"BOType_UDVisitor1",
         "id":"BOType_UDVisitor1"
      },
      "image":{
         "upload":"NO",
         "type":"Person",
         "name":null,
         "id":"0000109425"
      },
      "extradata":{
         "address":"Groene Loper 5, Eindhoven",
         "workmail":"m.h.visscher@student.tue.nl"
      },
      "desc":"Visscher, M.H. Maarten",
      "subtitle":"",
      "title":"Visscher, M.H. Maarten"
   },
   "dpv":{
      "icon":{
         "upload":"NO",
         "type":null,
         "name":"BOType_UDVisitor1",
         "id":"BOType_UDVisitor1"
      },
      "image":{
         "upload":"NO",
         "type":"Person",
         "name":null,
         "id":"0000109425"
      },
      "extradata":{
         "address":"Groene Loper 5, Eindhoven",
         "workmail":"m.h.visscher@student.tue.nl"
      },
      "desc":"Visscher, M.H. Maarten",
      "subtitle":"",
      "title":"Visscher, M.H. Maarten"
   },
   "name":"Visscher",
   "id":"0000109425"
}
```

# /FloorInformation getFreeRooms

## Request

```json
{
   "theProperty":{
      "loc":{
         "adr":"De Lampendriessen 31",
         "city":"Eindhoven",
         "nat":null,
         "pc":"5612AH"
      },
      "bdType":null,
      "name":"Luna",
      "id":"3100"
   }
}
```

## Response

```json
{
   "list":[
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"-2.380"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n-2\n-2.380"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"-2.382"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n-2\n-2.382"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"-2.384"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n-2\n-2.384"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"-2.386"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n-2\n-2.386"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"1.042"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n01\n1.042"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"1.050"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n01\n1.050"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"1.056"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n01\n1.056"
      },
      {
         "isRes":true,
         "oms":0,
         "people":0,
         "fac":null,
         "status":"FREE",
         "res":null,
         "ldpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":"1.240"
         },
         "dpv":{
            "icon":null,
            "image":null,
            "extradata":null,
            "desc":null,
            "subtitle":null,
            "title":null
         },
         "name":null,
         "id":"3100\n01\n1.240"
      }
   ],
   "searchQuery":null
}
```

# /FloorInformation getPropertyList

## Request
`{}`

## Response
```json
{
   "list":[
      {
         "loc":{
            "adr":"De Lismortel",
            "city":"Eindhoven",
            "nat":null,
            "pc":null
         },
         "bdType":null,
         "name":"Paviljoen Vleugel C",
         "id":"0163"
      },
      {
         "loc":{
            "adr":"De Zaale (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AJ"
         },
         "bdType":null,
         "name":"Traverse",
         "id":"1200"
      },
      {
         "loc":{
            "adr":"Den Dolech (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AZ"
         },
         "bdType":null,
         "name":"Auditorium",
         "id":"1300"
      },
      {
         "loc":{
            "adr":"Het Eeuwsel (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AS"
         },
         "bdType":null,
         "name":"IPO-gebouw",
         "id":"1400"
      },
      {
         "loc":{
            "adr":"Horsten 1",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AX"
         },
         "bdType":null,
         "name":"Multimediapaviljoen",
         "id":"2300"
      },
      {
         "loc":{
            "adr":"De Lampendriessen 31",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612AH"
         },
         "bdType":null,
         "name":"Luna",
         "id":"3100"
      },
      {
         "loc":{
            "adr":"Groene Loper 5",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612AE"
         },
         "bdType":null,
         "name":"MetaForum",
         "id":"4400"
      },
      {
         "loc":{
            "adr":"De Wielen (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AV"
         },
         "bdType":null,
         "name":"Vertigo",
         "id":"5100"
      },
      {
         "loc":{
            "adr":"De Wielen 6",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AV"
         },
         "bdType":null,
         "name":"Zwarte Doos",
         "id":"5500"
      },
      {
         "loc":{
            "adr":"Het Kranenveld",
            "city":"Eindhoven",
            "nat":null,
            "pc":null
         },
         "bdType":null,
         "name":"Helix-west",
         "id":"5801"
      },
      {
         "loc":{
            "adr":"Het Kranenveld",
            "city":"Eindhoven",
            "nat":null,
            "pc":null
         },
         "bdType":null,
         "name":"Helix-oost",
         "id":"5802"
      },
      {
         "loc":{
            "adr":null,
            "city":"Eindhoven",
            "nat":"Nederland",
            "pc":null
         },
         "bdType":null,
         "name":"Flux",
         "id":"7000"
      },
      {
         "loc":{
            "adr":"De Rondom (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AP"
         },
         "bdType":null,
         "name":"Cascade",
         "id":"7400"
      },
      {
         "loc":{
            "adr":"De Wielen (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AV"
         },
         "bdType":null,
         "name":"Gemini zuid",
         "id":"8100"
      },
      {
         "loc":{
            "adr":"De Wielen (Ongenummerd)",
            "city":"Eindhoven",
            "nat":null,
            "pc":"5612 AV"
         },
         "bdType":null,
         "name":"Gemini noord",
         "id":"8200"
      },
      {
         "loc":{
            "adr":"De Zaale (Ongenummerd)",
            "city":"Eindhoven",
            "nat":"Netherlands",
            "pc":"5612 AJ"
         },
         "bdType":null,
         "name":"Laplacegebouw",
         "id":"8300"
      }
   ],
   "searchQuery":null
}
```
# /FloorInformation getFloorsOfProperty

## Request
```json
{
   "theProperty":{
      "loc":{
         "adr":"De Lampendriessen 31",
         "city":"Eindhoven",
         "nat":null,
         "pc":"5612AH"
      },
      "bdType":null,
      "name":"Luna",
      "id":"3100"
   }
}
```

## Response

```json
{
   "findex":-1,
   "list":[
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - kelder",
         "id":"10000444"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - kelder",
         "id":"10000445"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - begane grond",
         "id":"10002297"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - eerste etage",
         "id":"10000446"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - tweede etage",
         "id":"10000447"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - derde etage",
         "id":"10000448"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - vierde etage",
         "id":"10000449"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - vijfde etage",
         "id":"10000450"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - zesde etage",
         "id":"10000451"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - zevende etage",
         "id":"10000452"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - achtste etage",
         "id":"10000453"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - negende etage",
         "id":"10000454"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - tiende etage",
         "id":"10000455"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - elfde etage",
         "id":"10000456"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - twaalfde etage",
         "id":"10002314"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - dertiende etage",
         "id":"10002315"
      },
      {
         "prop":{
            "loc":{
               "adr":"De Lampendriessen 31",
               "city":"Eindhoven",
               "nat":null,
               "pc":"5612AH"
            },
            "bdType":null,
            "name":"Luna",
            "id":"3100"
         },
         "name":"Luna - veertiende etage",
         "id":"10002316"
      }
   ],
   "searchQuery":null
}
```

# /RoomInformation reserveRoom

## Request
```json
{
   "end":"2017-10-28T23:30:00+0200",
   "oRoom":{
      "ldpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "dpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "fac":[

      ],
      "isRes":true,
      "oms":0,
      "people":1,
      "res":"3100.-2.380",
      "status":"FREE",
      "name":null,
      "id":"3100\n-2\n-2.380"
   },
   "start":"2017-10-28T21:30:00+0200"
}
```

## Response

```json
"true"
```

# /RoomInformation extendReservation

## Request

```json
{
   "end":"2017-10-28T23:45:00+0200",
   "oRoom":{
      "ldpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "dpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "fac":[

      ],
      "isRes":true,
      "oms":0,
      "people":1,
      "res":"3100.-2.380",
      "status":"FREE",
      "name":null,
      "id":"3100\n-2\n-2.380"
   },
   "reservation":{
      "cEnd":false,
      "fEnd":true,
      "fExt":true,
      "canShow":false,
      "showStat":"SHOWN_UP",
      "ldpv":{
         "desc":"Room reservation",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "image":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "subtitle":"",
         "title":"Room reservation"
      },
      "dpv":{
         "desc":"Room reservation",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "image":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "subtitle":"",
         "title":"Room reservation"
      },
      "maxExtTime":null,
      "locCode":{
         "ldpv":{
            "desc":null,
            "extradata":null,
            "icon":null,
            "image":null,
            "subtitle":null,
            "title":null
         },
         "dpv":{
            "desc":null,
            "extradata":null,
            "icon":null,
            "image":null,
            "subtitle":null,
            "title":null
         },
         "fac":null,
         "isRes":false,
         "oms":0,
         "people":0,
         "res":null,
         "status":"FREE",
         "name":null,
         "id":"3100\n-2\n-2.380"
      },
      "myRes":true,
      "pers":{
         "avail":"PRIVATE",
         "dep":{
            "name":"Faculteit Wiskunde \u0026 Informatica",
            "id":"32"
         },
         "ldpv":{
            "desc":"Visscher, M.H. Maarten",
            "extradata":{
               "address":"Groene Loper 5, Eindhoven",
               "workmail":"m.h.visscher@student.tue.nl"
            },
            "icon":{
               "type":null,
               "upload":"NO",
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "image":{
               "type":"Person",
               "upload":"NO",
               "name":null,
               "id":"0000109425"
            },
            "subtitle":"",
            "title":"Visscher, M.H. Maarten"
         },
         "dpv":{
            "desc":"Visscher, M.H. Maarten",
            "extradata":{
               "address":"Groene Loper 5, Eindhoven",
               "workmail":"m.h.visscher@student.tue.nl"
            },
            "icon":{
               "type":null,
               "upload":"NO",
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "image":{
               "type":"Person",
               "upload":"NO",
               "name":null,
               "id":"0000109425"
            },
            "subtitle":"",
            "title":"Visscher, M.H. Maarten"
         },
         "ema":"m.h.visscher@student.tue.nl",
         "fnm":"Maarten",
         "func":null,
         "mgr":null,
         "prefix":null,
         "registrationID":null,
         "room":null,
         "tel":{

         },
         "name":"Visscher",
         "id":"0000109425"
      },
      "resType":"MyReservation",
      "beg":"2017-10-28T21:30:00+0200",
      "end":"2017-10-28T23:30:00+0200",
      "canExt":false,
      "wsCode":null,
      "name":"Room reservation",
      "id":"1979544.00"
   }
}
```

## Response
```json
"true"
```

# /RoomInformation endReservation

## Request
```json
{
   "oRoom":{
      "ldpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "dpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "fac":[

      ],
      "isRes":true,
      "oms":0,
      "people":1,
      "res":"3100.-2.380",
      "status":"FREE",
      "name":null,
      "id":"3100\n-2\n-2.380"
   },
   "reservation":{
      "cEnd":false,
      "fEnd":true,
      "fExt":true,
      "canShow":false,
      "showStat":"SHOWN_UP",
      "ldpv":{
         "desc":"Room reservation",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "image":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "subtitle":"",
         "title":"Room reservation"
      },
      "dpv":{
         "desc":"Room reservation",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "image":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Reservation",
            "id":"BOType_Reservation"
         },
         "subtitle":"",
         "title":"Room reservation"
      },
      "maxExtTime":null,
      "locCode":{
         "ldpv":{
            "desc":null,
            "extradata":null,
            "icon":null,
            "image":null,
            "subtitle":null,
            "title":null
         },
         "dpv":{
            "desc":null,
            "extradata":null,
            "icon":null,
            "image":null,
            "subtitle":null,
            "title":null
         },
         "fac":null,
         "isRes":false,
         "oms":0,
         "people":0,
         "res":null,
         "status":"FREE",
         "name":null,
         "id":"3100\n-2\n-2.380"
      },
      "myRes":true,
      "pers":{
         "avail":"PRIVATE",
         "dep":{
            "name":"Faculteit Wiskunde \u0026 Informatica",
            "id":"32"
         },
         "ldpv":{
            "desc":"Visscher, M.H. Maarten",
            "extradata":{
               "address":"Groene Loper 5, Eindhoven",
               "workmail":"m.h.visscher@student.tue.nl"
            },
            "icon":{
               "type":null,
               "upload":"NO",
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "image":{
               "type":"Person",
               "upload":"NO",
               "name":null,
               "id":"0000109425"
            },
            "subtitle":"",
            "title":"Visscher, M.H. Maarten"
         },
         "dpv":{
            "desc":"Visscher, M.H. Maarten",
            "extradata":{
               "address":"Groene Loper 5, Eindhoven",
               "workmail":"m.h.visscher@student.tue.nl"
            },
            "icon":{
               "type":null,
               "upload":"NO",
               "name":"BOType_Reservation",
               "id":"BOType_Reservation"
            },
            "image":{
               "type":"Person",
               "upload":"NO",
               "name":null,
               "id":"0000109425"
            },
            "subtitle":"",
            "title":"Visscher, M.H. Maarten"
         },
         "ema":"m.h.visscher@student.tue.nl",
         "fnm":"Maarten",
         "func":null,
         "mgr":null,
         "prefix":null,
         "registrationID":null,
         "room":null,
         "tel":{

         },
         "name":"Visscher",
         "id":"0000109425"
      },
      "resType":"MyReservation",
      "beg":"2017-10-28T21:30:00+0200",
      "end":"2017-10-28T23:45:00+0200",
      "canExt":false,
      "wsCode":null,
      "name":"Room reservation",
      "id":"1979544.00"
   }
}
```

## Response

```json
null
```

# /RoomInformation getAllowedreservationTimes

## Request
```json
{
   "oRoom":{
      "ldpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "dpv":{
         "desc":"-2.380",
         "extradata":{

         },
         "icon":{
            "type":null,
            "upload":"NO",
            "name":"BOType_Space",
            "id":"BOType_Space"
         },
         "image":null,
         "subtitle":"",
         "title":"-2.380"
      },
      "fac":[

      ],
      "isRes":true,
      "oms":0,
      "people":1,
      "res":"3100.-2.380",
      "status":"FREE",
      "name":null,
      "id":"3100\n-2\n-2.380"
   }
}
```

## Response

```json
{
   "list":[
      {
         "endWorkDay":false,
         "end":"2017-10-28T22:00:00+0200",
         "start":"2017-10-28T21:30:00+0200",
         "length":1800
      }
   ],
   "searchQuery":null
}
```

# /RoomInformation getRoomWithCode
## Request
```json
{
   "searchcode":"3100\n-2\n-2.380"
}
```
## Response
```json
{
   "isRes":true,
   "oms":0,
   "people":1,
   "fac":[

   ],
   "status":"FREE",
   "res":"3100.-2.380",
   "ldpv":{
      "icon":{
         "upload":"NO",
         "type":null,
         "name":"BOType_Space",
         "id":"BOType_Space"
      },
      "image":null,
      "extradata":{

      },
      "desc":"-2.380",
      "subtitle":"",
      "title":"-2.380"
   },
   "dpv":{
      "icon":{
         "upload":"NO",
         "type":null,
         "name":"BOType_Space",
         "id":"BOType_Space"
      },
      "image":null,
      "extradata":{

      },
      "desc":"-2.380",
      "subtitle":"",
      "title":"-2.380"
   },
   "name":null,
   "id":"3100\n-2\n-2.380"
}
```