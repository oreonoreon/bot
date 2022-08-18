package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var bank int

func bankAddFakeMoney() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(5)
	switch random {
	case 0, 1:
		s, _ := strconv.Atoi(sum1)
		bank += s
	case 2, 3:
		s, _ := strconv.Atoi(sum2)
		bank += s
	case 4:
		s, _ := strconv.Atoi(sum3)
		bank += s
	}
	fmt.Printf("Random number is %v\n", random)
}
