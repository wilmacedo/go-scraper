package lib

import (
	"encoding/json"
	"go-scraper/src/config"
	"io"
	"net/http"
	"time"
)

type SearchQuery struct {
	Subject string
}

func FetchNews(searchQuery SearchQuery) (*ApiNewResponse, error) {
	env := config.GetEnv()

	request, err := http.NewRequest("GET", env.Endpoint+"/everything", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Api-Key", env.Key)

	query := request.URL.Query()
	query.Add("q", searchQuery.Subject)
	query.Add("searchIn", "title,description")

	request.URL.RawQuery = query.Encode()

	client := http.Client{
		Timeout: time.Duration(env.Timeout) * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data := &ApiNewResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
