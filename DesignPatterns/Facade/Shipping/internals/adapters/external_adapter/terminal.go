package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/QUDUSKUNLE/shipping/internals/core/ports"
)

func (terminal *ExternalRepository) TerminalRequest(endpoint ports.ENDPOINTS, method ports.METHOD, requestBody any) (map[string]interface{}, error) {
	var requestResult map[string]interface{}
	var bodyReader *bytes.Reader
	var err error
	var req *http.Request
	switch endpoint {
	case ports.SHIPMENT:
		switch method {
			case ports.GET:
				req, err = buildNewTerminalRequest(ports.GET.PrintMethod(), fmt.Sprintf("%s/%s?%s", ports.RATES.PrintEndpoint(), ports.SHIPMENT.PrintEndpoint(), requestBody), nil)
			}
	case ports.SHIPMENTS:
		switch method {
			case ports.GET:
				req, err = buildNewTerminalRequest(ports.GET.PrintMethod(), fmt.Sprintf("%s/%s", ports.RATES.PrintEndpoint(), ports.SHIPMENTS.PrintEndpoint()), nil)
			case ports.POST:
				bodyReader, err = byteReader(requestBody)
				if err != nil {
					log.Fatal(err)
					return requestResult, err
				}
				fmt.Println(requestBody, "**********")
				req, err = buildNewTerminalRequest(ports.POST.PrintMethod(), ports.SHIPMENTS.PrintEndpoint(), bodyReader)
		}
	case ports.PACKAGING:
		switch method {
			case ports.POST:
				bodyReader, err = byteReader(requestBody)
				if err != nil {
					log.Fatal(err)
					return requestResult, err
				}
				req, err = buildNewTerminalRequest(ports.POST.PrintMethod(), ports.PACKAGING.PrintEndpoint(), bodyReader)
		}
	case ports.ADDRESSES:
		switch method {
			case ports.POST:
				bodyReader, err = byteReader(requestBody)
				if err != nil {
					log.Fatal(err)
					return requestResult, err
				}
				req, err = buildNewTerminalRequest(ports.POST.PrintMethod(), ports.ADDRESSES.PrintEndpoint(), bodyReader)
		}
	case ports.PARCELS:
		switch method {
			case ports.POST:
				bodyReader, err = byteReader(requestBody)
				if err != nil {
					log.Fatal(err)
					return requestResult, err
				}
				req, err = buildNewTerminalRequest(ports.POST.PrintMethod(), ports.PARCELS.PrintEndpoint(), bodyReader)
		}
	}
	if err != nil {
		log.Fatal(err)
		return requestResult, err
	}
	requestResult, err = result(req, requestResult)
	if err != nil {
		log.Fatal(err)
		return requestResult, err
	}
	return requestResult, nil
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

func response(res *http.Response, data map[string]interface {}) (map[string]interface{}, error) {
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}

func result(req *http.Request, data map[string]interface {}) (map[string]interface{}, error) {
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
