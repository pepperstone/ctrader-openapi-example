package main

import (
	"ctrader/stubs/model"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func OnMessage(data []byte) {
	messagePb := &model.ProtoMessage{}
	err := proto.Unmarshal(data, messagePb)
	if err != nil {
		fmt.Println(err)
		return
	}
	// deal with messages
	switch uint32(*messagePb.PayloadType) {
	case uint32(model.ProtoPayloadType_HEARTBEAT_EVENT):
		break
	case uint32(model.ProtoOAPayloadType_PROTO_OA_TRADER_RES):
		m := &model.ProtoOATraderRes{}
		proto.Unmarshal(messagePb.Payload, m)
		fmt.Println(m)
		break
	case uint32(model.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_RES):
		m := &model.ProtoOASymbolsListRes{}
		proto.Unmarshal(messagePb.Payload, m)
		fmt.Println(m)
		break
	default:
		fmt.Println(messagePb)
	}

}

func OnError(data []byte) {
	fmt.Println("OnError", string(data))
}

func OnEnd(data []byte) {
	fmt.Println("OnEnd", string(data))
}
