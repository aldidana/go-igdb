package igdb

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Genres represent genres
type Genres struct {
	ID        *int    `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
	Slug      *string `json:"slug,omitempty"`
	Games     *[]int  `json:"games,omitempty"`
}

//GetGenres get list of genres
func (c Client) GetGenres(q string, limit int) ([]Genres, error) {
	var queryQ string
	var queryLimit int

	if q == "" {
		queryQ = "*"
	} else {
		queryQ = q
	}

	if limit == 0 {
		queryLimit = 40
	} else {
		queryLimit = limit
	}

	query := c.defaultBaseURL + "genres/?fields=" + queryQ + "&limit=" + strconv.Itoa(queryLimit)

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

	genres := []Genres{}

	if err := json.NewDecoder(response.Body).Decode(&genres); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return genres, nil
}
