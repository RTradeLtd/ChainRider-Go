package chainridergo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// NewClient is used to initialize our ChainRider client
func NewClient(opts *ConfigOpts) (*Client, error) {
	if opts == nil {
		opts = &ConfigOpts{
			APIVersion:      defaultAPIVersion,
			DigitalCurrency: defaultDigitalCurrency,
			Blockchain:      defaultBlockchain,
			Token:           "test",
		}
	}
	urlFormatted := fmt.Sprintf(urlTemplate, opts.APIVersion, opts.DigitalCurrency, opts.Blockchain)
	c := &Client{
		URL:   urlFormatted,
		Token: opts.Token,
		HC:    &http.Client{},
	}
	// generate our payload
	c.GeneratePayload()
	return c, nil
}

// GeneratePayload is used to generate our payload
func (c *Client) GeneratePayload() {
	c.Payload = fmt.Sprintf(payloadTemplate, c.Token)
}

// GetRateLimit is used to get our rate limit
func (c *Client) GetRateLimit() (*RateLimitResponse, error) {
	req, err := http.NewRequest("POST", rateLimitURL, strings.NewReader(c.Payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.HC.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	intf := RateLimitResponse{}
	if err = json.Unmarshal(bodyBytes, &intf); err != nil {
		return nil, err
	}
	return &intf, nil
}
