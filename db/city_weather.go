package db

import (
	"database/sql"
	"weather-demo/model"
)

func (c Client) GetCityWeathers() (cityWeathers []model.CityWeather, err error) {
	var rows *sql.Rows
	rows, err = c.DB.Query("select * from city_realtime_weather_com")
	if err != nil {
		return
	}

	for rows.Next() {
		var cityWeather model.CityWeather
		err = rows.Scan(
			&cityWeather.Id,
			&cityWeather.CityCode,
			&cityWeather.Country,
			&cityWeather.CountryEn,
			&cityWeather.Province,
			&cityWeather.City,
			&cityWeather.CityEn,
			&cityWeather.County,
			&cityWeather.Temp,
			&cityWeather.TempFahrenheit,
			&cityWeather.Wind,
			&cityWeather.WindScale,
			&cityWeather.Humidity,
			&cityWeather.Weather,
			&cityWeather.WeatherEn,
			&cityWeather.WeatherCode,
			&cityWeather.Aqi,
			&cityWeather.AqiPm25,
			&cityWeather.Rain,
			&cityWeather.Rain24h,
			&cityWeather.MaxTemp,
			&cityWeather.MinTemp,
			&cityWeather.DayAndNight,
			&cityWeather.Time,
			&cityWeather.Date,
			&cityWeather.UpsertTime,
		)

		if err != nil {
			return
		}

		cityWeathers = append(cityWeathers, cityWeather)
	}

	return
}
