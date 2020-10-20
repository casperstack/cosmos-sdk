package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	_ Authorization = &GenericAuthorization{}
)

func NewGenericAuthorization(methodName string) *GenericAuthorization {
	return &GenericAuthorization{
		MessageName: methodName,
	}
}

func (cap GenericAuthorization) MethodName() string {
	return cap.MessageName
}

func (cap GenericAuthorization) Accept(msg sdk.Msg, block tmproto.Header) (allow bool, updated Authorization, delete bool) {
	return true, &cap, false
}