package main

import (
	"ctrader/stubs/model"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewMessage(reqType uint32, payload protoreflect.ProtoMessage) []byte {
	message := &model.ProtoMessage{}
	payloadBytes, _ := proto.Marshal(payload)
	message.Payload = payloadBytes
	message.PayloadType = &reqType
	// @todo: handle errors
	bytes, _ := proto.Marshal(message)
	return bytes
}

func ApplicationAuthReq(clientID string, clientSecret string) []byte {
	reqType := model.ProtoOAPayloadType_PROTO_OA_APPLICATION_AUTH_REQ
	authMessage := &model.ProtoOAApplicationAuthReq{
		PayloadType:  &reqType,
		ClientId:     &clientID,
		ClientSecret: &clientSecret,
	}
	return NewMessage(uint32(reqType), authMessage)
}

func AuthAccountReq(token string, account int64) []byte {
	reqType := model.ProtoOAPayloadType_PROTO_OA_ACCOUNT_AUTH_REQ
	authMessage := &model.ProtoOAAccountAuthReq{
		PayloadType:         &reqType,
		AccessToken:         &token,
		CtidTraderAccountId: &account,
	}
	return NewMessage(uint32(reqType), authMessage)
}

func TraderDataReq(account int64) []byte {
	reqType := model.ProtoOAPayloadType_PROTO_OA_TRADER_REQ
	payload := &model.ProtoOATraderReq{
		PayloadType:         &reqType,
		CtidTraderAccountId: &account,
	}
	return NewMessage(uint32(reqType), payload)
}

func SymbolListReq(account int64) []byte {
	reqType := model.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_REQ
	payload := &model.ProtoOASymbolsListReq{
		PayloadType:         &reqType,
		CtidTraderAccountId: &account,
	}
	return NewMessage(uint32(reqType), payload)

}

func SubscribeToEuerusdReq(account int64) []byte {
	symbol := []int64{1}
	reqType := model.ProtoOAPayloadType_PROTO_OA_SUBSCRIBE_SPOTS_REQ
	payload := &model.ProtoOASubscribeSpotsReq{
		PayloadType:         &reqType,
		CtidTraderAccountId: &account,
		SymbolId:            symbol,
	}
	return NewMessage(uint32(reqType), payload)

}
