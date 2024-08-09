package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (terminal *ExternalRepository) GetPackagingRates() (interface{}, error) {
	fmt.Println("Get Terminal Shipping rates...")
	req, err := buildNewTerminalRequest(GET, RATES, nil)
	var data interface{}
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	data, err = result(req, data)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	return data, nil
}

func (terminal *ExternalRepository) Rate(rateID string) error {
	fmt.Printf("Get Terminal Shipping rate with ID: %s", rateID)
	return nil
}

func (terminal *ExternalRepository) RatesForShipment() (interface{}, error) {
	fmt.Println("Get Terminal Shipping rate...")
	req, err := buildNewTerminalRequest(GET, fmt.Sprintf("%s/%s", RATES,SHIPMENT), nil)
	var data interface{}
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	data, err = result(req, data)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	return data, nil
}

func (terminal *ExternalRepository) CreatePackaging(packaging domain.PackagingDTO) (interface{}, error) {
	fmt.Println("Create a new package for shipping")
	var data interface{}
	requestBody := map[string]interface{}{
		"height": packaging.Height,
		"length": packaging.Length,
		"name": packaging.Name,
		"size_unit": packaging.Size_Unit,
		"type": packaging.Type,
		"width": packaging.Width,
		"weight": packaging.Weight,
		"weight_unit": packaging.Weight_Unit,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := buildNewTerminalRequest(POST, PACKAGING, bodyReader)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	data, err = result(req, data)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	return data, nil
}

func buildNewTerminalRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", os.Getenv("TERMINAL_AFRICA_TEST_ENDPOINT"), endpoint), body)
	if err != nil {
		log.Fatal(err)
		return req, err
	}
	req = setHeader(req)
	return req, nil
}

func setHeader(req *http.Request) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TERMINAL_AFRICA_SECRET_KEY")))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func request(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return response, err
	}
	return response, nil
}

func response(res *http.Response, data interface {}) (interface{}, error) {
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}

func result(req *http.Request, data interface {}) (interface{}, error) {
	res, err := request(req)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	data, err = response(res, data)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	return data, nil
}
