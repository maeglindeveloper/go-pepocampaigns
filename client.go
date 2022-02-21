package pepo

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Client
type Client struct {
	client    *http.Client
	apiKey    string
	apiSecret string
}

// NewClient instanciates a new pepo client
func NewClient(key, secret string) *Client {
	return &Client{
		client:    &http.Client{},
		apiKey:    key,
		apiSecret: secret,
	}
}

// GetBaseURL returns the base URL for pepocampaigns v1
func (c *Client) GetBaseURL() string {
	return "https://pepocampaigns.com/api/v1"
}

// generates a signature depending on requestTime & url
// https://know.pepocampaigns.com/articles/get-started-with-api/
func (c *Client) generateSignature(url string, requestTime time.Time) string {
	str := url + "::" + requestTime.Format(time.RFC3339)
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) call(ctx context.Context, method string, endpoint string, params *url.Values, body []byte, response interface{}) error {

	endpointURL, _ := url.Parse(c.GetBaseURL())
	endpointURL.Path = path.Join(endpointURL.Path, endpoint) + "/"

	requestTime := time.Now()
	signature := c.generateSignature(endpointURL.Path, requestTime)

	baseParams := url.Values{}
	baseParams.Add("api-key", c.apiKey)
	baseParams.Add("signature", signature)
	baseParams.Add("request-time", requestTime.Format(time.RFC3339))
	if params != nil {
		for k, p := range *params {
			baseParams.Add(k, p[0])
		}
	}

	endpointURL.RawQuery = baseParams.Encode()
	request, err := http.NewRequest(method, endpointURL.String(), bytes.NewReader(body))
	if err != nil {
		return err
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
