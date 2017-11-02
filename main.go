package main

import (
	"github.com/mhvis/planon/planonlib"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	// Read config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Get config values
	twoWayAuthUrl := viper.GetString("twowayauth_url")
	jUsername := viper.GetString("j_username")
	jPassword := viper.GetString("j_password")

	p := planonlib.NewPlanonSession(twoWayAuthUrl, jUsername, jPassword)
	versionDetails, err := p.VersionDetails()
	if err != nil {
		log.Println(err)
	}
	log.Println(versionDetails)
	log.Println(versionDetails.ServerVersion)
	log.Println(versionDetails.Version)

	err = p.GetReservation("3100\\n-2\\n-2.380", 1, time.Now(), time.Now().Add(time.Hour))
	if err != nil {
		log.Println(err)
	}

	err = p.GetRoomWithCode("3100\n-2\n-2.380")
	if err != nil {
		log.Println(err)
	}

}
