package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	models "github.com/kaifei-bianjie/common-parser/types"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin
)
