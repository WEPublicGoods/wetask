package limit_keeper

import (
	"github.com/WEPublicGoods/wetask/pkg/eth/order"
	"github.com/ethereum/go-ethereum/common"
)

type payload struct {
	NetworkName                 string
	AutomationCompatibleAddress common.Address
	Keeper                      common.Address
	LimitOrder                  order.LimitOrderExecuteInput
}

type cancelPayload struct {
	NetworkName                 string
	AutomationCompatibleAddress common.Address
	Keeper                      common.Address
	Order                       order.Order
}
