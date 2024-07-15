package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"time"
	"github.com/weather-api/types"
)



func parseXMLToStruct(xmlURL string) (*types.Data, error) {
    // Fetch XML data from URL
    resp, err := http.Get(xmlURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the body of the response into a []byte
    xmlData, err := io.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    // Parse XML data
    var data types.Data
    err = xml.Unmarshal(xmlData, &data)
    if err != nil {
        return nil, err
    }

    return &data, nil
}


func ParseCityFromXml(url string) (*types.City, error){
	data, errorParse := parseXMLToStruct(url)
	
	if(errorParse != nil){
		return nil, errorParse
	}

	weatherArray := []types.Weather{}
	cityItem := types.City{}

	for _, item := range data.MetaData {

		var iconUrl = data.IconURLBase

		temp := types.WeatherTemp{
			Low: item.LowTemp,
			Max: item.MaxTemp,
			Unit: item.LowTempUnit,
		}

		wind := types.WeatherWind{
			Icon: iconUrl+item.IconWind+"."+data.IconFormat,
			Direction: item.WindDirection,
			DirectionLong:  item.WindDirectionLong,
			Speed: item.WindSpeed,
			Unit: item.WindUnit,
		}


		var dayParsed = strings.Split(item.Day, " ")


		weatherItem := types.Weather{ 
			Day: dayParsed[0], 
			Valid: item.Valid, 
			Icon: iconUrl +item.Icon+"."+ data.IconFormat,
			Temp: temp,
			Wind: wind,
		 }

		weatherArray = append(weatherArray, weatherItem)

		if cityItem.City == "" {

			layout := "02.01.2006 15:04 UTC"
			timeUpdated, err := time.Parse(layout, item.UpdatedAtUTC)
			
			if err != nil {
				return nil, err
			}

			timeUnix := timeUpdated.Unix()

			cityItem = types.City{
				City: item.City,
				Country: strings.ToLower(item.CountyCode),
				UpdatedAt: timeUnix,
			}
		}

	}

	cityItemFinal := types.City{
		City: cityItem.City,
		Country: cityItem.Country,
		UpdatedAt: cityItem.UpdatedAt,
		Weather: weatherArray,
	}

	return &cityItemFinal, nil
}