package models

// WeatherData is struct to read the weather data
type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
		TempKf    float64 `json:"temp_kf"`
	} `json:"main"`
	Visibility float64 `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Dt  float64 `json:"dt"`
	Sys struct {
		Type    float64 `json:"type"`
		ID      float64 `json:"id"`
		Country string  `json:"country"`
		Sunrise float64 `json:"sunrise"`
		Sunset  float64 `json:"sunset"`
		Area    string  `json:"area"`
		City    string  `json:"city"`
	} `json:"sys"`
	CountryName string  `json:"country_name"`
	Timezone    float64 `json:"timezone"`
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	Cod         float64 `json:"cod"`
}
