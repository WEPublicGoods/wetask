package order

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var LimitOrderExecuteInputABI abi.Arguments

func init() {
	limitOrderExecuteInputType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "order", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "account", Type: "address"},
			{Name: "index", Type: "uint256"},
			{Name: "orderType", Type: "uint256"},
			{Name: "executeFee", Type: "uint256"},
		}},
		{Name: "routes", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "dexId", Type: "uint16"},
			{Name: "tokenIn", Type: "address"},
			{Name: "tokenOut", Type: "address"},
			{Name: "amountIn", Type: "uint256"},
			{Name: "amountOutMin", Type: "uint256"},
			{Name: "extraData", Type: "bytes"},
		}},
		{Name: "tokenIn", Type: "address"},
		{Name: "tokenOut", Type: "address"},
		{Name: "amountIn", Type: "uint256"},
		{Name: "amountOutMin", Type: "uint256"},
		{Name: "amountOutExpected", Type: "uint256"},
		{Name: "feeReceiver", Type: "address"},
	})
	if err != nil {
		panic(err)
	}
	LimitOrderExecuteInputABI = append(LimitOrderExecuteInputABI, abi.Argument{Type: limitOrderExecuteInputType})
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
	Routes            []SwapRoute
	AmountIn          *big.Int
	AmountOutMin      *big.Int
	AmountOutExpected *big.Int
	FeeReceiver       common.Address
}
