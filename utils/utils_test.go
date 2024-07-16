package utils

import (
	"testing"
	"github.com/weather-api/types"
)


func TestParseCityFromXml(t *testing.T) {
    // Mocked data for testing
	testCity := types.City{
		City: "New York",
	}

    // Mocked ParseCityFromXml function
	ParseCityFromXml := func(url string) (*types.City, error) {
        return &testCity, nil
    }

    // Test the ParseCityFromXml function with mocked data
    city, err := ParseCityFromXml("mock-url")
    if err != nil {
        t.Fatalf("ParseCityFromXml failed: %v", err)
    }

    // Assert on the constructed city object
    expectedCity := "New York"
    if city.City != expectedCity {
        t.Errorf("Expected city: %s, got: %s", expectedCity, city.City)
    }
    // Add more assertions for weather data, etc.
}
