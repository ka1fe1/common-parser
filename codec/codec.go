package codec

import (
	"github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

var (
	AppModules []module.AppModuleBasic
	Encodecfg  EncodingConfig
)

type EncodingConfig struct {
	InterfaceRegistry ctypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

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
