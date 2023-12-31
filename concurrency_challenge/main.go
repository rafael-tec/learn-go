package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	httpClient = http.Client{}
	defaultCep = "01310300"
)

func main() {
	chBrasilApi := make(chan string)
	chViaCep := make(chan string)

	go fetchAddressInBrasilApi(defaultCep, chBrasilApi)
	go fetchAddressInViaCep(defaultCep, chViaCep)

	select {
	case result := <-chBrasilApi:
		log.Printf("Address by BrasilAPI: %s", result)
	case result := <-chViaCep:
		log.Printf("Address by ViaCEP: %s", result)
	case <-time.After(time.Second * 1):
		log.Fatalln("Timeout was reached and no service returned")
	}
}

func fetchAddressInBrasilApi(cep string, result chan<- string) {
	baseUrl := "https://brasilapi.com.br/api/cep/v1/%s"
	url := fmt.Sprintf(baseUrl, cep)

	log.Printf("Fetching address for CEP %s in BrasilAPI", cep)

	res, err := httpClient.Get(url)
	if err != nil {
		panic(err)
	}

	log.Println("Request to BrasilAPI was successful")

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	result <- string(body)
}

func fetchAddressInViaCep(cep string, result chan<- string) {
	baseUrl := "http://viacep.com.br/ws/%s/json/"
	url := fmt.Sprintf(baseUrl, cep)

	log.Printf("Fetching address for CEP %s in ViaCEP", cep)

	res, err := httpClient.Get(url)
	if err != nil {
		panic(err)
	}

	log.Println("Request to ViaCEP was successful")

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	result <- string(body)
}
