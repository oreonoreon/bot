package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	method string
	Url    string
	values url.Values
}

func (i Info) RequestToYooMoneyApi() []byte {

	i.Url = i.Url + i.method

	client := http.Client{}
	req, err := http.NewRequest("POST", i.Url, strings.NewReader(i.values.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	/*RecievedUrl := resp.Request.URL
	fmt.Println("Получаю ответ ", RecievedUrl)*/

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))

	return result
}

type respOperationHistory struct {
	NextRecord string `json:"next_record"`
	Operations []struct {
		OperationId string    `json:"operation_id"`
		Status      string    `json:"status"`
		PatternId   string    `json:"pattern_id,omitempty"`
		Direction   string    `json:"direction"`
		Label       string    `json:"label"`
		Amount      float64   `json:"amount"`
		Datetime    time.Time `json:"datetime"`
		Title       string    `json:"title"`
		Type        string    `json:"type"`
	} `json:"operations"`
}

func putInType(result []byte) []struct {
	OperationId string    `json:"operation_id"`
	Status      string    `json:"status"`
	PatternId   string    `json:"pattern_id,omitempty"`
	Direction   string    `json:"direction"`
	Label       string    `json:"label"`
	Amount      float64   `json:"amount"`
	Datetime    time.Time `json:"datetime"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
} {
	respOperation := new(respOperationHistory)
	if err := json.Unmarshal(result, respOperation); err != nil {
		log.Println("error in Unmarshal ", err)
	}
	//fmt.Println(respOperation.Operations[0].Amount)
	return respOperation.Operations
}

func (U *Cache) checkPayment(label int64) string {
	u := reqToYoomoneyByLabel(label)
	operation := putInType(u)

	var text string

	if len(operation) == 0 {
		fmt.Printf("%T\n", operation)
		return payTheWitcher
	}
	fmt.Println(operation)
	if operation[0].Status == "success" {
		if *U.Users[label].OperationId == "" || *U.Users[label].OperationId != operation[0].OperationId {
			*U.Users[label].OperationId = operation[0].OperationId
			*U.Users[label].Amount += operation[0].Amount
			if operation[0].Amount > 1 && operation[0].Amount <= 2 {
				fmt.Println("Получил билет за 100")
				text = "Получил билет за 100" + fmt.Sprintf(textGetTicket, operation[0].OperationId)
				U.Users[label].Tickets[sum1] += 1
				fmt.Println("ADD Tickets ", U)
			} else if operation[0].Amount > 900 && operation[0].Amount <= 1000 {
				fmt.Println("Получил билет за 1000")
				text = "Получил билет за 1000" + fmt.Sprintf(textGetTicket, operation[0].OperationId)
				U.Users[label].Tickets[sum2] += 1
			} else if operation[0].Amount > 9000 && operation[0].Amount <= 10000 {
				fmt.Println("Получил билет за 10000")
				text = "Получил билет за 10000" + fmt.Sprintf(textGetTicket, operation[0].OperationId)
				U.Users[label].Tickets[sum3] += 1
			}
			//U.Set(label, operation[0].OperationId, operation[0].Amount)
			return text
		}
	}
	return text
}
func reqToYoomoneyByLabel(label int64) []byte {
	strLabel := strconv.FormatInt(label, 10) //переведём наш label который chatID из цыферек в строку
	u := newUrlValues()
	u.Add("label", strLabel)
	/*u.Set("from", "")
	u.Set("till", "")
	u.Set("start_record", "")
	u.Set("records", "")
	u.Set("details", "")*/

	check := Info{
		method: "api/operation-history",
		Url:    "https://yoomoney.ru/",
		values: u,
	}
	return check.RequestToYooMoneyApi()
}
func newUrlValues() url.Values {
	u := url.Values{}
	u.Add("type", "deposition")
	return u
}
func reqYoomoneyGetAllOperation() []struct {
	OperationId string    `json:"operation_id"`
	Status      string    `json:"status"`
	PatternId   string    `json:"pattern_id,omitempty"`
	Direction   string    `json:"direction"`
	Label       string    `json:"label"`
	Amount      float64   `json:"amount"`
	Datetime    time.Time `json:"datetime"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
} {
	u := newUrlValues()
	u.Add("records", "100")
	check := Info{
		method: "api/operation-history",
		Url:    "https://yoomoney.ru/",
		values: u,
	}
	operation := putInType(check.RequestToYooMoneyApi())

	return operation
}
