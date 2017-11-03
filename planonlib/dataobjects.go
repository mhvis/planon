package planonlib

import "time"

type VersionDetails struct {
	ServerVersion string `json:"serverversion"`
	Version       string `json:"version"`
}

type Department struct {
	Id   string
	Name string
}

type Room struct {
	Id       string
	Status   string
	Reserved bool
	People   int
}

type Person struct {
	Id         string
	LastName   string
	FirstName  string
	NamePrefix string
	Email      string
	Avail      string
	Department
}

type Reservation struct {
	Id    string
	Room
	Person
	Mine  bool
	Start time.Time
	End   time.Time
	Type  string
	Name  string
}

type Property struct {
	Id   string
	Name string
	Location
}

type Location struct {
	Address    string
	City       string
	PostalCode string
}

type Floor struct {
	Id   string
	Name string
	Property
}

// Extraction functions which extract data from a given JSON decoded object.

func extractDepartment(obj interface{}) Department {
	o := obj.(map[string]interface{})
	return Department{
		Id:   o["id"].(string),
		Name: o["name"].(string),
	}
}

func extractRoom(obj interface{}) Room {
	o := obj.(map[string]interface{})
	return Room{
		Id:       o["id"].(string),
		People:   o["people"].(int),
		Reserved: o["isRes"].(bool),
		Status:   o["status"].(string),
	}
}

func extractPerson(obj interface{}) Person {
	o := obj.(map[string]interface{})
	return Person{
		Id:         o["id"].(string),
		Email:      o["ema"].(string),
		FirstName:  o["fnm"].(string),
		LastName:   o["name"].(string),
		NamePrefix: o["prefix"].(string),
		Department: extractDepartment(o["dep"]),
		Avail:      o["avail"].(string),
	}
}

func extractReservation(obj interface{}) Reservation {
	o := obj.(map[string]interface{})

	// Extract start/end time
	start, err := time.Parse(time.RFC3339, o["beg"].(string))
	if err != nil {
		// Todo: maybe nicer error handling
		panic(err)
	}
	end, err := time.Parse(time.RFC3339, o["end"].(string))
	if err != nil {
		panic(err)
	}

	return Reservation{
		Id:     o["id"].(string),
		Name:   o["name"].(string),
		Type:   o["resType"].(string),
		Start:  start,
		End:    end,
		Mine:   o["myRes"].(bool),
		Person: extractPerson(o["pers"]),
		Room:   extractRoom(o["locCode"]),
	}
}

func extractProperty(obj interface{}) Property {
	o := obj.(map[string]interface{})
	return Property{
		Name:     o["name"].(string),
		Id:       o["id"].(string),
		Location: extractLocation(o["loc"]),
	}
}

func extractLocation(obj interface{}) Location {
	o := obj.(map[string]interface{})
	return Location{
		Address:    o["adr"].(string),
		City:       o["city"].(string),
		PostalCode: o["pc"].(string),
	}
}

func extractFloor(obj interface{}) Floor {
	o := obj.(map[string]interface{})
	return Floor{
		Id:       o["id"].(string),
		Name:     o["name"].(string),
		Property: extractProperty(o["prop"]),
	}
}
