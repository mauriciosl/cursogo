package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func check(url string) {
	resp, err := http.Get(url)
	if err == nil && resp.StatusCode == 200 {
		fmt.Println(url, "OK")
		return
	}
	fmt.Println(url, "NOK")
}

func main() {
	urls := []string{
		"http://janus.hotelurbano.com/healthcheck",
		"http://novabusca.hotelurbano.com/healthcheck",
		"http://hotels.hotelurbano.com/healthcheck",
	}
	for _, url := range urls {
		check(url)
	}
}

// END OMIT
