package api

import (
	"fmt"

	"github.com/tomnomnom/linkheader"

	resty "gopkg.in/resty.v0"
)

// InfraClient represents the infra client state for the API.
type InfraClient struct {
	RestyClient *resty.Client
}

// NewInfra returns a new Infra Client for the specified apiKey.
func NewInfra(config Config) InfraClient {
	r := resty.New()

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://infra-api.newrelic.com/v2"
	}

	r.SetHeader("X-Api-Key", config.APIKey)
	r.SetHostURL(baseURL)

	if config.TLSConfig != nil {
		r.SetTLSClientConfig(config.TLSConfig)
	}
	if config.Debug {
		r.SetDebug(true)
	}

	c := InfraClient{
		RestyClient: r,
	}

	return c
}

// Do exectes an API request with the specified parameters.
func (c *InfraClient) Do(method string, path string, body interface{}, response interface{}) (string, error) {
	r := c.RestyClient.R().
		SetError(&ErrorResponse{})

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	apiResponse, err := r.Execute(method, path)

	if err != nil {
		return "", err
	}

	nextPath := ""
	header := apiResponse.Header().Get("Link")
	if header != "" {
		links := linkheader.Parse(header)

		for _, link := range links.FilterByRel("next") {
			nextPath = link.URL
			break
		}
	}

	statusClass := apiResponse.StatusCode() / 100 % 10

	if statusClass == 2 {
		return nextPath, nil
	}

	rawError := apiResponse.Error()

	if rawError != nil {
		apiError := rawError.(*ErrorResponse)

		if apiError.Detail != nil {
			return "", apiError
		}
	}

	return "", fmt.Errorf("Unexpected status %v returned from API", apiResponse.StatusCode())
}
