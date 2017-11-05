// Package planon provides high-level and lower-level functions for Planon BookMySpace.
package planonlib

import (
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"time"
	"errors"
	"fmt"
)

// Planon accepts following time format, which is almost the same as ISO8601/RFC3339, however slight different ._.
var planonTime = "2006-01-02T15:04:05-0700"

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

// JSON request/response samples: https://github.com/mhvis/planon/blob/master/planonlib/samples.md


// GetReservations
func (p *Planon) GetReservations(roomId string, start, end time.Time) ([]Reservation, error) {
	// Request
	params := map[string]interface{}{
		"oRoom": map[string]interface{}{
			"id": roomId,
		},
		"start": start.Format(planonTime),
		"end":   end.Format(planonTime),
	}
	var response interface{}
	err := p.Call("/RoomInformation", "getReservation", params, &response)
	if err != nil {
		return nil, err
	}

	// Extra data from response
	obj := response.(map[string]interface{})
	list := obj["list"].([]interface{})

	// Extract reservations
	reservations := make([]Reservation, 0, len(list))
	for _, v := range list {
		// Skip reservations with empty id
		if v.(map[string]interface{})["id"] == nil {
			continue
		}
		// Append reservation to the list
		reservation := extractReservation(v)
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

// ReserveRoom
func (p *Planon) ReserveRoom(roomId, name string, start, end time.Time) error {
	// Request
	params := map[string]interface{}{
		"oRoom": map[string]interface{}{
			"id": roomId,
		},
		"start": start.Format(planonTime),
		"end":   end.Format(planonTime),
	}
	var response bool
	err := p.Call("/RoomInformation", "reserveRoom", params, &response)
	if err != nil {
		return err
	}
	// Check response
	if response != true {
		return errors.New(fmt.Sprintf("unexpected response: %v", response))
	}
	return nil
}

// ExtendReservation
func (p *Planon) ExtendReservation(reservationId string, end time.Time) error {
	// Request
	params := map[string]interface{}{
		"reservation": map[string]interface{}{
			"id": reservationId,
		},
		"end":   end.Format(planonTime),
	}
	var response string
	err := p.Call("/RoomInformation", "extendReservation", params, &response)
	if err != nil {
		return err
	}
	// Check response
	if response != "true" {
		return errors.New("unexpected RPC return value: " + response)
	}
	return nil
}

// EndReservation
func (p *Planon) EndReservation(reservationId string) error {
	// Request
	params := map[string]interface{}{
		"reservation": map[string]interface{}{
			"id": reservationId,
		},
	}
	var response interface{}
	err := p.Call("/RoomInformation", "endReservation", params, &response)
	if err != nil {
		return err
	}
	// Check response
	if response != nil {
		return errors.New("unexpected RPC return value")
	}
	return nil
}

// GetMe
func (p *Planon) GetMe() (Person, error) {
	var response interface{}
	err := p.Call("/UserInformation", "getMe", nil, &response)
	if err != nil {
		return Person{}, err
	}
	return extractPerson(response), nil
}


func (p *Planon) GetRoomWithCode(searchcode string) error {
	params := map[string]interface{}{"searchcode": searchcode}
	err := p.Call("/RoomInformation", "getRoomWithCode", params, nil)
	return err

}

// GetFreeRooms
func (p *Planon) GetFreeRooms(propertyId string) ([]Room, error) {
	params := map[string]interface{}{
		"theProperty": map[string]interface{}{
			"id": propertyId,
		},
	}
	var response map[string]interface{}
	err := p.Call("/FloorInformation", "getFreeRooms", params, &response)
	if err != nil {
		return nil, err
	}

	list := response["list"].([]interface{})
	rooms := make([]Room, len(list), len(list))
	for i, v := range list {
		rooms[i] = extractRoom(v)
	}
	return rooms, nil
}

// GetPropertyList
func (p *Planon) GetPropertyList() ([]Property, error) {
	var response map[string]interface{}
	err := p.Call("/FloorInformation", "getPropertyList", nil, &response)
	if err != nil {
		return nil, err
	}

	list := response["list"].([]interface{})
	properties := make([]Property, len(list), len(list))
	for i, v := range list {
		properties[i] = extractProperty(v)
	}
	return properties, nil
}

// GetFloorsOfProperty
func (p *Planon) GetFloorsOfProperty(propertyId string) ([]Floor, error) {
	params := map[string]interface{}{
		"theProperty": map[string]interface{}{
			"id": propertyId,
		},
	}

	var response map[string]interface{}
	err := p.Call("/FloorInformation", "getFloorsOfProperty", params, &response)
	if err != nil {
		return nil, err
	}

	list := response["list"].([]interface{})
	floors := make([]Floor, len(list), len(list))
	for i, v := range list {
		floors[i] = extractFloor(v)
	}
	return floors, nil
}

// GetVersionDetails
func (p *Planon) GetVersionDetails() (VersionDetails, error) {
	var versionDetails VersionDetails
	err := p.Call("/VersionCheck", "getVersionDetails", nil, &versionDetails)
	return versionDetails, err
}