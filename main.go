package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	IPV4 := "65.204.3.85"
	UAString := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.37"

	payload := []byte(fmt.Sprintf(`{"ip":"%s","ua_string":"%s"}`, IPV4, UAString))

	req, err := http.NewRequest("POST", "http://localhost:14000/audience/v1/connections/nocreate", bytes.NewBuffer(payload))
	check(err)

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
