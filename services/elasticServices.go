package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/billzayy/elastic-golang/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

func ConnectElastic() (*elasticsearch.Client, error) {
	godotenv.Load()
	cfg := elasticsearch.Config{
		CloudID: os.Getenv("CloudId"),
		APIKey:  os.Getenv("APIKey"),
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		resErr := fmt.Sprintf("Error creating the client: %s", err)
		return &elasticsearch.Client{}, errors.New(resErr)
	}

	return es, nil
}

func IndexDoc() {
	es, err := ConnectElastic()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch",
	}
	data, _ := json.Marshal(document)

	es.Index("my_index", bytes.NewReader(data))
}

func SearchDoc(inputName string) ([]models.LocationResponse, error) {
	es, err := ConnectElastic()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var r models.SearchResults
	var result []models.LocationResponse

	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"place_name": {
					"query": "%v"
				}
			}
		}
	}`, inputName)

	res, _ := es.Search(
		es.Search.WithIndex("vietnam_location"),
		es.Search.WithBody(strings.NewReader(query)),
	)

	json.NewDecoder(res.Body).Decode(&r)

	data := r.Hits.Hits

	for i := range data {
		result = append(result, data[i].Source)
	}
	return result, nil
}

func AllLocation() ([]models.ResultLocation, error) {
	es, err := ConnectElastic()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var r models.SearchAllResults
	var result []models.ResultLocation

	query := `{"size": 63,"track_total_hits": true}`

	res, _ := es.Search(
		es.Search.WithIndex("vietnam_location"),
		es.Search.WithBody(strings.NewReader(query)),
	)

	json.NewDecoder(res.Body).Decode(&r)

	data := r.Hits.Hits

	for i := range data {
		result = append(result, data[i].Source)
	}

	return result, nil
}
