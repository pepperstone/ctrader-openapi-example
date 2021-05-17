package main

import (
	"fmt"
	"time"

	OpenAPI "github.com/ahmad-pepperstone/openapi-go"
)

var (
	ClientID     = "2641_2lPgFs39C0aNEP0CbI2t7uQCAulKHN7Mz6izU6aJHA1vU1LJ5c"
	ClientSecret = "jOPor8J7zOhQcAKKPGsTtwuyicIwFlU2fNPdM9NoARySnzSJQh"
	Token        = "PpsP_qldH6qqfbW-ic05bbrSv1fzauTZKXO9VuGrqR8"
	AccountID    = int64(21058862)
	Host         = "demo.ctraderapi.com:5035"
)

func main() {

	client := OpenAPI.NewClient(OpenAPI.ClientConfig{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Address:      Host,
		CertFile:     "server.crt",
		KeyFile:      "server.key",
	})

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	// Messages are published as []byte
	client.On("message", OnMessage)
	client.On("error", OnMessage)
	client.On("end", OnEnd)
	// Send few requests to the server
	fmt.Println("-- Authorize app --")
	err = client.SendMessage(ApplicationAuthReq(ClientID, ClientSecret))
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n-- Authorize Account --")
	err = client.SendMessage(AuthAccountReq(Token, AccountID))
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

	fmt.Println("\n-- Request Trader Data --")
	err = client.SendMessage(TraderDataReq(AccountID))
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("\n-- Request Symbols List --")
	err = client.SendMessage(SymbolListReq(AccountID))
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("\n-- Subscribe to EURUSD prices --")
	err = client.SendMessage(SubscribeToEuerusdReq(AccountID))
	if err != nil {
		panic(err)
	}

	select {}
}
