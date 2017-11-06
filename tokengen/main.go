package main

import (
	"encoding/base64"
	"fmt"
	"github.com/mhvis/planon/jsonrpc"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	a := os.Args
	if len(a) < 3 {
		fmt.Printf("Usage: %v <ticket> <new token>\n", a[0])
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

	ticket := a[1]
	token := a[2]
	domain, username, password := decodeTicket(ticket)
	fmt.Println("domain:", domain)
	fmt.Println("username:", username)
	fmt.Println("password:", password)

	// Do getMe request
	req, err := jsonrpc.EncodeRequest(domain+"/JSONrpc/UserInformation", "getMe", nil, 0)
	if err != nil {
		panic(err)
	}
	// Print response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("getMe response:", resp)

	// Do login request
	resp, err = client.PostForm(domain+"/j_security_check", url.Values{
		"j_username": {username},
		"j_password": {password + token},
	})
	if err != nil {
		panic(err)
	}
	// Print response
	fmt.Println("login response:", resp, getBody(resp))

	// Do logout
	req, err = jsonrpc.EncodeRequest(domain+"/JSONrpc/UserInformation", "logout", nil, 0)
	if err != nil {
		panic(err)
	}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	// Print response
	fmt.Println("logout response:", resp, getBody(resp))

	// Try another login, now with token
	req, err = jsonrpc.EncodeRequest(domain+"/JSONrpc/UserInformation", "getMe", nil, 0)
	if err != nil {
		panic(err)
	}
	// Print response
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("getMe response:", resp)

	// Do login request
	resp, err = client.PostForm(domain+"/j_security_check", url.Values{
		"j_username": {username},
		"j_password": {token},
	})
	if err != nil {
		panic(err)
	}
	// Print response
	fmt.Println("second login response:", resp, getBody(resp))

}

// Takes the useful information out of a ticket (code based on Planon APK function 'decodeResultAndLoginTwoWayAuth'
func decodeTicket(ticket string) (domain, username, password string) {
	decoded := decodeBase64(ticket)
	values := strings.Split(decoded, "@ @")
	domain = values[0]
	username = decodeBase64(values[1])
	password = values[2]
	return
}

func decodeBase64(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getBody(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	return string(body)
}
