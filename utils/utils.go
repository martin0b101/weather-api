package utils

import (
	"encoding/xml"
	"io"
	"fmt"
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


func ParseCityFromXml(url string) (*types.City, error) {
    data, err := parseXMLToStruct(url)
    if err != nil {
        return nil, err
    }

    var city types.City
    var weather []types.Weather

    for _, item := range data.MetaData {
        if city.City == "" {
            layout := "02.01.2006 15:04 UTC"
            timeUpdated, err := time.Parse(layout, item.UpdatedAtUTC)
            if err != nil {
                return nil, err
            }
            city = types.City{
                City:     item.City,
                Country:  strings.ToLower(item.CountyCode),
                UpdatedAt: timeUpdated.Unix(),
            }
        }

        iconURL := data.IconURLBase
        windIcon := fmt.Sprintf("%s%s.%s", iconURL, item.IconWind, data.IconFormat)
        weatherIcon := fmt.Sprintf("%s%s.%s", iconURL, item.Icon, data.IconFormat)

        temp := types.WeatherTemp{
            Low:  item.LowTemp,
            Max:  item.MaxTemp,
            Unit: item.LowTempUnit,
        }

        wind := types.WeatherWind{
            Icon:           windIcon,
            Direction:      item.WindDirection,
            DirectionLong:  item.WindDirectionLong,
            Speed:          item.WindSpeed,
            Unit:           item.WindUnit,
        }

        day := strings.Split(item.Day, " ")[0]
        weather = append(weather, types.Weather{
            Day:   day,
            Valid: item.Valid,
            Icon:  weatherIcon,
            Temp:  temp,
            Wind:  wind,
        })
    }

    return &types.City{
        City:     city.City,
        Country:  city.Country,
        UpdatedAt: city.UpdatedAt,
        Weather:  weather,
    }, nil
}
