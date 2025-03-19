package limit_keeper

import (
	"math/big"

	"github.com/WEPublicGoods/wetask/pkg/eth/order"
	"github.com/ethereum/go-ethereum/common"
)

type payload struct {
	NetworkName                 string
	AutomationCompatibleAddress common.Address
	Keeper                      common.Address
	LimitOrder                  order.LimitOrderExecuteInput
	BasefeeWiggleMultiplier     *big.Int
	GasLimitMultiplier          float64
}

type cancelPayload struct {
	NetworkName                 string
	AutomationCompatibleAddress common.Address
	Keeper                      common.Address
	Order                       order.Order
}
