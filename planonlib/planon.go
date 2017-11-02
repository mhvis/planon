// Package planon provides high-level and lower-level functions for Planon BookMySpace.
package planonlib

import (
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"time"
	"log"
)

type Planon struct {
	client        *http.Client
	id            int
	twoWayAuthUrl string
	jUsername     string
	jPassword     string
}

func NewPlanonSession(twoWayAuthUrl, jUsername, jPassword string) *Planon {
	// Create cookie jar
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}
	// Create HTTP client
	client := &http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}
	return &Planon{
		client,
		0,
		twoWayAuthUrl,
		jUsername,
		jPassword,
	}
}

type VersionDetails struct {
	ServerVersion string `json:"serverversion"`
	Version       string `json:"version"`
}

func (p *Planon) ReserveRoom() error {
	return nil
}

func (p *Planon) VersionDetails() (VersionDetails, error) {
	var versionDetails VersionDetails
	err := p.Call("/VersionCheck", "getVersionDetails", nil, &versionDetails)
	return versionDetails, err
}

/*
REQUEST SAMPLES

getReservation
{
  "ldpv": {
    "desc": "-2.380",
    "extradata":{},
    "icon": {
      "type": null,
      "upload": "NO",
      "name": "BOType_Space",
      "id": "BOType_Space"
    },
    "image": null,
    "subtitle": "",
    "title": "-2.380"
  },
  "dpv": {
    "desc": "-2.380",
    "extradata": {},
    "icon": {
      "type":null,
      "upload\":\"NO\",
      \"name\":\"BOType_Space\",
      \"id\":\"BOType_Space\"
    },
    \"image\":null,
    \"subtitle\":\"\",
    \"title\":\"-2.380\"
  },
  "fac":[],
  "isRes":true,
  "oms\":0,
  "people\":1,
  "res\":\"3100.-2.380\",
  "status\":\"FREE\",
  "name\":null,
  "id":"3100\\n-2\\n-2.380"
}
*/

type reservationParams struct {
	ORoom oRoom  `json:"oRoom"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type oRoom struct {
	Id     string `json:"id"`
	People int    `json:"people"`
}

/*
{
	"list": [
		{
			"cEnd":false,
			\"fEnd\":false,
			\"fExt\":false,
			\"canShow\":false,
			\"myRes\":false,
			\"maxExtTime\":null,
			\"beg\":\"2017-11-02T20:33:00+0100\",
			\"end\":\"2017-11-02T22:45:00+0100\",
			\"wsCode\":null,
			\"pers\":null,
			\"resType\":\"ReserveNowFreeTime\",
			\"showStat\":null,
			\"locCode\":null,
			\"canExt\":false,
			\"ldpv\":{
				\"icon\":null,
				\"image\":null,
				\"extradata\":null,
				\"desc\":\"\",
				\"subtitle\":\"\",
				\"title\":\"Meet Now!\"
			},
			\"dpv\":{
				\"icon\":null,
				\"image\":null,
				\"extradata\":null,
				\"desc\":\"\",
				\"subtitle\":\"\",
				\"title\":\"Meet Now!\"
			},
			\"name\":null,
			\"id\":null
		}
	],
	\"searchQuery\":null
}
 */



func (p *Planon) GetReservation(id string, people int, start, end time.Time) error {
	params := make(map[string]interface{})
	params["oRoom"] = oRoom{id, people}
	params["start"] = start.Format(time.RFC3339)
	params["end"] = end.Format(time.RFC3339)

	result := make(map[string]interface{})
	err := p.Call("/RoomInformation", "getReservation", params, result)
	log.Println(result)
	return err
}

func (p *Planon) GetRoomWithCode(searchcode string) error {
	params := map[string]interface{}{"searchcode": searchcode}
	err := p.Call("/RoomInformation", "getRoomWithCode", params, nil)
	return err
	/*
			Response:

			{
			  \"isRes\":true,
			  \"oms\":0,
			  \"people\":1,
			  \"fac\":[],
			  \"status\":\"FREE\",
			  \"res\":\"3100.-2.380\",
			  \"ldpv\":{
			    \"icon\":{
			      \"upload\":\"NO\",
			      \"type\":null,
			      \"name\":\"BOType_Space\",
			      \"id\":\"BOType_Space\"
			    },
			    \"image\":null,
			    \"extradata\":{},
			    \"desc\":\"-2.380\",
			    \"subtitle\":\"\",
			    \"title\":\"-2.380\"
			  },
			  \"dpv\":{
			    \"icon\":{
			      \"upload\":\"NO\",
		          \"type\":null,
		          \"name\":\"BOType_Space\",
		          \"id\":\"BOType_Space\"
		        },
		        \"image\":null,
		        \"extradata\":{},
		        \"desc\":\"-2.380\",
		        \"subtitle\":\"\",
		        \"title\":\"-2.380\"
		      },
		      \"name\":null,
		      \"id\":\"3100\\n-2\\n-2.380\"
		    }
	*/
}
