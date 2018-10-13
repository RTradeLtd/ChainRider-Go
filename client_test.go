package chainridergo_test

import (
	"fmt"
	"os"
	"testing"

	chainridergo "github.com/RTradeLtd/ChainRider-Go"
)

const (
	testTxHash  = "e71dc1a959a80956ac5dfb385e4113203ba8636ba61cc6fd7c2c41d99026ff48"
	testAddress = "XrWe3E96h7QDRkvwgqY4LmdPBE9txYPfrV"
)

func TestChainRiderGo(t *testing.T) {
	token := os.Getenv("TOKEN")
	if token == "" {
		t.Fatal("TOKEN env var is empty")
	}
	opts := chainridergo.ConfigOpts{
		APIVersion:      "v1",
		DigitalCurrency: "dash",
		Blockchain:      "main",
		Token:           token,
	}
	c, err := chainridergo.NewClient(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if resp, err := c.GetRateLimit(); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}

	if resp, err := c.GetInformation(); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}

	if resp, err := c.TransactionByHash(testTxHash); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}

	if resp, err := c.TransactionsForAddress(testAddress); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}

	if balance, err := c.BalanceForAddress(testAddress); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(balance)
	}
}
