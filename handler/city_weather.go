package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"weather-demo/model"

	"github.com/gin-gonic/gin"
)

func (h Handler) SearchCityWeathers(c *gin.Context) {
	city := c.PostForm("city")
	body := fmt.Sprintf(`{"query": {"multi_match": {"query": "%s", "fields": ["city", "city_en"]}}}`, city)

	res, err := h.ESClient.ES.Search(
		h.ESClient.ES.Search.WithContext(context.Background()),
		h.ESClient.ES.Search.WithIndex("city_weather"),
		h.ESClient.ES.Search.WithBody(strings.NewReader(body)),
		h.ESClient.ES.Search.WithPretty(),
		)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer res.Body.Close()

	if res.IsError() {
		c.JSON(res.StatusCode, gin.H{"error": res.Status()})
		return
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var weathers []interface{}
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		weather := hit.(map[string]interface{})["_source"]
		weathers = append(weathers, weather)
	}

	c.JSON(http.StatusOK, gin.H{"data": weathers})
}

// SyncCityWeathers sync data from db to es
func (h Handler) SyncCityWeathers() (err error) {
	var weathers []model.CityWeather
	weathers, err = h.DBClient.GetCityWeathers()
	if err != nil {
		return
	}

	for _, weather := range weathers {
		var body []byte
		body, err = json.Marshal(weather)
		if err != nil {
			return
		}

		_, err = h.ESClient.CreateDocument(
			"city_weather",
			strconv.FormatInt(weather.Id, 10),
			strings.NewReader(string(body)),
		)
		if err != nil {
			return
		}
	}

	return nil
}
