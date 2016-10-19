package main

import (
	"time"
	"flag"
	 "log"
	"net/http"
)

func main() {

 	urlPtr := flag.String("u", "http://example.com/", "url to perform GET ping")
 	secPtr := flag.Int("s", 3600, "# of seconds between pings")
	flag.Parse()

	//TODO flags validation

	for true {
		resp, err := http.Get(*urlPtr)
		if err != nil {
			log.Println("Error:", err)
		}
		log.Println("Response: ", resp.Status)
				
		time.Sleep(time.Duration(*secPtr) * time.Second)
	}
}
