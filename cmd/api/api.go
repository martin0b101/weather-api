package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weather-api/service/weather"
)


type ApiServer struct {
	address string
}


func NewApiServer(address string) (*ApiServer) {
	return &ApiServer{
		address: address,
	}
}

func (s *ApiServer) Run() error {
	router := gin.Default()

	weatherStore := weather.NewWeatherStore()

	weatherHandler := weather.NewHandler(weatherStore)
	weatherHandler.RegisterRoutes(router)

	return router.Run(s.address)
}