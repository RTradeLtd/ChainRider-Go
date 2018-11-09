package dash_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/RTradeLtd/ChainRider-Go/dash"
)

const (
	testTxHash             = "e71dc1a959a80956ac5dfb385e4113203ba8636ba61cc6fd7c2c41d99026ff48"
	testAddress            = "XrWe3E96h7QDRkvwgqY4LmdPBE9txYPfrV"
	testBlockHash          = "0000000000000004689b8833fa05a687a042650e4f4e8da9953a5e0dba5fc1a0"
	testDestinationAddress = "yfLFuyfSNHNtwKbfaGXh17maGKAAgd2A4z"
	testPaymentForwardID1  = "Uu0smyqh8mRAOOPZLtV71zYBJzpwpBZt"
	testPaymentForwardID2  = "9risbkgOTkKHq18X198zikCza8GuTJfO"
)

func TestChainRiderGo(t *testing.T) {
	token := os.Getenv("TOKEN")
	if token == "" {
		t.Fatal("TOKEN env var is empty")
	}
	opts := dash.ConfigOpts{
		APIVersion:      "v1",
		DigitalCurrency: "dash",
		Blockchain:      "testnet",
		Token:           token,
	}
	c, err := dash.NewClient(&opts)
	if err != nil {
		t.Fatal(err)
	}
	var forwardID string
	if resp, err := c.CreatePaymentForward(&dash.PaymentForwardOpts{DestinationAddress: testDestinationAddress}); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		forwardID = resp.PaymentForwardID
		fmt.Printf("%+v\n", resp)
	}
	if err = c.DeletePaymentForwardByID(forwardID); err != nil {
		t.Fatal(err)
	}
	t.Skip()
	if resp, err := c.GetPaymentForwardByID(testPaymentForwardID1); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		layout := "2006-01-02T15:04:05.000Z"
		if len(resp.ProcessedTxs) > 0 {
			parsedTime, err := time.Parse(layout, resp.ProcessedTxs[0].ProcessedDate)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("parsed time ", parsedTime)
		}
		fmt.Printf("response\n%+v\n", resp)
	}
	if resp, err := c.GetPaymentForwardByID(testPaymentForwardID2); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		layout := "2006-01-02T15:04:05.000Z"
		if len(resp.ProcessedTxs) > 0 {
			parsedTime, err := time.Parse(layout, resp.ProcessedTxs[0].ProcessedDate)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("parsed time ", parsedTime)
		}
		fmt.Printf("response\n%+v\n", resp)
	}
	t.Skip()
	if resp, err := c.GetRateLimit(); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
	time.Sleep(time.Second * 2)
	if resp, err := c.GetInformation(); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
	time.Sleep(time.Second * 2)
	if resp, err := c.TransactionByHash(testTxHash); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
	time.Sleep(time.Second * 2)
	if resp, err := c.TransactionsForAddress(testAddress); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
	time.Sleep(time.Second * 2)
	if balance, err := c.BalanceForAddress(testAddress); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(balance)
	}
	time.Sleep(time.Second * 2)
	if resp, err := c.GetBlockByHash(testBlockHash); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
	time.Sleep(time.Second * 2)
	if resp, err := c.GetLastBlockHash(); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
	time.Sleep(time.Second * 2)
	if resp, err := c.GetBlockchainSyncStatus(); err != nil {
		t.Fatal(err)
	} else if resp == nil {
		t.Fatal("resp is nil but no error occured, unexpected issue")
	} else {
		fmt.Printf("%+v\n", resp)
	}
}
