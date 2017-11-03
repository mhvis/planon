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
	twowayauthUrl := viper.GetString("twowayauth_url")
	jUsername := viper.GetString("j_username")
	jPassword := viper.GetString("j_password")

	p := planonlib.NewPlanonSession(twowayauthUrl, jUsername, jPassword)
	//versionDetails, err := p.VersionDetails()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(versionDetails)
	//log.Println(versionDetails.ServerVersion)
	//log.Println(versionDetails.Version)

	wednesdayMorning := time.Date(2017, 10, 29, 8, 0, 0, 0, time.Local)
	wednesdayEvening := wednesdayMorning.Add(80*time.Hour)

	err = p.GetReservation("3100\n-2\n-2.386", wednesdayMorning, wednesdayEvening)
	if err != nil {
		log.Println(err)
	}

	//err = p.GetRoomWithCode("3100\n-2\n-2.380")
	//if err != nil {
	//	log.Println(err)
	//}

}
