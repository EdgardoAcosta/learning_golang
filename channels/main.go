package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com.mx",
	}
	// creat a new channel
	c := make(chan string)

	for _, link := range links {
		// run function in a sub-routine/thread
		go checkLink(link, c)
	}

	// Print any msg from the channel, receiving msg is also a blocking statement
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send msg to the channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}
