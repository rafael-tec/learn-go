package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AddressResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in http request: %v\n", err)
		}

		defer res.Body.Close()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in reade response: %v\n", err)
		}

		var address AddressResponse
		err = json.Unmarshal(resBody, &address)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in parse response: %v\n", err)
		}

		fmt.Println(address)

		file, err := os.Create("via_cep_response.json")
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in file creation: %v\n", err)
		}

		defer file.Close()
		_, err = file.WriteString(string(resBody))
	}
}
