package limit_keeper

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/WEPublicGoods/wetask/pkg/eth/com"
	"github.com/WEPublicGoods/wetask/pkg/eth/eclient"
	"github.com/WEPublicGoods/wetask/pkg/eth/order"
	"github.com/WEPublicGoods/wetask/pkg/pool"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hibiken/asynq"
	"github.com/tinkler/moonmist/pkg/jsonz/cjson"
)

func parsePayloadFrom(t *asynq.Task) (*payload, error) {
	var p payload
	if err := cjson.Unmarshal(t.Payload(), &p); err != nil {
		return nil, fmt.Errorf("cjson.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	return &p, nil
}

func Handle(ctx context.Context, t *asynq.Task) error {
	p, err := parsePayloadFrom(t)
	if err != nil {
		return err
	}
	client, ok := pool.GetClient(ctx, p.NetworkName)
	if !ok {
		return fmt.Errorf("network %s is not support, %w", p.NetworkName, asynq.SkipRetry)
	}
	orderData, err := order.LimitOrderExecuteInputABI.Pack(
		p.LimitOrder.Order,
		false,
		p.LimitOrder.TokenIn,
		p.LimitOrder.TokenOut,
		p.LimitOrder.RemainingAmountIn,
		p.LimitOrder.Routes,
		p.LimitOrder.AmountIn,
		p.LimitOrder.AmountOutMin,
		p.LimitOrder.AmountOutExpected,
		p.Keeper,
	)
	if err != nil {
		return fmt.Errorf("pack order data %v error:%s, %w", p.LimitOrder, err.Error(), asynq.SkipRetry)
	}
	callable, _, err := checkUpkeep(ctx, client, &bind.CallOpts{
		From: p.Keeper,
	}, p.AutomationCompatibleAddress, orderData)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			return fmt.Errorf("contract is not exist %s, %w", p.AutomationCompatibleAddress.Hex(), asynq.SkipRetry)
		}
		return fmt.Errorf("check upkeep error %s,%w", err.Error(), asynq.SkipRetry)
	}
	if callable {
		transactOpts, err := pool.GetSignedTransactOpts(ctx, p.NetworkName, p.Keeper)
		if err != nil {
			return fmt.Errorf("%s, %w", err.Error(), asynq.SkipRetry)
		}
		if p.BasefeeWiggleMultiplier != nil || p.GasLimitMultiplier > 0 {
			backend, err := client.GetClient(ctx)
			if err != nil {
				return err
			}
			head, err := backend.HeaderByNumber(ctx, nil)
			if err != nil {
				return err
			}
			if head.BaseFee == nil {
				return fmt.Errorf("can not get BaseFee")
			}
			tip, err := backend.SuggestGasTipCap(ctx)
			if err != nil {
				return err
			}
			transactOpts.GasTipCap = tip
			if p.BasefeeWiggleMultiplier != nil {
				transactOpts.GasFeeCap = new(big.Int).Add(
					transactOpts.GasTipCap,
					new(big.Int).Mul(head.BaseFee, p.BasefeeWiggleMultiplier),
				)
			} else {
				transactOpts.GasFeeCap = new(big.Int).Add(
					transactOpts.GasTipCap,
					new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
				)
			}

			if p.GasLimitMultiplier > 0 {
				parsed, err := com.AutomationCompatibleMetaData.GetAbi()
				if err != nil {
					return err
				}
				input, err := parsed.Pack("performUpkeep", orderData)
				if err != nil {
					return err
				}
				msg := ethereum.CallMsg{
					From:      transactOpts.From,
					To:        &p.AutomationCompatibleAddress,
					GasPrice:  nil,
					GasTipCap: transactOpts.GasTipCap,
					GasFeeCap: transactOpts.GasFeeCap,
					Value:     transactOpts.Value,
					Data:      input,
				}
				gasLimit, err := backend.EstimateGas(ctx, msg)
				if err != nil {
					return err
				}
				transactOpts.GasLimit = uint64(float64(gasLimit) * p.GasLimitMultiplier)
			}
		}

		performTx, err := performUpkeep(ctx, client, transactOpts, p.AutomationCompatibleAddress, orderData)
		if err != nil {
			return err
		}
		_, err = client.WaitForReceipt(ctx, performTx.Hash())
		return err
	}

	return nil
}

func parseCancelPayloadFrom(t *asynq.Task) (*cancelPayload, error) {
	var p cancelPayload
	if err := cjson.Unmarshal(t.Payload(), &p); err != nil {
		return nil, fmt.Errorf("cjson.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	return &p, nil
}

func HandleCancel(ctx context.Context, t *asynq.Task) error {
	p, err := parseCancelPayloadFrom(t)
	if err != nil {
		return err
	}
	client, ok := pool.GetClient(ctx, p.NetworkName)
	if !ok {
		return fmt.Errorf("network %s is not support, %w", p.NetworkName, asynq.SkipRetry)
	}
	orderData, err := order.LimitOrderExecuteInputABI.Pack(p.Order,
		true,
		common.Address{},
		common.Address{},
		big.NewInt(0),
		[]order.SwapRoute{},
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		p.Keeper)
	if err != nil {
		return fmt.Errorf("pack order data %v error:%s, %w", p.Order, err.Error(), asynq.SkipRetry)
	}
	callable, _, err := checkUpkeep(ctx, client, &bind.CallOpts{
		From: p.Keeper,
	}, p.AutomationCompatibleAddress, orderData)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			return fmt.Errorf("contract is not exist %s, %w", p.AutomationCompatibleAddress.Hex(), asynq.SkipRetry)
		}
		return fmt.Errorf("check upkeep error %s,%w", err.Error(), asynq.SkipRetry)
	}
	if callable {
		transactOpts, err := pool.GetSignedTransactOpts(ctx, p.NetworkName, p.Keeper)
		if err != nil {
			return fmt.Errorf("%s, %w", err.Error(), asynq.SkipRetry)
		}
		performTx, err := performUpkeep(ctx, client, transactOpts, p.AutomationCompatibleAddress, orderData)
		if err != nil {
			return err
		}
		_, err = client.WaitForReceipt(ctx, performTx.Hash())
		return err
	}

	return nil
}

func checkUpkeep(ctx context.Context, client eclient.Ethclient, opts *bind.CallOpts, automationCompatibleAddress common.Address, checkData []byte) (callable bool, executeData []byte, err error) {
	executeData = make([]byte, 0)

	stub, err := client.GetClient(ctx)
	if err != nil {
		return false, nil, err
	}
	contract, err := com.NewAutomationCompatible(automationCompatibleAddress, stub)
	if err != nil {
		return false, nil, err
	}
	result := []interface{}{
		&callable,
		&executeData,
	}
	results := []interface{}{
		&result,
	}
	err = (&com.AutomationCompatibleRaw{Contract: contract}).
		Call(opts, &results, "checkUpkeep", checkData)
	return
}

func performUpkeep(ctx context.Context, client eclient.Ethclient, opts *bind.TransactOpts, automationCompatibleAddress common.Address, performData []byte) (*types.Transaction, error) {
	stub, err := client.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	contract, err := com.NewAutomationCompatible(automationCompatibleAddress, stub)
	if err != nil {
		return nil, err
	}
	return contract.PerformUpkeep(opts, performData)
}
