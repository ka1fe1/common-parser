package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/tendermint/tendermint/types"
	"github.com/tharsis/ethermint/encoding"
)

var (
	AppModules []module.AppModuleBasic
	Encodecfg  params.EncodingConfig
)

// MakeEncodingConfig 初始化账户地址前缀
func MakeEncodingConfig() {
	moduleBasics := module.NewBasicManager(AppModules...)
	Encodecfg = encoding.MakeConfig(moduleBasics)
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
