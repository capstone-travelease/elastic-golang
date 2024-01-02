package services

import (
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

func SearchDoc(inputName string) ([]models.LocationResponse, error) {
	es, err := ConnectElastic()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	if inputName == "" {
		query := `{"size": 63,"track_total_hits": true}`

		return SearchLogic(es, query)
	} else {
		query := fmt.Sprintf(`{
			"query": {
				"multi_match" : {
					"query": "%v",
					"fields": [ "placeName", "codePlaceName" ],
					"fuzziness": "AUTO",
					"prefix_length": 3,
					"fuzzy_transpositions": "true"
					}
				}
		}`, inputName)

		return SearchLogic(es, query)
	}
}

func SearchLogic(es *elasticsearch.Client, query string) ([]models.LocationResponse, error) {
	var r models.SearchResults
	var result []models.LocationResponse

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
