// Package planon provides high-level and lower-level functions for Planon BookMySpace.
package planonrpc

import (
	"errors"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// ~~Planon accepts following time format, which is almost the same as ISO8601/RFC3339, however slight different ._.~~
// NEW PLANON VERSION USES RFC3339!!!
var planonTime = time.RFC3339

// Service stores data for a web service connection.
type Service struct {
	client        *http.Client
	id            int
	twowayauthUrl string
	jUsername     string
	jPassword     string
}

func NewService(twowayauthUrl, jUsername, jPassword string) *Service {
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
	return &Service{
		client,
		0,
		twowayauthUrl,
		jUsername,
		jPassword,
	}
}

// JSON request/response samples: https://github.com/mhvis/planon/blob/master/planonlib/samples.md

// GetReservations
func (p *Service) GetReservations(roomId string, start, end time.Time) ([]Reservation, error) {
	// Request
	args := []string{
		fmt.Sprintf("{\"id\":\"%s\"}", roomId),
		fmt.Sprintf("\"%s\"", start.Format(planonTime)),
		fmt.Sprintf("\"%s\"", end.Format(planonTime)),
	}
	var response interface{}
	err := p.Call("/RoomInformation", "getReservation", nil, args, &response)
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
func (p *Service) ReserveRoom(roomId, name string, start, end time.Time) error {
	// Request
	args := []string{
		fmt.Sprintf("{\"id\":\"%s\"}", roomId),
		fmt.Sprintf("\"%s\"", start.Format(planonTime)),
		fmt.Sprintf("\"%s\"", end.Format(planonTime)),
	}
	var response bool
	err := p.Call("/RoomInformation", "reserveRoom", nil, args, &response)
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
func (p *Service) ExtendReservation(reservationId string, end time.Time) error {
	return errors.New("this method is not updated to the new Planon API")

	// Request
	params := map[string]interface{}{
		"reservation": map[string]interface{}{
			"id": reservationId,
		},
		"end": end.Format(planonTime),
	}
	var response string
	err := p.Call("/RoomInformation", "extendReservation", params, []string{}, &response)
	if err != nil {
		return err
	}
	// Check response
	if response != "true" {
		return errors.New("unexpected RPC return value: " + response)
	}
	return nil
}

// EndReservation, a bit hacky
func (p *Service) EndReservation(roomId string, start, end time.Time) error {
	// TODO: clean up this function

	// TODO: below is duplication of GetReservations!
	// Request
	args := []string{
		fmt.Sprintf("{\"id\":\"%s\"}", roomId),
		fmt.Sprintf("\"%s\"", start.Format(planonTime)),
		fmt.Sprintf("\"%s\"", end.Format(planonTime)),
	}
	var response interface{}
	err := p.Call("/RoomInformation", "getReservation", nil, args, &response)
	if err != nil {
		return err
	}

	// Extra data from response
	obj := response.(map[string]interface{})
	reservationList := obj["list"].([]interface{})

	// Extract reservations
	reservations := make([]Reservation, 0, len(reservationList))
	for _, v := range reservationList {
		// Skip reservations with empty id
		if v.(map[string]interface{})["id"] == nil {
			continue
		}
		// Append reservation to the list
		reservation := extractReservation(v)
		reservations = append(reservations, reservation)
	}
	// END OF DUPLICATION

	// Search for reservation with equal start and end
	resIdx := -1
	for i, res := range reservations {
		if res.Start.Equal(start) && res.End.Equal(end) {
			resIdx = i
			break
		}
	}
	if resIdx == -1 {
		return errors.New("Planon reservation not found")
	}

	// FETCH ROOM
	args = []string{
		fmt.Sprintf("\"%s\"", roomId),
	}
	var room interface{}
	err = p.Call("/RoomInformation", "getRoomWithCode", nil, args, &room)
	if err != nil {
		return err
	}

	// End reservation using the fetched reservation data
	roomEncoded, err := json.Marshal(room)
	if err != nil {
		panic(err)
	}
	reservationEncoded, err := json.Marshal(reservationList[resIdx])
	if err != nil {
		panic(err)
	}
	args = []string{
		string(roomEncoded),
		string(reservationEncoded),
	}
	err = p.Call("/RoomInformation", "endReservation", nil, args, &response)
	if err != nil {
		return err
	}
	// Discard response
	//if response != nil {
	//	return errors.New("unexpected Planon response")
	//}
	return nil
}

// GetMe
func (p *Service) GetMe() (Person, error) {
	var response interface{}
	err := p.Call("/UserInformation", "getMe", nil, []string{}, &response)
	if err != nil {
		return Person{}, err
	}
	return extractPerson(response), nil
}

func (p *Service) GetRoomWithCode(searchcode string) error {
	return errors.New("this method is not updated to the new Planon API")

	params := map[string]interface{}{"searchcode": searchcode}
	err := p.Call("/RoomInformation", "getRoomWithCode", params, []string{}, nil)
	return err

}

// GetFreeRooms
func (p *Service) GetFreeRooms(propertyId string) ([]Room, error) {
	return nil, errors.New("this method is not updated to the new Planon API")

	params := map[string]interface{}{
		"theProperty": map[string]interface{}{
			"id": propertyId,
		},
	}
	var response map[string]interface{}
	err := p.Call("/FloorInformation", "getFreeRooms", params, []string{}, &response)
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
func (p *Service) GetPropertyList() ([]Property, error) {
	var response map[string]interface{}
	err := p.Call("/FloorInformation", "getPropertyList", nil, []string{}, &response)
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
func (p *Service) GetFloorsOfProperty(propertyId string) ([]Floor, error) {
	return nil, errors.New("this method is not updated to the new Planon API")

	params := map[string]interface{}{
		"theProperty": map[string]interface{}{
			"id": propertyId,
		},
	}

	var response map[string]interface{}
	err := p.Call("/FloorInformation", "getFloorsOfProperty", params, []string{}, &response)
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
func (p *Service) GetVersionDetails() (VersionDetails, error) {
	var versionDetails VersionDetails
	err := p.Call("/VersionCheck", "getVersionDetails", nil, []string{}, &versionDetails)
	return versionDetails, err
}
