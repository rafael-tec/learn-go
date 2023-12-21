package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	baseUrl     = "http://google.com"
	contentType = "application/json"
)

func main() {
	c := http.Client{Timeout: time.Duration(2) * time.Second}

	GetRequestWithContext(c)
	GetRequest(c)
	GetCustomRequest(c)
	PostRequest(c)
}

func GetRequestWithContext(c http.Client) {
	fmt.Println("Sending get request with context...")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func GetCustomRequest(c http.Client) {
	fmt.Println("Sending get custom request...")

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request_id", uuid.NewString())

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func PostRequest(c http.Client) {
	fmt.Println("Sending post request...")

	jsonBuff := bytes.NewBuffer([]byte(`{"name": "Rafael"}`))

	res, err := c.Post(baseUrl, contentType, jsonBuff)
	if err != nil {
		panic(err)
	}

	io.CopyBuffer(os.Stdout, res.Body, nil)
}

func GetRequest(c http.Client) {
	fmt.Println("Sending get request...")

	res, err := c.Get(baseUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
