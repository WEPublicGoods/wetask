package order

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var LimitOrderExecuteInputABI abi.Arguments

func init() {
	addressInputType, err := abi.NewType("address", "", nil)
	if err != nil {
		panic(err)
	}
	uint256InputType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(err)
	}
	{
		// order
		orderInputType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
			{Name: "account", Type: "address"},
			{Name: "index", Type: "uint256"},
			{Name: "orderType", Type: "uint256"},
			{Name: "executeFee", Type: "uint256"},
		})
		if err != nil {
			panic(err)
		}
		LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: orderInputType})
	}
	{
		// isExpired
		boolType, err := abi.NewType("bool", "", nil)
		if err != nil {
			panic(err)
		}
		LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: boolType})
	}
	// tokenIn
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: addressInputType})
	// tokenOut
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: addressInputType})
	// remainingAmountIn
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: uint256InputType})
	{
		// routes
		inputType, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
			{Name: "dexId", Type: "uint16"},
			{Name: "tokenIn", Type: "address"},
			{Name: "tokenOut", Type: "address"},
			{Name: "amountIn", Type: "uint256"},
			{Name: "amountOutMin", Type: "uint256"},
			{Name: "extraData", Type: "bytes"},
		})
		if err != nil {
			panic(err)
		}
		LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: inputType})
	}
	// amountIn
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: uint256InputType})
	// amountOutMin
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: uint256InputType})
	// amountOutExpected
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: uint256InputType})
	// feeReceiver
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: addressInputType})
}

type SwapRoute struct {
	DexId        uint16
	TokenIn      common.Address
	TokenOut     common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	ExtraData    []byte
}

type LimitOrderExecuteInput struct {
	Order             Order
	TokenIn           common.Address
	TokenOut          common.Address
	RemainingAmountIn *big.Int
	Routes            []SwapRoute
	AmountIn          *big.Int
	AmountOutMin      *big.Int
	AmountOutExpected *big.Int
	FeeReceiver       common.Address
}
