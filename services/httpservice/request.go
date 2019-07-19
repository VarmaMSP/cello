package httpservice

import (
	"fmt"
	"net/http"
)

func (client *Client) MakeRequests(urlChan <-chan string, respChan chan<- *http.Response) {
	semaphore := make(chan int, 7)

	for {
		go func(url string) {
			semaphore <- 0
			fmt.Printf("CRAWLING: %s\n", url)
			req, _ := http.NewRequest("GET", url, nil)
			resp, err := client.httpClient.Do(req)
			if err != nil {
				fmt.Printf("REQ FAILED: GET %s - %s\n\n", url, err.Error())
			}
			<-semaphore
			respChan <- resp
		}(<-urlChan)
	}
}
