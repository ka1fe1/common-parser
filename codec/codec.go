package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/tendermint/tendermint/types"
)

var (
	AppModules []module.AppModuleBasic
	Encodecfg  params.EncodingConfig
)

func GetTxDecoder() sdk.TxDecoder {
	return Encodecfg.TxConfig.TxDecoder()
}

func GetMarshaler() codec.Codec {
	return Encodecfg.Codec
}

func GetSigningTx(txBytes types.Tx) (signing.Tx, error) {
	Tx, err := GetTxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	signingTx := Tx.(signing.Tx)
	return signingTx, nil
}

func RegisterAppModules(module ...module.AppModuleBasic) {
	AppModules = append(AppModules, module...)
}
