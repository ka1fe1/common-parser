package common_parser

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
