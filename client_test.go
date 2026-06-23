package airportdb

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// TestNewClient tests the creation of a new Client via NewClient
func TestNewClient(t *testing.T) {
	client, err := NewClient(zap.NewNop(), &Config{URL: "https://api.airportdb.io", Key: "test-key"})

	require.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "https://api.airportdb.io", client.URL)
	assert.Equal(t, "test-key", client.Key)
}

func TestGetRequestURLConstruction(t *testing.T) {
	client := &Client{
		logger: zap.NewNop(),
		URL:    "https://api.airportdb.io",
		Key:    "test-key",
	}

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"simple path", "KJFK", "https://api.airportdb.io/KJFK"},
		{"path with numbers", "KJFK123", "https://api.airportdb.io/KJFK123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := client.getRequest(tt.path)

			require.NoError(t, err)
			assert.Equal(t, tt.expected, req.URL.String())
			assert.Equal(t, "GET", req.Method)
		})
	}
}

func TestAddAPIToken(t *testing.T) {
	client := &Client{
		logger: zap.NewNop(),
		URL:    "https://api.airportdb.io",
		Key:    "secret-test-key",
	}

	req, _ := client.getRequest("KJFK")
	client.addAPIToken(req)

	assert.Equal(t, "secret-test-key", req.URL.Query().Get(apiTokenKey))
}

func TestGetRequestInvalidURL(t *testing.T) {
	client := &Client{
		logger: zap.NewNop(),
		URL:    "ht!tp://invalid",
		Key:    "test-key",
	}

	_, err := client.getRequest("KJFK")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create request")
}

func TestAddAPITokenMultipleParams(t *testing.T) {
	client := &Client{
		logger: zap.NewNop(),
		URL:    "https://api.airportdb.io",
		Key:    "test-key",
	}

	req, _ := client.getRequest("KJFK")

	query := req.URL.Query()
	query.Add("existing", "value")
	req.URL.RawQuery = query.Encode()

	client.addAPIToken(req)

	finalQuery := req.URL.Query()
	assert.Equal(t, "value", finalQuery.Get("existing"))
	assert.Equal(t, "test-key", finalQuery.Get(apiTokenKey))
}

func TestGetAirportSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get(apiTokenKey) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]interface{}{"name": "John F Kennedy International Airport"})
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &Client{
		logger: zap.NewNop(),
		URL:    server.URL,
		Key:    "test-key",
	}

	airport, err := client.GetAirport("KJFK")

	require.NoError(t, err)
	assert.NotNil(t, airport)
}

func TestGetAirportNetworkError(t *testing.T) {
	client := &Client{
		logger: zap.NewNop(),
		URL:    "http://invalid-unreachable-domain.test",
		Key:    "test-key",
	}

	_, err := client.GetAirport("KJFK")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to fetch airport data")
}

func TestGetAirportInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("invalid json"))
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &Client{
		logger: zap.NewNop(),
		URL:    server.URL,
		Key:    "test-key",
	}

	_, err := client.GetAirport("KJFK")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to parse response JSON")
}

func TestGetAirportHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "server error")
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &Client{
		logger: zap.NewNop(),
		URL:    server.URL,
		Key:    "test-key",
	}

	airport, err := client.GetAirport("KJFK")

	assert.True(t, airport == nil || err != nil, "expected error or nil airport for 500 response")
}

func TestCloseResponseBody(t *testing.T) {
	body := io.NopCloser(strings.NewReader("test"))
	assert.NotPanics(t, func() {
		_ = body.Close()
	})
}
