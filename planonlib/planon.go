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

func (p *Planon) GetReservation(id string, start, end time.Time) error {
	params := map[string]interface{}{
		"oRoom": map[string]interface{}{
			"id": id,
		},
		"start": start.Format(time.RFC3339),
		"end":   end.Format(time.RFC3339),
	}

	result := make(map[string]interface{})
	err := p.Call("/RoomInformation", "getReservation", params, &result)
	log.Println(result)
	return err
}

func (p *Planon) GetRoomWithCode(searchcode string) error {
	params := map[string]interface{}{"searchcode": searchcode}
	err := p.Call("/RoomInformation", "getRoomWithCode", params, nil)
	return err

}
