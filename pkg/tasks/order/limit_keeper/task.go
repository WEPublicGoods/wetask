package limit_keeper

import (
	"fmt"

	ethorder "github.com/WEPublicGoods/wetask/pkg/eth/order"
	"github.com/WEPublicGoods/wetask/pkg/tasks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hibiken/asynq"
	"github.com/tinkler/moonmist/pkg/jsonz/cjson"
)

func NewNormalTask(networkName string, automationCompatibleAddr string, keeper string, order ethorder.LimitOrder) (*asynq.Task, error) {
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
	if order.TokenIn == (common.Address{}) {
		return nil, fmt.Errorf("tokenIn address cannot be empty")
	}
	if order.TokenOut == (common.Address{}) {
		return nil, fmt.Errorf("tokenOut address cannot be empty")
	}
	if order.AmountIn == nil {
		return nil, fmt.Errorf("amountIn cannot be nil")
	}
	if order.AmountOut == nil {
		return nil, fmt.Errorf("amountOut cannot be nil")
	}
	if order.OrderFee == nil {
		return nil, fmt.Errorf("orderFee cannot be nil")
	}
	if order.Deadline == nil {
		return nil, fmt.Errorf("deadline cannot be nil")
	}
	if order.RemainingAmountIn == nil {
		return nil, fmt.Errorf("remainingAmountIn cannot be nil")
	}
	if order.FilledAmountOut == nil {
		return nil, fmt.Errorf("filledAmountOut cannot be nil")
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

	p, err := cjson.Marshal(payload{
		NetworkName:                 networkName,
		AutomationCompatibleAddress: common.HexToAddress(automationCompatibleAddr),
		Keeper:                      common.HexToAddress(keeper),
		LimitOrder:                  order,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(tasks.ACT_LIMIT_ORDER, p, asynq.MaxRetry(5)), nil
}
