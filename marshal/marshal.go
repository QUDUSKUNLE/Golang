package marshal

import (
	"encoding/json"
)

// Decoding JSON to structured data
type Bird struct {
	Species string
	Description string
}

// Decoding JSON using Marshal
type BirdMarshal struct {
	Species string `json:"birdType"`
	Description string `json:"what it does"`
}

type Dimension struct {
	Height int
	Width int
}

type BirdsDimension struct {
	Bird
	Dimension Dimension
}

func (t *Bird) ReturnBird(bird string) (Bird, error) {
	var birde Bird
	if err := json.Unmarshal([]byte(bird), &birde); err != nil {
		return Bird{}, err
	}
	return birde, nil
}

func (t *Bird) ReturnBirdsArray(bird string) ([]Bird, error) {
	var birde []Bird
	if err := json.Unmarshal([]byte(bird), &birde); err != nil {
		return []Bird{}, err
	}
	return birde, nil
}

func (t *BirdsDimension) ReturnBirdWithDimesion(bird string) (BirdsDimension, error) {
	var birdDim BirdsDimension
	if err := json.Unmarshal([]byte(bird), &birdDim); err != nil {
		return BirdsDimension{}, err
	}
	return birdDim, nil
}

// Decoding JSON to maps - Unstructured data
func UnknownResult(bird string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(bird), &result); err != nil {
		return map[string]interface{}{}, err
	}
	return result, nil
}

// Validate Json Data
func (t *Bird) ValidateJsonData(bird string) (bool) {
	return json.Valid([]byte(bird))
}

// Marshaling Structured Data
func (t *BirdMarshal) BirdMarshalResponse(birdMarshal BirdMarshal) ([]byte, error) {
	response, err := json.Marshal(&birdMarshal);
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}

// Marshaling Slice
func (t *BirdMarshal) BirdMarshalArray(birdMarshal BirdMarshal) ([]byte, error) {
	response, err := json.Marshal([]*BirdMarshal{&birdMarshal})
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}

// Marshaling Maps
func (t *BirdMarshal) MarshalingMaps(data map[string]interface{}) ([]byte, error) {
	response, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}
