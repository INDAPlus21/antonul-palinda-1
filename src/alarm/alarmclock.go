package main

import (
	"fmt"
	"time"
)

const (
	hhmmss = "15:04:05"
)

func currentTime() string {
	return "The time is: " + time.Now().Format(hhmmss)
}

func timeTo(todo string) {
	switch todo {
	case "eat":
		for {
			fmt.Println(currentTime() + ": Time to eat")
			time.Sleep(10 * time.Second)
		}
	case "work":
		for {
			fmt.Println(currentTime() + ": Time to work")
			time.Sleep(30 * time.Second)
		}
	case "sleep":
		for {
			fmt.Println(currentTime() + ": Time to sleep")
			time.Sleep(60 * time.Second)
		}
	default:
		fmt.Println("U fail, u bad")
	}
}

func main() {
	go timeTo("eat")
	go timeTo("work")
	go timeTo("sleep")
	select {}
}
