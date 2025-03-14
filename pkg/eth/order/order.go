package order

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var OrderABI abi.Arguments

func init() {
	var (
		order abi.Argument
		err   error
	)
	order.Type, err = abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "account", Type: "address"},
		{Name: "index", Type: "uint256"},
		{Name: "orderType", Type: "uint256"},
		{Name: "executeFee", Type: "uint256"},
	})
	if err != nil {
		panic(err)
	}
	OrderABI = append(OrderABI, order)
}

type Order struct {
	Account    common.Address
	Index      *big.Int
	OrderType  *big.Int
	ExecuteFee *big.Int
}
