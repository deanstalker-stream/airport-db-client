package airportdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	Namespace = "feed.airportdb"

	apiTokenKey       = "apiToken"
	endpointDelimiter = "/"
)

// BindEnvs registers environment variable mappings for this feed's config namespace.
func BindEnvs(v *viper.Viper) {
	_ = v.BindEnv("feed.airportdb.url", "FEED_AIRPORTDB_URL")
	_ = v.BindEnv("feed.airportdb.key", "FEED_AIRPORTDB_KEY")
}

// Client represents a client for the airportdb API.
type Client struct {
	logger *zap.Logger

	URL string
	Key string
}

// NewClient creates a new airportdb client.
func NewClient(logger *zap.Logger, cfg *Config) (*Client, error) {
	return &Client{
		logger: logger.Named(Namespace),
		URL:    cfg.URL,
		Key:    cfg.Key,
	}, nil
}

// GetAirport returns the airport with the given ICAO code.
func (c *Client) GetAirport(icao string) (*Airport, error) {
	path := icao
	req, err := c.getRequest(path)

	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	c.addAPIToken(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch airport data: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}

	var output *Airport
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, fmt.Errorf("unable to parse response JSON: %w", err)
	}

	return output, nil
}

// buildRequest constructs an HTTP request object for the given endpoint.
func (c *Client) getRequest(path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s%s", c.URL, endpointDelimiter, path)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return req, nil
}

// addAPIToken appends the API token to the request's query parameters.
func (c *Client) addAPIToken(req *http.Request) {
	query := req.URL.Query()
	query.Add(apiTokenKey, c.Key)
	req.URL.RawQuery = query.Encode()
}
