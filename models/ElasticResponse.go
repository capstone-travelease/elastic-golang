package models

import (
	"encoding/json"
)

type SearchResults struct {
	Took   int                        `json:"took"`
	Shards json.RawMessage            `json:"_shards"`
	Hits   Hits                       `json:"hits"`
	Aggs   map[string]json.RawMessage `json:"aggregations"`
}

type Hits struct {
	Total int
	Hits  []HitData `json:"hits"`
}

type HitData struct {
	Index  string           `json:"_index"`
	Id     string           `json:"_id"`
	Score  float64          `json:"_score"`
	Source LocationResponse `json:"_source"`
}

type DataResponse struct {
	Date             string `json:"date"`
	ShortDescription string `json:"short_description"`
	TimeStamp        string `json:"@timestamp"`
	Link             string `json:"link"`
	Category         string `json:"category"`
	HeadLine         string `json:"headline"`
	Authors          string `json:"authors"`
}

type LocationResponse struct {
	PlaceName     string `json:"placeName"`
	Code          int    `json:"code"`
	DivisionType  string `json:"divisionType"`
	CodePlaceName string `json:"codePlaceName"`
	PhoneCode     int    `json:"phoneCode"`
}
