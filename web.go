package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mhvis/planon/planonrpc"
	"net/http"
	"time"
)

func listenAndServe(apiKeys []string, p *planonrpc.Service) {
	r := gin.Default()

	r.Use(auth(apiKeys))

	r.POST("/getReservations", getReservations(p))
	r.POST("/reserveRoom", reserveRoom(p))
	r.POST("/endReservation", endReservation(p))
	r.POST("/extendReservation", extendReservation(p))
	r.POST("/getPropertyList", getPropertyList(p))
	r.POST("/getFloorsOfProperty", getFloorsOfProperty(p))
	r.POST("/getFreeRooms", getFreeRooms(p))
	r.POST("/getMe", getMe(p))

	r.Run()
}

// Handlers

func getReservations(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		start, ok := parseTime(c, c.PostForm("start"))
		if !ok {
			return
		}
		end, ok := parseTime(c, c.PostForm("end"))
		if !ok {
			return
		}

		reservations, err := p.GetReservations(c.PostForm("room_id"), start, end)
		setResponse(c, reservations, err)
	}
}

func reserveRoom(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		start, ok := parseTime(c, c.PostForm("start"))
		if !ok {
			return
		}
		end, ok := parseTime(c, c.PostForm("end"))
		if !ok {
			return
		}

		err := p.ReserveRoom(c.PostForm("room_id"), c.PostForm("name"), start, end)
		setResponse(c, nil, err)
	}
}

func endReservation(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		start, ok := parseTime(c, c.PostForm("start"))
		if !ok {
			return
		}
		end, ok := parseTime(c, c.PostForm("end"))
		if !ok {
			return
		}

		err := p.EndReservation(c.PostForm("room_id"), start, end)
		setResponse(c, nil, err)
	}
}

func getPropertyList(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		properties, err := p.GetPropertyList()
		setResponse(c, properties, err)
	}
}

func extendReservation(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		end, ok := parseTime(c, c.PostForm("end"))
		if !ok {
			return
		}
		err := p.ExtendReservation(c.PostForm("reservation_id"), end)
		setResponse(c, nil, err)
	}
}

func getFloorsOfProperty(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		floors, err := p.GetFloorsOfProperty(c.PostForm("property_id"))
		setResponse(c, floors, err)
	}
}

func getFreeRooms(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms, err := p.GetFreeRooms(c.PostForm("property_id"))
		setResponse(c, rooms, err)
	}
}

func getMe(p *planonrpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		person, err := p.GetMe()
		setResponse(c, person, err)
	}
}

// Middleware & helper functions

func auth(validKeys []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")

		// Check if key is valid
		for _, validKey := range validKeys {
			if key == validKey {
				return
			}
		}

		// If not valid, abort response
		c.Abort()
		c.Status(http.StatusUnauthorized)
	}
}

// Tries to parse the given time string as RFC3339.
// If it fails, it updates the response and the second bool will be false.
func parseTime(c *gin.Context, s string) (time.Time, bool) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return t, false
	}
	return t, true
}

// Sets the response body with the given data or error and sets the status code accordingly.
func setResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.String(http.StatusTeapot, err.Error())
	} else if data != nil {
		c.IndentedJSON(http.StatusOK, data)
	} else {
		c.Status(http.StatusNoContent)
	}
}
