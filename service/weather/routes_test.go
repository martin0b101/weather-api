package weather

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weather-api/types"
)


func TestGetAllWeatherHandler(t *testing.T){

	weatherStore := &mockWeatherStore{}
	handler := NewHandler(weatherStore)

	t.Run("Should retrun all weathers", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/weather", nil)

		if err != nil{
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/weather", handler.handleGetWeather)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK{
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}


func TestGetWeatherByCityHandler(t *testing.T){
	
	weatherStore := &mockWeatherStore{}
	handler := NewHandler(weatherStore)

	t.Run("Should retrun Zadar weather", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/weather/Zadar", nil)

		if err != nil{
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()


		router := gin.Default()
		router.GET("/weather/:city", handler.handleGetWeatherByCity)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK{
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}


func TestGetWeatherByCityHandlerNotFound(t *testing.T){

	weatherStore := &mockWeatherStore{}
	handler := NewHandler(weatherStore)
	t.Run("Should retrun not found response", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/weather/Zadar1", nil)

		if err != nil{
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/weather/:city", handler.handleGetWeatherByCity)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound{
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
		
	})
}

func TestGetWeatherByCityHandlerInternalError(t *testing.T){
	weatherStore := &mockWeatherStore{}
	handler := NewHandler(weatherStore)
	t.Run("Should retrun not found response", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/weather/1", nil)

		if err != nil{
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/weather/:city", handler.handleGetWeatherByCity)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError{
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
		}
		
	})
}


type mockWeatherStore struct {}

func (m *mockWeatherStore) GetAllWeather() ([]types.City, error){

	return nil, nil
}

func (m *mockWeatherStore) GetWeatherByCity(filterCity string) (*types.City, error){

	if filterCity == "Zadar"{
		return &types.City{City: "Zadar"}, nil
	}

	if filterCity == "Zadar1"{
		return &types.City{City: ""}, nil
	}

	if filterCity == "1"{
		return nil, fmt.Errorf("City has to be a string!")
	}

	return nil, nil
}


