package model

type CityWeather struct {
	Id int64 `json:"id"`
	CityCode string `json:"city_code"`
	Country string `json:"country"`
	CountryEn string `json:"country_en"`
	Province string `json:"province"`
	City string `json:"city"`
	CityEn string `json:"city_en"`
	County string `json:"county"`
	Temp string `json:"temp"`
	TempFahrenheit string `json:"temp_fahrenheit"`
	Wind string `json:"wind"`
	WindScale string `json:"wind_scale"`
	Humidity string `json:"humidity"`
	Weather string `json:"weather"`
	WeatherEn string `json:"weather_en"`
	WeatherCode string `json:"weather_code"`
	Aqi string `json:"aqi"`
	AqiPm25 string `json:"aqi_pm25"`
	Rain string `json:"rain"`
	Rain24h string `json:"rain_24h"`
	MaxTemp string `json:"max_temp"`
	MinTemp string `json:"min_temp"`
	DayAndNight int `json:"day_and_night"`
	Time string `json:"time"`
	Date string `json:"date"`
	UpsertTime int64 `json:"upsert_time"`
}
