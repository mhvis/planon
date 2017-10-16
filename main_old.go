package main

import (
	_ "net/http"
	"net/url"
)
import _ "fmt"
import "log"
import _ "io/ioutil"
import "github.com/headzoo/surf"
import "github.com/headzoo/surf/browser"
import "github.com/spf13/viper"
import "bytes"

func main_old() {
	// Read config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Get config values
	baseUrl := viper.GetString("baseurl")
	username := viper.GetString("username")
	password := viper.GetString("password")

	// Go and surf
	bow := surf.NewBrowser()

	// SAML page 1
	err = bow.Open(baseUrl)
	if err != nil {
		panic(err)
	}
	//log.Println(htmlString(bow))
	fm, err := bow.Form("#loginForm")
	if err != nil {
		panic(err)
	}
	fm.Input("UserName", username)
	fm.Input("Password", password)

	// SAML page 2
	err = fm.Submit()
	if err != nil {
		panic(err)
	}
	//log.Println(htmlString(bow))
	fm, err = bow.Form("form")
	if err != nil {
		panic(err)
	}

	// SAML page 3
	err = fm.Submit()
	if err != nil {
		panic(err)
	}
	//log.Println(htmlString(bow))
	fm, err = bow.Form("#id1")
	if err != nil {
		panic(err)
	}
	// SAML page 4
	err = fm.Submit()
	if err != nil {
		panic(err)
	}
	//log.Println(htmlString(bow))

	// Logged in!

	// POST
	v := url.Values{}
	v.Set("buildingBlocks:inlineBuildingBlocks:block2:subCaseParts:inlineSubCaseParts:0:casepart:buildingBlockGroups:group:form:buildingBlocks:inlineBuildingBlocks:block1:form:fieldGroupRepeater:0:fieldGroup:fieldRepeater:1:field:fieldContent:field-editor:datetimepicker",
		"12-10-2017 21:34")
	v.Set("buildingBlocks:inlineBuildingBlocks:block2:subCaseParts:inlineSubCaseParts:0:casepart:buildingBlockGroups:group:form:buildingBlocks:inlineBuildingBlocks:block1:form:fieldGroupRepeater:0:fieldGroup:fieldRepeater:2:field:fieldContent:field-editor:datetimepicker",
		"12-10-2017 22:34")
	v.Set("buildingBlocks:inlineBuildingBlocks:block2:subCaseParts:inlineSubCaseParts:0:casepart:buildingBlockGroups:group:form:buildingBlocks:inlineBuildingBlocks:block1:form:fieldGroupRepeater:1:fieldGroup:fieldRepeater:1:field:fieldContent:field-editor:datetimepicker",
		"3100, Luna")
	v.Set("buildingBlocks:inlineBuildingBlocks:block2:subCaseParts:inlineSubCaseParts:0:casepart:buildingBlockGroups:group:form:buildingBlocks:inlineBuildingBlocks:block1:form:fieldGroupRepeater:1:fieldGroup:fieldRepeater:5:field:fieldContent:field-editor:textfield",
		"1")

	err = bow.PostForm("https://agnes.campus.tue.nl:18443/case/tue/RES_001", v)
	checkErr(err)

	log.Println(htmlString(bow))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// htmlString returns the HTML code of the current page in the browser.
func htmlString(bow *browser.Browser) string {
	var b bytes.Buffer

	_, err := bow.Download(&b)
	if err != nil {
		panic(err)
	}

	return b.String()
}
