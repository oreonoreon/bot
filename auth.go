package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func Connect() {
	u := url.Values{}
	u.Set("client_id", "716F2AD40A075568A2A750C572AE9E897EEB2FAE5DBA77483BB92932743ABD87")
	u.Set("response_type", "code")
	u.Set("redirect_uri", "https://tglink.ru/test_app1")
	u.Set("scope", "account-info operation-history operation-details incoming-transfers payment-p2p")
	//u.Set("instance_name", "")
	Url := "https://yoomoney.ru/oauth/authorize"
	resp, err := http.PostForm(Url, u)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("URL is ", resp.Request.URL)

	resp.Body.Close()
}
func getToken() {
	u2 := url.Values{}
	u2.Set("code", "4E9C1E2122789FBED698C42FF7B7270873C0DF33110776814A5B00C7D107D0247FDB58E3BD53E4DAB9C4AFAD05AC4BF9E62CAA09BCC21E099DA6CD1E3ECE916F21F7608D50511847DC8D5922257102EE264C4EE1A61208D6967074ACA654AEA874536BE109F453F70F2878AB85C2356D4791BE9478029579B623CD100C002D4C")
	u2.Set("client_id", "716F2AD40A075568A2A750C572AE9E897EEB2FAE5DBA77483BB92932743ABD87")
	u2.Set("grant_type", "authorization_code")
	u2.Set("redirect_uri", "https://tglink.ru/test_app1")
	Url := "https://yoomoney.ru/oauth/token"
	resp, err := http.PostForm(Url, u2)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}
