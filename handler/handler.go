package handler

import (
	"weather-demo/db"
	"weather-demo/es"
)

type Handler struct {
	DBClient db.Client
	ESClient es.Client
}

func New(dbClient db.Client, esClient es.Client) *Handler {
	return &Handler{
		DBClient: dbClient,
		ESClient: esClient,
	}
}
