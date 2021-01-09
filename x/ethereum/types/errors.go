package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	// Code 1 is a reserved code for internal errors and should not be used for anything else
	_           = sdkerrors.Register(ModuleName, 1, "internal error")
	ErrEthereum = sdkerrors.Register(ModuleName, 2, "eth bridge error")
)
