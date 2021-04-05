package main

import (
	"fmt"
	"net/url"
	"os"
	"time"

	framework "github.com/hyplabs/dfinity-oracle-framework"
	"github.com/hyplabs/dfinity-oracle-framework/models"
)

func generateEndpoints(cityName string, countryCode string) []models.Endpoint {
	weatherapiAPIKey := os.Getenv("WEATHERAPI_API_KEY")
	weatherbitAPIKey := os.Getenv("WEATHERBIT_API_KEY")
	openweathermapAPIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	return []models.Endpoint{
		{
			Endpoint: fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%v&q=%v,%v", url.QueryEscape(weatherapiAPIKey), url.QueryEscape(cityName), url.QueryEscape(countryCode)),
			JSONPaths: map[string]string{
				"temperature_celsius": "$.current.temp_c",
				"pressure_mbar":       "$.current.pressure_mb",
				"humidity_pct":        "$.current.humidity",
			},
		},
		{
			Endpoint: fmt.Sprintf("https://api.weatherbit.io/v2.0/current?key=%v&city=%v&country=%v", url.QueryEscape(weatherbitAPIKey), url.QueryEscape(cityName), url.QueryEscape(countryCode)),
			JSONPaths: map[string]string{
				"temperature_celsius": "$.data[0].temp",
				"pressure_mbar":       "$.data[0].pres",
				"humidity_pct":        "$.data[0].rh",
			},
		},
		{
			Endpoint: fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%v&q=%v,%v", url.QueryEscape(openweathermapAPIKey), url.QueryEscape(cityName), url.QueryEscape(countryCode)),
			JSONPaths: map[string]string{
				"temperature_kelvin": "$.main.temp",
				"pressure_mbar":      "$.main.pressure",
				"humidity_pct":       "$.main.humidity",
			},
			NormalizeFunc: func(val map[string]interface{}) (map[string]float64, error) {
				return map[string]float64{
					"temperature_celsius": val["temperature_kelvin"].(float64) - 273.15,
					"pressure_mbar":       val["pressure_mbar"].(float64),
					"humidity_pct":        val["humidity_pct"].(float64),
				}, nil
			},
		},
	}
}

func main() {
	config := models.Config{
		CanisterName: "weather_oracle",

		// this is based on the rate limits on all the APIs - the strictest rate limit is 500 calls/day from WeatherBit
		UpdateInterval: 1 * time.Hour,
	}

	// based on https://en.wikipedia.org/wiki/List_of_largest_cities#List
	engine := models.Engine{
		Metadata: []models.MappingMetadata{
			{Key: "Tokyo", Endpoints: generateEndpoints("Tokyo", "JP")},
			{Key: "Delhi", Endpoints: generateEndpoints("Delhi", "IN")},
			{Key: "Shanghai", Endpoints: generateEndpoints("Shanghai", "CN")},
			{Key: "São Paulo", Endpoints: generateEndpoints("São Paulo", "BR")},
			{Key: "Mexico City", Endpoints: generateEndpoints("Mexico City", "MX")},
			{Key: "Cairo", Endpoints: generateEndpoints("Cairo", "EG")},
			{Key: "Mumbai", Endpoints: generateEndpoints("Mumbai", "IN")},
			{Key: "Beijing", Endpoints: generateEndpoints("Beijing", "CN")},
			{Key: "Dhaka", Endpoints: generateEndpoints("Dhaka", "BD")},
			{Key: "Osaka", Endpoints: generateEndpoints("Osaka", "JP")},
			{Key: "New York City", Endpoints: generateEndpoints("New York City", "US")},
			{Key: "Karachi", Endpoints: generateEndpoints("Karachi", "PK")},
			{Key: "Buenos Aires", Endpoints: generateEndpoints("Buenos Aires", "AR")},
			{Key: "Chongqing", Endpoints: generateEndpoints("Chongqing", "CN")},
			{Key: "Istanbul", Endpoints: generateEndpoints("Istanbul", "TR")},
			{Key: "Kolkata", Endpoints: generateEndpoints("Kolkata", "IN")},
			{Key: "Manila", Endpoints: generateEndpoints("Manila", "PH")},
			{Key: "Lagos", Endpoints: generateEndpoints("Lagos", "NG")},
			{Key: "Rio de Janeiro", Endpoints: generateEndpoints("Rio de Janeiro", "BR")},
			{Key: "Tianjin", Endpoints: generateEndpoints("Tianjin", "CN")},
		},
	}

	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}
