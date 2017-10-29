package main

import (
	"github.com/mhvis/planon/planon"
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
	jsonRpcUrl := viper.GetString("json_rpc_url")
	jUsername := viper.GetString("j_username")
	jPassword := viper.GetString("j_password")
	jSecurityCheckUrl := viper.GetString("j_security_check_url")

	p := planon.NewPlanonSession(jsonRpcUrl, jSecurityCheckUrl, jUsername, jPassword)
	versionDetails, err := p.VersionDetails()
	if err != nil {
		log.Println(err)
	}
	log.Println(versionDetails)

	err = p.GetReservation(time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}

	err = p.GetRoomWithCode("3100\n-2\n-2.380")
	if err != nil {
		log.Println(err)
	}

}
