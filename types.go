package chainridergo

import (
	"net/http"
)

type Client struct {
	URL     string `json:"url"`
	Payload string `json:"payload"`
	Token   string `json:"token"`
	HC      *http.Client
}

type ConfigOpts struct {
	APIVersion      string `json:"api_version"`
	DigitalCurrency string `json:"digital_currency"`
	Blockchain      string `json:"blockchain"`
	Token           string `json:"token"`
}

const (
	payloadTemplate        = "{\n  \"token\": \"%s\"\n}"
	urlTemplate            = "https://api.chainrider.io/%s/%s/%s"
	rateLimitURL           = "https://api.chainrider.io/v1/ratelimit/"
	defaultAPIVersion      = "v1"
	defaultDigitalCurrency = "dash"
	defaultBlockchain      = "testnet"
)

type RateLimitResponse struct {
	Message struct {
		Hour struct {
			Usage    int `json:"usage"`
			Limit    int `json:"limit"`
			TimeLeft int `json:"time_left"`
		} `json:"hour"`
		Day struct {
			Usage    int `json:"usage"`
			Limit    int `json:"limit"`
			TimeLeft int `json:"time_left"`
		} `json:"day"`
		Forward struct {
			Usage    int `json:"usage"`
			Limit    int `json:"limit"`
			TimeLeft int `json:"time_left"`
		} `json:"forward"`
	} `json:"message"`
}

type InformationResponse struct {
	Info struct {
		Version         int     `json:"version"`
		InsightVerison  string  `json:"insightversion"`
		ProtocolVersion int     `json:"protocolversion"`
		Blocks          int     `json:"blocks"`
		TimeOffset      int     `json:"timeoffset"`
		Connections     int     `json:"connections"`
		Proxy           string  `json:"proxy"`
		Difficulty      float64 `json:"difficulty"`
		Testnet         bool    `json:"testnet"`
		RelayFee        float64 `json:"relayfee"`
		Errors          string  `json:"errors"`
		Network         string  `json:"network"`
	} `json:"info"`
}