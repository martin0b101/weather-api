package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weather-api/types"
)


type Handler struct {
	store types.WeatherStore
}

func NewHandler(store types.WeatherStore) *Handler{
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.Engine){
	router.GET("/weather", h.handleGetWeather)
	router.GET("/weather/:city", h.handleGetWeatherByCity)
}


func (h *Handler) handleGetWeather(c *gin.Context) {
	cities, err := h.store.GetAllWeather()

	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, types.Response{
			Status: http.StatusInternalServerError,
			Error: true,
			Data: make([]types.City, 0),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, types.Response{
		Status: http.StatusOK,
		Error: false,
		Data: cities,
	})
}

func (h *Handler) handleGetWeatherByCity(c *gin.Context){
	cityFilter := c.Param("city")

	if cityFilter == ""{
		c.IndentedJSON(http.StatusInternalServerError, types.Response{
			Status: http.StatusInternalServerError,
			Error: true,
			Data: struct{}{},
		})
		return
	}

	city, err := h.store.GetWeatherByCity(cityFilter)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.Response{
			Status: http.StatusInternalServerError,
			Error: true,
			Data: struct{}{},
		})
		return
	}

	if city.City == "" {
		c.IndentedJSON(http.StatusNotFound, types.Response{
			Status: http.StatusNotFound,
			Error: false,
			Data: struct{}{},
		})
		return
	}
	c.IndentedJSON(http.StatusOK, types.Response{
		Status: http.StatusOK,
		Error: false,
		Data: city,
	})
	
}