package es

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Client struct {
	ES *elasticsearch.Client
}

func Init(Address []string) (Client, error) {
	client := Client{}
	es, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: Address})
	if err != nil {
		return client, err
	}

	_, err = es.Info()
	if err != nil {
		return client, err
	}

	client.ES = es
	return client, nil
}

func (c Client) CreateDocument(index string, documentId string, body io.Reader) (*esapi.Response, error) {
	req := esapi.IndexRequest{
		Index: index,
		DocumentID: documentId,
		Body: body,
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), c.ES)
	if err != nil {
		return res, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	return res, nil
}

func (c Client) CreateIndex(index string, body io.Reader) (*esapi.Response, error) {
	req := esapi.IndexRequest{
		Index: index,
		Body: body,
	}

	res, err := req.Do(context.Background(), c.ES)
	if err != nil {
		return res, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	return res, nil
}
