package requests

import (
	"ctrader/stubs/model"

	"google.golang.org/protobuf/proto"
)

func ApplicationAuthReq(clientID string, clientSecret string) *model.ProtoMessage {
	reqType := model.ProtoOAPayloadType_PROTO_OA_APPLICATION_AUTH_REQ
	authMessage := &model.ProtoOAApplicationAuthReq{
		PayloadType:  &reqType,
		ClientId:     &clientID,
		ClientSecret: &clientSecret,
	}
	payload, _ := proto.Marshal(authMessage)
	payloadType := uint32(reqType)
	message := &model.ProtoMessage{}
	message.Payload = payload
	message.PayloadType = &payloadType
	return message
}

func AuthAccountReq(token string, account int64) *model.ProtoMessage {
	// Payload
	reqType := model.ProtoOAPayloadType_PROTO_OA_ACCOUNT_AUTH_REQ
	authMessage := &model.ProtoOAAccountAuthReq{
		PayloadType:         &reqType,
		AccessToken:         &token,
		CtidTraderAccountId: &account,
	}
	payload, _ := proto.Marshal(authMessage)
	payloadType := uint32(reqType)
	// ProtoMessage
	message := &model.ProtoMessage{}
	message.Payload = payload
	message.PayloadType = &payloadType
	return message
}

func TraderDataReq(account int64) *model.ProtoMessage {
	// Payload
	reqType := model.ProtoOAPayloadType_PROTO_OA_TRADER_REQ
	content := &model.ProtoOATraderReq{
		PayloadType:         &reqType,
		CtidTraderAccountId: &account,
	}
	payload, _ := proto.Marshal(content)
	payloadType := uint32(reqType)
	// ProtoMessage
	message := &model.ProtoMessage{}
	message.Payload = payload
	message.PayloadType = &payloadType
	return message
}

func SymbolListReq(account int64) *model.ProtoMessage {
	// Payload
	reqType := model.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_REQ
	content := &model.ProtoOASymbolsListReq{
		PayloadType:         &reqType,
		CtidTraderAccountId: &account,
	}
	payload, _ := proto.Marshal(content)
	payloadType := uint32(reqType)
	// ProtoMessage
	message := &model.ProtoMessage{}
	message.Payload = payload
	message.PayloadType = &payloadType
	return message
}
