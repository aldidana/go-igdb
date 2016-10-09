package igdb

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Companies represent companies
type Companies struct {
	ID                 *int    `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	CreatedAt          *int64  `json:"created_at,omitempty"`
	UpdatedAt          *int64  `json:"updated_at,omitempty"`
	Slug               *string `json:"slug,omitempty"`
	URL                *string `json:"url,omitempty"`
	StartDateCategory  *int    `json:"start_date_category,omitempty"`
	ChangeDateCategory *int    `json:"change_date_category,omitempty"`
}

//GetCompanies list of companies
func (c Client) GetCompanies(q string, limit int, offset int) ([]Companies, error) {
	query := c.defaultBaseURL + "companies/?fields=" + q + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)

	req, err := http.NewRequest("GET", query, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Set(c.headerMashape, c.apiKey)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer response.Body.Close()

	companies := []Companies{}

	if err := json.NewDecoder(response.Body).Decode(&companies); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return companies, nil
}
