package main

import (
	"log"
	"weather-demo/db"
	"weather-demo/es"
	"weather-demo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "weather-demo:weather-demo@tcp(127.0.0.1:3306)/weather?charset=utf8mb4&parseTime=True"
	dbClient, dbErr := db.Init(dsn)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	address := []string{"http://127.0.0.1:9200"}
	esClient, esErr := es.Init(address)
	if esErr != nil {
		log.Fatal(esErr)
	}

	h := handler.New(dbClient, esClient)
	syncErr := h.SyncCityWeathers()
	if syncErr != nil {
		log.Fatal(syncErr)
	}

	r := gin.Default()
	r.POST("/city-weathers/search", h.SearchCityWeathers)

	log.Fatal(r.Run(":8081"))
}
