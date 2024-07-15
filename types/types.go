package types


type CityURLMapping struct {
	City string `json:"city"`
	Url string `json:"xml-url"`
}

type JsonDB struct {
	CityURLs []CityURLMapping `json:"xmls"`
}

type City struct {
	City string `json:"city"`
	Country string `json:"country"`
	UpdatedAt int64 `json:"updatedAt"`
	Weather []Weather `json:"weather"`
}

type Weather struct {
	Day string `json:"day"`
	Valid string `json:"valid"`
	Icon string `json:"icon"`
	Temp WeatherTemp `json:"temp"`
	Wind WeatherWind `json:"wind"`
}

type WeatherTemp struct {
	Low int `json:"low"`
	Max int `json:"max"`
	Unit string `json:"unit"`
}

type WeatherWind struct {
	Icon string `json:"icon"`
	Direction string `json:"direction"`
	DirectionLong string `json:"directionLong"`
	Speed float32 `json:"speed"`
	Unit string `json:"unit"`
}

type Response struct {
	Status int `json:"status"`
	Error bool `json:"error"`
	Data interface{} `json:"data"`
}

type Data struct {
	ID                 string   `xml:"id,attr"`
    KatalonExtensionID string   `xml:"katalonExtensionId,attr"`
    Language           string   `xml:"language"`
    Credit             string   `xml:"credit"`
    CreditURL          string   `xml:"credit_url"`
    ImageUrl           string   `xml:"image_url"`
    SuggestedPickup    string   `xml:"suggested_pickup"`
    SuggestedPickupPeriod int    `xml:"suggested_pickup_period"`
    WebcamURLBase      string   `xml:"webcam_url_base"`
    IconURLBase        string   `xml:"icon_url_base"`
    IconFormat         string   `xml:"icon_format"`
    DocsURL            string   `xml:"docs_url"`
    DisclaimerURL      string   `xml:"disclaimer_url"`
    CopyrightURL       string   `xml:"copyright_url"`
    PrivacyPolicyURL   string   `xml:"privacy_policy_url"`
    ManagingEditor     string   `xml:"managing_editor"`
    WebMaster          string   `xml:"web_master"`
    Generator          string   `xml:"generator"`
    MeteosIURL         string   `xml:"meteosi_url"`
	MetaData []MetaData `xml:"metData"`
}

type MetaData struct {

	City string `xml:"domain_longTitle"`
	CountyCode string `xml:"domain_countryIsoCode2"`
	UpdatedAtUTC string `xml:"tsUpdated_UTC"`

	//weather
	Day string `xml:"valid_day"`
	Valid string `xml:"valid"`
	Icon string `xml:"nn_icon"`
	//temp
	LowTemp int `xml:"tnsyn"`
	LowTempUnit string `xml:"tnsyn_var_unit"`
	LowTempInDegreesC string `xml:"tnsyn_degreesC"`
	MaxTemp int `xml:"txsyn"`
	MaxTempUnit string `xml:"txsyn_var_unit"`
	MaxTempInDegreesC string `xml:"txsyn_degreesC"`

	//wind
	IconWind string `xml:"dd_icon"`
	WindDirection string `xml:"dd_shortText"`
	WindDirectionLong string `xml:"dd_longText"`
	WindUnit string `xml:"ff_var_unit"`
	WindSpeed float32 `xml:"ff_val"`

}



type WeatherStore interface{
	GetAllWeather() ([]City, error)
	GetWeatherByCity(filterCity string) (*City, error)
}