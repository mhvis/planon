package main

import (
	"github.com/mhvis/planon/planonlib"
	"github.com/spf13/viper"
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
	apiKeys := viper.GetStringSlice("api_keys")

	p := planonlib.NewPlanonSession(twowayauthUrl, jUsername, jPassword)

	//reservations, err := p.GetReservations("3100\n-2\n-2.380", time.Now(), time.Now().Add(48*time.Hour))
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	log.Println(reservations)
	//}

	listenAndServe(apiKeys, p)
}
