package app

import (
	"github.com/ixofoundation/ixo-blockchain/v3/app/keepers"
	"github.com/ixofoundation/ixo-blockchain/v3/app/params"

	"github.com/cosmos/cosmos-sdk/std"
)

var encodingConfig params.EncodingConfig = MakeEncodingConfig()

func GetEncodingConfig() params.EncodingConfig {
	return encodingConfig
}

// MakeEncodingConfig creates an EncodingConfig.
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	keepers.AppModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
