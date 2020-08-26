package main

import (
	"github.com/mhvis/planon/planonrpc"
	"github.com/spf13/viper"
	"net/http"
	"crypto/tls"
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
	serverName := viper.GetString("server_name")

	if serverName != "" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ServerName: serverName}
	}

	p := planonrpc.NewService(twowayauthUrl, jUsername, jPassword)

	listenAndServe(apiKeys, p)
}
