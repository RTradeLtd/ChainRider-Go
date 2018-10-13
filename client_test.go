package chainridergo_test

import (
	"fmt"
	"os"
	"testing"

	chainridergo "github.com/RTradeLtd/ChainRider-Go"
)

func TestChainRiderGo(t *testing.T) {
	token := os.Getenv("TOKEN")
	if token == "" {
		t.Fatal("TOKEN env var is empty")
	}
	opts := chainridergo.ConfigOpts{
		APIVersion:      "v1",
		DigitalCurrency: "dash",
		Blockchain:      "testnet",
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
}
