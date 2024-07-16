package weather

import (
	"github.com/weather-api/helper"
	"github.com/weather-api/types"
	"github.com/weather-api/utils"
)


type Store struct {
}


func NewWeatherStore() *Store {
	return &Store{}
}

func (s *Store) GetAllWeather() ([]types.City, error){

	cacheMutex.Lock()
	if cache != nil {
		cacheMutex.Unlock()
		return *cache, nil
	}

	cacheMutex.Unlock()

	data, err := helper.LoadFile()

	if err != nil{
		return nil, err
	}

	var cities []types.City

	for _, item := range data.CityURLs{
		cityParsed, err := utils.ParseCityFromXml(item.Url)

		if err != nil{
			return []types.City{}, err
		}

		cities = append(cities, *cityParsed)
	}

	cacheMutex.Lock()
	cache = &cities
	cacheMutex.Unlock()

	return cities, nil
}

func (s *Store) GetWeatherByCity(filterCity string) (*types.City, error){

	cacheMutex.Lock()
	if cache != nil {
		for _, item := range *cache{
			if(item.City == filterCity){
				cacheMutex.Unlock()
				return &item, nil
			}
		}
	}
	cacheMutex.Unlock()

	data, err := helper.LoadFile()

	if err != nil{
		return nil, err
	}

	for _, item := range data.CityURLs{
		if item.City != filterCity{
			continue
		}
		var cityReturned, errorFetching = utils.ParseCityFromXml(item.Url)

		if(errorFetching != nil){
			return nil, errorFetching
		}

		return cityReturned, nil
	}
	return &types.City{}, nil;
}
