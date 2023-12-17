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
	http.HandleFunc("/", GetCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func GetCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid URL!\n"))
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP is param required!\n"))
		return
	}

	address, err := SearchCEP(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func SearchCEP(cep string) (*AddressResponse, error) {
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

	return &address, nil
}
