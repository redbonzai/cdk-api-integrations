package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"net/http"
	"os"
	_ "os"
)

type WeatherRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type WeatherResponse struct {
	Condition   string `json:"condition"`
	Description string `json:"description"`
	Temperature string `json:"temperature"`
}

type WeatherAPIResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func WeatherMapHandler(ctx context.Context, request WeatherRequest) (*WeatherResponse, error) {
	response, err := fetchWeatherData(request.Latitude, request.Longitude)
	if err != nil {
		return nil, err
	}

	saveWeatherDataToDynamoDB(response)
	return response, nil
}

func fetchWeatherData(lat, lon string) (*WeatherResponse, error) {
	weatherMapKey := os.Getenv("WEATHER_API_KEY")
	weatherMapBaseUrl := os.Getenv("WEATHER_MAP_BASE_URL")
	url := fmt.Sprintf(weatherMapBaseUrl+"/weather?lat=%s&lon=%s&appid=%s", lat, lon, weatherMapKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding weather data: %v", err)
	}

	temperature := "moderate"
	tempCelsius := apiResponse.Main.Temp - 273.15
	if tempCelsius < 10 {
		temperature = "cold"
	} else if tempCelsius > 25 {
		temperature = "hot"
	}

	response := &WeatherResponse{
		Condition:   apiResponse.Weather[0].Main,
		Description: apiResponse.Weather[0].Description,
		Temperature: temperature,
	}

	return response, nil
}

func saveWeatherDataToDynamoDB(data *WeatherResponse) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)
	item, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		fmt.Printf("Error marshalling data: %v\n", err)
		return
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("WeatherData"),
		Item:      item,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Printf("Error saving data to DynamoDB: %v\n", err)
	}
}

func main() {
	lambda.Start(WeatherMapHandler)
}
