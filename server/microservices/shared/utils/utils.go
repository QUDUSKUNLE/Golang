package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/types/known/structpb"
)

type Config struct {
	Port        string
	KafkaBroker string
	KafkaTopic  string
	KafkaGroup  string
	DB_URL      string
}

func LoadEnvironmentVariable() error {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
		return nil // Return nil to allow fallback to system environment variables
	}
	return nil
}

func LoadConfig() (*Config, error) {
	if err := LoadEnvironmentVariable(); err != nil {
		return nil, err
	}

	config := &Config{
		Port:        os.Getenv("PORT"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
		KafkaGroup:  os.Getenv("KAFKA_GROUP_ID"),
		DB_URL:      os.Getenv("DB_URL"),
	}
	// Validate required fields
	if config.Port == "" || config.KafkaBroker == "" || config.DB_URL == "" {
		return nil, fmt.Errorf("missing required environment variables: PORT, KAFKA_BROKER, or DB_URL")
	}
	return config, nil
}

func Haversine(lat1, lon1, lat2, lon2 float64) (float64, error) {
	// Validate latitude and longitude ranges
	if lat1 < -90 || lat1 > 90 || lat2 < -90 || lat2 > 90 {
		return 0, fmt.Errorf("latitude must be between -90 and 90")
	}
	if lon1 < -180 || lon1 > 180 || lon2 < -180 || lon2 > 180 {
		return 0, fmt.Errorf("longitude must be between -180 and 180")
	}

	const earthRadius = 6371 // Earth's radius in kilometers
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c, nil
}

func JsonMarshal(data interface{}) ([]byte, error) {
	// Marshal the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	return jsonData, nil
}

func MapToStruct(data map[string]interface{}) (*structpb.Struct, error) {
	if data == nil {
		return nil, fmt.Errorf("input map cannot be nil")
	}

	structData, err := structpb.NewStruct(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert map to struct: %w", err)
	}
	return structData, nil
}

func PaginationParams(limit, offset int32) (int32, int32) {
	if limit <= 0 {
		fmt.Printf("Invalid limit: %d. Using default limit: %d\n", limit, constants.DefaultLimit)
		limit = constants.DefaultLimit
	}
	if offset < 0 {
		fmt.Printf("Invalid offset: %d. Using default offset: %d\n", offset, constants.DefaultOffset)
		offset = constants.DefaultOffset
	}
	return limit, offset
}
