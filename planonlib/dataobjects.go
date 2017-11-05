package planonlib

import "time"

type VersionDetails struct {
	ServerVersion string `json:"serverversion"`
	Version       string `json:"version"`
}

type Department struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	Reserved bool   `json:"reserved"`
	People   int    `json:"people"`
}

type Person struct {
	Id         string     `json:"id"`
	LastName   string     `json:"last_name"`
	FirstName  string     `json:"first_name"`
	NamePrefix string     `json:"name_prefix"`
	Email      string     `json:"email"`
	Avail      string     `json:"avail"`
	Department Department `json:"department"`
}

type Reservation struct {
	Id     string    `json:"id"`
	Room   Room      `json:"room"`
	Person Person    `json:"person"`
	Mine   bool      `json:"mine"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Type   string    `json:"type"`
	Name   string    `json:"name"`
}

type Property struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

type Floor struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Property Property `json:"property"`
}

// Extraction functions which extract data from a given JSON decoded object.

func extractDepartment(obj interface{}) Department {
	o := obj.(map[string]interface{})
	return Department{
		Id:   asString(o["id"]),
		Name: asString(o["name"]),
	}
}

func extractRoom(obj interface{}) Room {
	o := obj.(map[string]interface{})
	return Room{
		Id:       asString(o["id"]),
		People:   int(o["people"].(float64)),
		Reserved: o["isRes"].(bool),
		Status:   asString(o["status"]),
	}
}

func extractPerson(obj interface{}) Person {
	o := obj.(map[string]interface{})
	return Person{
		Id:         asString(o["id"]),
		Email:      asString(o["ema"]),
		FirstName:  asString(o["fnm"]),
		LastName:   asString(o["name"]),
		NamePrefix: asString(o["prefix"]),
		Department: extractDepartment(o["dep"]),
		Avail:      asString(o["avail"]),
	}
}

func extractReservation(obj interface{}) Reservation {
	o := obj.(map[string]interface{})

	// Extract start/end time
	start, err := time.Parse(planonTime, asString(o["beg"]))
	if err != nil {
		// Todo: maybe nicer error handling
		panic(err)
	}
	end, err := time.Parse(planonTime, asString(o["end"]))
	if err != nil {
		panic(err)
	}

	return Reservation{
		Id:     asString(o["id"]),
		Name:   asString(o["name"]),
		Type:   asString(o["resType"]),
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
		Name:     asString(o["name"]),
		Id:       asString(o["id"]),
		Location: extractLocation(o["loc"]),
	}
}

func extractLocation(obj interface{}) Location {
	o := obj.(map[string]interface{})
	return Location{
		Address:    asString(o["adr"]),
		City:       asString(o["city"]),
		PostalCode: asString(o["pc"]),
	}
}

func extractFloor(obj interface{}) Floor {
	o := obj.(map[string]interface{})
	return Floor{
		Id:       asString(o["id"]),
		Name:     asString(o["name"]),
		Property: extractProperty(o["prop"]),
	}
}

// Tries to convert interface to string. If it fails, an empty string is returned.
func asString(obj interface{}) string {
	str, ok := obj.(string)
	if !ok {
		return ""
	} else {
		return str
	}
}