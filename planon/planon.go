// Package planon provides high-level and lower-level functions for Planon BookMySpace.
package planon

import (
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Planon struct {
	client            *http.Client
	id                int
	jsonRpcUrl        string
	jSecurityCheckUrl string
	jUsername         string
	jPassword         string
}

func NewPlanonSession(jsonRpcUrl, jSecurityCheckUrl, jUsername, jPassword string) *Planon {
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
		jsonRpcUrl,
		jSecurityCheckUrl,
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

func (p *Planon) GetReservation(start, end time.Time) error {
	err := p.Call("/RoomInformation", "getReservation", nil, nil)
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
		      \"upload\":\"NO\",\"type\":null,\"name\":\"BOType_Space\",\"id\":\"BOType_Space\"},\"image\":null,\"extradata\":{},\"desc\":\"-2.380\",\"subtitle\":\"\",\"title\":\"-2.380\"},\"name\":null,\"id\":\"3100\\n-2\\n-2.380\"}
	*/
}
