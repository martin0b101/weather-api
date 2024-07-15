package helper

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"github.com/weather-api/types"
)

func LoadFile() (*types.JsonDB, error){
	data, err := openAndLoad()
	if err != nil {
		log.Fatal(err)
	}
	return data, err;
}

func openAndLoad() (*types.JsonDB, error){
	jsonFile, err := os.Open("arso-xmls.json")
	if err != nil {
		return nil, err;
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil{
		return nil, err
	}

	var xmls types.JsonDB

	json.Unmarshal(byteValue, &xmls)

	return &xmls, nil
}