package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

type User struct {
	messageId   *int
	Card        *string
	OperationId *string
	Amount      *float64
	Tickets     map[string]int
	time        *time.Time
	FirstName   string
	LastName    string
	UserName    string
}

//var Tickets = map[string]int{sum1: 0, sum2: 0, sum3: 0}

type Cache struct {
	Users map[int64]User
}

func New() *Cache {
	users := make(map[int64]User)

	cache := Cache{
		Users: users,
	}
	return &cache
}
func (U *Cache) CheckNewUser(update tgbotapi.Update) {

	if _, ok := U.Users[update.SentFrom().ID]; !ok {
		user := User{
			messageId: new(int), Card: new(string),
			OperationId: new(string), Amount: new(float64),
			Tickets:   map[string]int{sum1: 0, sum2: 0, sum3: 0},
			time:      new(time.Time),
			FirstName: update.SentFrom().FirstName,
			LastName:  update.SentFrom().LastName,
			UserName:  update.SentFrom().UserName,
		}
		U.Users[update.SentFrom().ID] = user
		fmt.Println(U, ok)
	} else {
		fmt.Println("IT'S ALIVE! ", U, ok)
	}

}

func (U *Cache) Get(userID int64) string {
	return *U.Users[userID].OperationId
}
func (U *Cache) getTickets(userID int64) string {
	var str string
	for k, v := range U.Users[userID].Tickets {
		str += strconv.Itoa(v) + biletZa + k + "\n"
	}

	return str
}

func (U *Cache) FullFillCache(ID int64) {

	if _, ok := U.Users[ID]; !ok {
		user := User{
			messageId: new(int), Card: new(string),
			OperationId: new(string), Amount: new(float64),
			Tickets:   map[string]int{sum1: 0, sum2: 0, sum3: 0},
			time:      new(time.Time),
			FirstName: "",
			LastName:  "",
			UserName:  "",
		}
		U.Users[ID] = user
		fmt.Println(U, ok)
		fmt.Println(U.Users[ID].time)
	} else {
		fmt.Println("IT'S ALIVE! ", U, ok)
	}

}
