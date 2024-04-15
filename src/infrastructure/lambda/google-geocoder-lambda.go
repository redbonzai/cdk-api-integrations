package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type GeoRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type GeoResponse struct {
	Address string `json:"address"`
}

func GeoCoderHandler(ctx context.Context, request GeoRequest) (*GeoResponse, error) {
	response, err := fetchGeoData(request.Latitude, request.Longitude)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func fetchGeoData(lat, lon string) (*GeoResponse, error) {
	key := os.Getenv("GOOGLE_API_KEY")
	geocoderBaseUrl := os.Getenv("GEOCODE_BASE_URL")
	url := fmt.Sprintf(geocoderBaseUrl+"?latlng=%s,%s&key=%s", lat, lon, key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching geocoding data: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding geocoding data: %v", err)
	}

	address := result["results"].([]interface{})[0].(map[string]interface{})["formatted_address"].(string)
	return &GeoResponse{Address: address}, nil
}

func main() {
	lambda.Start(GeoCoderHandler)
}
