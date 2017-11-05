package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/mhvis/planon/planonlib"
	"net/http"
)

func listenAndServe(apiKeys []string, p *planonlib.Planon) {
	r := gin.Default()

	r.Use(auth(apiKeys))

	r.GET("/reservations", getReservations(p))
	r.POST("/reservations", createReservation(p))
	r.DELETE("/reservations", deleteReservation(p))

	r.GET("/properties", getProperties(p))

	r.Run()
}

func getReservations(p *planonlib.Planon) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomId := c.Query("room_id")
		start, ok := parseTime(c, c.Query("start"))
		if !ok {
			return
		}
		end, ok := parseTime(c, c.Query("end"))
		if !ok {
			return
		}
		reservations, err := p.GetReservations(roomId, start, end)
		if err != nil {
			c.String(http.StatusTeapot, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, reservations)
	}
}

func createReservation(p *planonlib.Planon) gin.HandlerFunc {
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
		if err != nil {
			c.String(http.StatusTeapot, err.Error())
			return
		}
		c.String(http.StatusCreated, "Created")
	}
}

func deleteReservation(p *planonlib.Planon) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := p.EndReservation(c.Query("reservation_id"))
		if err != nil {
			c.String(http.StatusTeapot, err.Error())
			return
		}
		c.String(http.StatusNoContent, "Deleted")
	}
}

func getProperties(p *planonlib.Planon) gin.HandlerFunc {
	return func(c *gin.Context) {
		properties, err := p.GetPropertyList()
		if err != nil {
			c.String(http.StatusTeapot, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, properties)
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
		c.String(http.StatusUnauthorized, "Unauthorized")
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
