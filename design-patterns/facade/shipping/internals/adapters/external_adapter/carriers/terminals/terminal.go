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

func (terminal *ExternalRepository) TerminalGetRates(query string) (map[string]interface{}, error) {
	req, err := buildNewTerminalRequest(GET.PrintMethod(), fmt.Sprintf("%s/%s?%s", RATES.PrintEndpoint(), SHIPMENT.PrintEndpoint(), query), nil)
	var rates map[string]interface{}
	if err != nil {
		log.Fatal(err)
		return rates, err
	}
	rates, err = result(req, rates)
	if err != nil {
		log.Fatal(err)
		return rates, err
	}
	return rates, nil
}

func (terminal *ExternalRepository) Rate(rateID string) error {
	fmt.Printf("Get Terminal Shipping rate with ID: %s", rateID)
	return nil
}

func (terminal *ExternalRepository) RatesForShipment() (map[string]interface{}, error) {
	req, err := buildNewTerminalRequest(GET.PrintMethod(), fmt.Sprintf("%s/%s", RATES.PrintEndpoint(), SHIPMENTS.PrintEndpoint()), nil)
	var shipments map[string]interface{}
	if err != nil {
		log.Fatal(err)
		return shipments, err
	}
	shipments, err = result(req, shipments)
	if err != nil {
		log.Fatal(err)
		return shipments, err
	}
	return shipments, nil
}

func (terminal *ExternalRepository) TerminalCreatePackaging(pack interface{}) (map[string]interface{}, error) {
	var packaging map[string]interface{}
	bodyReader, err := byteReader(pack)
	if err != nil {
		log.Fatal(err)
		return packaging, err
	}
	req, err := buildNewTerminalRequest(POST.PrintMethod(), PACKAGING.PrintEndpoint(), bodyReader)
	if err != nil {
		log.Fatal(err)
		return packaging, err
	}
	packaging, err = result(req, packaging)
	if err != nil {
		log.Fatal(err)
		return packaging, err
	}
	return packaging, nil
}

func (terminal *ExternalRepository) TerminalUpdatePackaging(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"packaging": shipment}, nil
}
func (terminal *ExternalRepository) TerminalDeletePackaging(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"packaging": shipment}, nil
}

func (terminal *ExternalRepository) TerminalGetPackaging(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"packaging": shipment}, nil
}

func (terminal *ExternalRepository) TerminalCreateAddress(add interface{}) (map[string]interface{}, error) {
	var address map[string]interface{}
	bodyReader, err := byteReader(add)
	if err != nil {
		log.Fatal(err)
		return address, err
	}
	req, err := buildNewTerminalRequest(POST.PrintMethod(), ADDRESSES.PrintEndpoint(), bodyReader)
	if err != nil {
		log.Fatal(err)
		return address, err
	}
	address, err = result(req, address)
	if err != nil {
		log.Fatal(err)
		return address, err
	}
	return address, nil
}

func (terminal *ExternalRepository) TerminalCreateParcel(parce interface{}) (map[string]interface{}, error) {
	var parcel map[string]interface{}
	bodyReader, err := byteReader(parce)
	if err != nil {
		log.Fatal(err)
		return parcel, err
	}
	req, err := buildNewTerminalRequest(POST.PrintMethod(), PARCELS.PrintEndpoint(), bodyReader)
	if err != nil {
		log.Fatal(err)
		return parcel, err
	}
	parcel, err = result(req, parcel)
	if err != nil {
		log.Fatal(err)
		return parcel, err
	}
	return parcel, nil
}

func (terminal *ExternalRepository) TerminalGetParcels() (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (terminal *ExternalRepository) TerminalGetParcel(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalDeleteParcel(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalUpdateParcel(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalCreateShipment(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalTrackShipment(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalCancelShipment(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalDeleteShipment(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
}

func (terminal *ExternalRepository) TerminalGetShipment(shipment interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"shipment": shipment}, nil
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
