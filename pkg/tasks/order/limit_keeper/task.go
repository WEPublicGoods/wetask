package limit_keeper

import (
	"fmt"
	"math/big"

	ethorder "github.com/WEPublicGoods/wetask/pkg/eth/order"
	"github.com/WEPublicGoods/wetask/pkg/tasks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hibiken/asynq"
	"github.com/tinkler/moonmist/pkg/jsonz/cjson"
)

func NewNormalTask(networkName string, automationCompatibleAddr string, keeper string, order ethorder.LimitOrderExecuteInput, opts ...asynq.Option) (*asynq.Task, error) {
	if networkName == "" {
		return nil, fmt.Errorf("network name cannot be empty")
	}
	if automationCompatibleAddr == "" {
		return nil, fmt.Errorf("automation compatible address cannot be empty")
	}
	if keeper == "" {
		return nil, fmt.Errorf("keeper address cannot be empty")
	}
	// Validate LimitOrder fields
	if order.AmountIn == nil {
		return nil, fmt.Errorf("amountIn cannot be nil")
	}
	// Validate nested Order fields
	if order.Order.Account == (common.Address{}) {
		return nil, fmt.Errorf("order account cannot be empty")
	}
	if order.Order.Index == nil {
		return nil, fmt.Errorf("order index cannot be nil")
	}
	if order.Order.OrderType == nil {
		return nil, fmt.Errorf("order type cannot be nil")
	}
	if order.Order.ExecuteFee == nil {
		return nil, fmt.Errorf("execute fee cannot be nil")
	}

	if !common.IsHexAddress(automationCompatibleAddr) {
		return nil, fmt.Errorf("the address of AutomationCompatible is invalid: %s", automationCompatibleAddr)
	}
	if !common.IsHexAddress(keeper) {
		return nil, fmt.Errorf("the address of keeper is invalid: %s", keeper)
	}
	pl := &payload{
		NetworkName:                 networkName,
		AutomationCompatibleAddress: common.HexToAddress(automationCompatibleAddr),
		Keeper:                      common.HexToAddress(keeper),
		LimitOrder:                  order,
	}
	for _, opt := range opts {
		switch opt := opt.(type) {
		case basefeeWiggleMultiplierOption:
			v := opt.Value().(big.Int)
			if v.Cmp(big.NewInt(2)) < 0 {
				return nil, fmt.Errorf("the basefee wiggle multiplier is less than 2")
			}
			pl.BasefeeWiggleMultiplier = &v
		}
	}
	p, err := cjson.Marshal(pl)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(tasks.ACT_LIMIT_ORDER, p, append([]asynq.Option{asynq.MaxRetry(0)}, opts...)...), nil
}

func NewCancelTask(networkName string, automationCompatibleAddr string, keeper string, order ethorder.Order, opts ...asynq.Option) (*asynq.Task, error) {
	if networkName == "" {
		return nil, fmt.Errorf("network name cannot be empty")
	}
	if automationCompatibleAddr == "" {
		return nil, fmt.Errorf("automation compatible address cannot be empty")
	}
	if keeper == "" {
		return nil, fmt.Errorf("keeper address cannot be empty")
	}
	// Validate nested Order fields
	if order.Account == (common.Address{}) {
		return nil, fmt.Errorf("order account cannot be empty")
	}
	if order.Index == nil {
		return nil, fmt.Errorf("order index cannot be nil")
	}
	if order.OrderType == nil {
		return nil, fmt.Errorf("order type cannot be nil")
	}
	if order.ExecuteFee == nil {
		order.ExecuteFee = big.NewInt(0)
	}
	p, err := cjson.Marshal(cancelPayload{
		NetworkName:                 networkName,
		AutomationCompatibleAddress: common.HexToAddress(automationCompatibleAddr),
		Keeper:                      common.HexToAddress(keeper),
		Order:                       order,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(tasks.ACT_CANCEL_LIMIT_ORDER, p, append([]asynq.Option{asynq.MaxRetry(5)}, opts...)...), nil
}
