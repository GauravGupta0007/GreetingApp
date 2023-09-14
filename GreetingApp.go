package main

import (
	"fmt"
	"time"
)

func main() {
	var cur_time = time.Now().Hour()

	var greeting string

	if cur_time >= 6 && cur_time <= 11 {
		greeting = "Good morning!"
	} else if cur_time >= 11 && cur_time <= 16 {
		greeting = "Good afternoon!"
	} else if cur_time >= 16 && cur_time <= 21 {
		greeting = "Good evening!"
	} else {
		greeting = "Good night!"
	}

	fmt.Println(greeting)
}