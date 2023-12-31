package main

import (
	"fmt"
	"net/http"
	"time"
)

// Go Routines📝: Go Routines help in concurrent programming.
// 		When a go program runs it executes the main routine (threat) and executes from top-to-bottom order. And can not execute anything else until one code block is done executing.
// 		Go Routines: help with concurrent programming (running multiple threads at once).

// Go Channel📝: In general main go routine does not care if a child routine is done processing and we might not be able to track the progress of these concurrent programs,
// 		that's where go channels come into place, channels create link between main and child routines(threads) so that main routine is aware of the child routine status.
// 		One go channel can only communicate through one data type, meaning channel of int will not communicate string data.

// Function LIteral📝: have the same behaviour as anonymous functions in JavaScript. In go function literals are unnamed functions that gets wrapped with some chunk of codes
// 		which could be executed at some point in future.

var website_up []string
var website_down []string

func main() {
	fmt.Println("Hello from Go Routines!!")

	websites := []string{
		"http://google.com1",
		"http://fb.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://instagram.com",
	}

	// Creating a Go Channel to communicate between main and child Go Routines ⬇️
	channel := make(chan string)

	for _, web := range websites {
		fmt.Println("Checking:", web)
		// checkWebsite(web)
		// Running go routine to do concurrent programming ⏩
		go checkWebsite(web, channel)
	}

	// Receiving message from Channel
	// fmt.Println("Channel message received: " + <-channel)
	// fmt.Println("Channel message received: " + <-channel)
	// fmt.Println("Channel message received: " + <-channel)
	// fmt.Println("Channel message received: " + <-channel)
	// fmt.Println("Channel message received: " + <-channel)

	// Receiving Channel messages using loop 🔽
	for i := 0; i < len(websites); i++ {
		fmt.Println("Channel message received => " + <-channel)
	}

	// Adding condition to keep static check on websites, continuously keep making the website to check status 🔽1️⃣
	// for {
	// 	go checkWebsite(<-channel, channel)
	// }

	// Another of running this for loop 2️⃣
	for site := range channel {
		// time.Sleep(time.Second * 2)
		// go checkWebsite(site, channel)

		// Replacing the direct function call with function literals 🔽
		go func(web_site string) {
			time.Sleep(time.Second * 2)
			checkWebsite(web_site, channel)
		}(site) // here we pass the site as arg to this function literal
	}

	// fmt.Println("Waiting for Go Routines to finish ..")
	// time.Sleep(time.Second)

	// fmt.Println("Websites up:", website_up)
	// fmt.Println("Websites down:", website_down)
}

func checkWebsite(web string, channel chan string) {
	_, err := http.Get(web)
	if err != nil {
		// fmt.Println("Error:", err)
		fmt.Println(web + " is down")
		// Sending messsage to Channel
		// channel <- web + " might be down"
		channel <- web
		// website_down = append(website_down, web)
		return
	}
	fmt.Println(web + " is up")
	// website_up = append(website_up, web)..
	// channel <- web + " is up"
	channel <- web
}
