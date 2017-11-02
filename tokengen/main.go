package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
	"net/http/cookiejar"
	"golang.org/x/net/publicsuffix"
	"net/url"
	"io/ioutil"
)

/*
02406edb31a-2864-466b-89e1-9c2ce53a8feb
024645101c7-bfe0-463a-93cd-245e1930f94a

 */

func main() {
	a := os.Args
	if len(a) < 5 {
		fmt.Printf("Usage: %v <twowayauth url> <username> <uuid> <new token>\n", a[0])
		os.Exit(2)
	}

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

	// Get variables
	authUrl := a[1] + "/j_security_check"
	username := a[2]
	uuid := a[3]
	token := a[4]

	password := uuid + token

	// Do request
	resp, err := client.PostForm(authUrl, url.Values{"j_username": {username}, "j_password": {password}})
	if err != nil {
		panic(err)
	}

	// Examine/print response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	fmt.Println("Response:", resp)
	fmt.Println("Response body:", string(body))
}