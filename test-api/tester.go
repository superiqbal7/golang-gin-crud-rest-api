package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/api/videos"
	method := "POST"

	payload := strings.NewReader(`{
    "title": "cool Golang / Go Gin Framework Crash Course 04 | HTML, Templates and Multi-Route Grouping",
    "description": "In this video we are going to take a look at HTML, Templates and Multi-Route Grouping using Golang's Gin HTTP Framework.",
    "url": "https://youtu.be/sDJLQMZzzM4?list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w",
    "author": {
        "firstname": "rahim",
        "lastname": "iqbal",
        "age": 24,
        "email": "aa@a.com"
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
