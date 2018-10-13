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

type TransactionByHashResponse struct {
	TxID     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		TxID      string `json:"txid"`
		Vout      int    `json:"vout"`
		N         int    `json:"n"`
		ScriptSig struct {
			Hex string `json:"hex"`
			Asm string `json:"asm"`
		} `json:"scriptsig"`
		Addr            string  `json:"addr"`
		ValueSat        int     `json:"valueSat"`
		Value           float64 `json:"value"`
		DoubleSpentTxID string  `json:"doubleSpentTxID,omitempty"`
	} `json:"vin"`
	Vout []struct {
		Value        string `json:"value"` // this might need to be a string
		N            int    `json:"n"`
		ScriptPubKey struct {
			Hex       string   `json:"hex"`
			Asm       string   `json:"asm"`
			Addresses []string `json:"addresses"`
			Type      string   `json:"type"`
		} `json:"scriptPubKey"`
		SpentTxID   string `json:"spentTxId"`
		SpentIndex  int    `json:"spentIndex"`
		SpentHeight int    `json:"spentHeight"`
	} `json:"vout"`
	BlockHash     string  `json:"blockhash"`
	BlockHeight   int     `json:"blockheight"`
	Confirmations int     `json:"confirmations"`
	Time          int     `json:"time"`      // looks like its unix timestamp
	BlockTime     int     `json:"blocktime"` // looks like its unix timestamps
	ValueOut      float64 `json:"valueOut"`
	Size          int     `json:"size"`
	ValueIn       float64 `json:"valueIn"`
	Fees          float64 `json:"fees"`
	TxLock        bool    `json:"txlock"`
}

type TransactionsForAddressResponse struct {
	PagesTotal   int                         `json:"pagesTotal"`
	Transactions []TransactionByHashResponse `json:"txs"`
}
