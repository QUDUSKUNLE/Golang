package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (terminal *ExternalRepository) TerminalGetPackagingRates() (interface{}, error) {
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

func (terminal *ExternalRepository) TerminalCreatePackaging(packaging interface{}) (interface{}, error) {
	fmt.Println("Create a new package for shipping")
	var data interface{}
	bodyReader, err := byteReader(packaging)
	if err != nil {
		log.Fatal(err)
	}
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

func (terminal *ExternalRepository) TerminalCreateAddress(address interface{}) (interface{}, error) {
	fmt.Println("Create a new address for shipping")
	var data interface{}
	bodyReader, err := byteReader(address)
	if err != nil {
		log.Fatal(err)
	}
	req, err := buildNewTerminalRequest(POST, ADDRESSES, bodyReader)
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

func byteReader(reader interface{}) (*bytes.Reader, error) {
	var result *bytes.Reader
	jsonBody, err := json.Marshal(reader)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return bytes.NewReader(jsonBody), nil
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
