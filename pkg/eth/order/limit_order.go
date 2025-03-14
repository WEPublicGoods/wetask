package order

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var LimitOrderABI abi.Arguments

func init() {
	limitOrderType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "order", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "account", Type: "address"},
			{Name: "index", Type: "uint256"},
			{Name: "orderType", Type: "uint256"},
			{Name: "executeFee", Type: "uint256"},
		}},
		{Name: "tokenIn", Type: "address"},
		{Name: "tokenOut", Type: "address"},
		{Name: "amountIn", Type: "uint256"},
		{Name: "amountOut", Type: "uint256"},
		{Name: "orderFee", Type: "uint256"},
		{Name: "deadline", Type: "uint256"},
		{Name: "remainingAmountIn", Type: "uint256"},
		{Name: "filledAmountOut", Type: "uint256"},
	})
	if err != nil {
		panic(err)
	}
	LimitOrderABI = append(LimitOrderABI, abi.Argument{Type: limitOrderType})
}

type LimitOrder struct {
	Order             Order
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	AmountOut         *big.Int
	OrderFee          *big.Int
	Deadline          *big.Int
	RemainingAmountIn *big.Int
	FilledAmountOut   *big.Int
}
