package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/viper"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"
)

func main() {
	var err error

	// Read config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Get config values
	baseUrl := viper.GetString("baseurl")
	username := viper.GetString("username")
	password := viper.GetString("password")

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt, cdp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, planonLogin(baseUrl, username, password))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	//	log.Printf("saved screenshot from search result listing `%s` (%s)", res, site)
}

// Task list for Planon login
func planonLogin(url, username, password string) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(url),
		cdp.SendKeys()
		cdp.Sleep(30 * time.Second),
	}
}

func googleSearch(q, text string, site, res *string) cdp.Tasks {
	var buf []byte
	sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
	return cdp.Tasks{
		cdp.Navigate(`https://www.google.com`),
		cdp.WaitVisible(`#hplogo`, cdp.ByID),
		cdp.SendKeys(`#lst-ib`, q+"\n", cdp.ByID),
		cdp.WaitVisible(`#res`, cdp.ByID),
		cdp.Text(sel, res),
		cdp.Click(sel),
		cdp.WaitVisible(`a[href="/brankas-for-business"]`, cdp.ByQuery),
		cdp.WaitNotVisible(`.preloader-content`, cdp.ByQuery),
		cdp.Location(site),
		cdp.ScrollIntoView(`.banner-section.third-section`, cdp.ByQuery),
		cdp.Sleep(2 * time.Second), // wait for animation to finish
		cdp.Screenshot(`.banner-section.third-section`, &buf, cdp.ByQuery),
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			return ioutil.WriteFile("screenshot.png", buf, 0644)
		}),
	}
}
