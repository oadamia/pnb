package newsapi

import "github.com/go-resty/resty/v2"

const (
	sourcePath = "/top-headlines/sources"
)

type Client struct {
	APIKey  string
	rClient *resty.Client
}

func NewNewsAPI(baseurl string, apiKey string) *Client {
	r := resty.New().
		SetBaseURL(baseurl).
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json")

	return &Client{
		APIKey:  apiKey,
		rClient: r,
	}

}

func (n *Client) GetSources() (*SourcesResult, error) {
	var result SourcesResult

	_, err := n.rClient.R().
		SetQueryParam("language", "en").
		SetQueryParam("country", "us").
		SetResult(&result).
		Get(sourcePath)

	if err != nil {
		return nil, err
	}

	return &result, nil

}
