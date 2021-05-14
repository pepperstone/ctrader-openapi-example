package main

import (
	"ctrader/openclient"
	"ctrader/requests"
	"ctrader/stubs/model"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
)

var (
	ClientID     = "2641_2lPgFs39C0aNEP0CbI2t7uQCAulKHN7Mz6izU6aJHA1vU1LJ5c"
	ClientSecret = "jOPor8J7zOhQcAKKPGsTtwuyicIwFlU2fNPdM9NoARySnzSJQh"
	Token        = "PpsP_qldH6qqfbW-ic05bbrSv1fzauTZKXO9VuGrqR8"
	AccountID    = int64(21058862)
	Host         = "demo.ctraderapi.com:5035"
)

func main() {
	client := &openclient.Client{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Address:      Host,
		CertFile:     "server.crt",
		KeyFile:      "server.key",
	}
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			message, err := client.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(message)
			switch uint32(*message.PayloadType) {
			case uint32(model.ProtoOAPayloadType_PROTO_OA_TRADER_RES):
				m := &model.ProtoOATraderRes{}
				proto.Unmarshal(message.Payload, m)
				fmt.Println(m)
				break
			case uint32(model.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_RES):
				fmt.Println(message)
				break
			default:
				fmt.Println(message)
			}
		}
	}()

	err = client.SendMessage(requests.ApplicationAuthReq(ClientID, ClientSecret))
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	err = client.SendMessage(requests.AuthAccountReq(Token, AccountID))
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	err = client.SendMessage(requests.TraderDataReq(AccountID))
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	err = client.SendMessage(requests.SymbolListReq(AccountID))
	if err != nil {
		panic(err)
	}
	// Keep app running
	select {}
}
