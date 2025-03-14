package limit_keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/WEPublicGoods/wetask/pkg/eth/com"
	"github.com/WEPublicGoods/wetask/pkg/eth/eclient"
	"github.com/WEPublicGoods/wetask/pkg/eth/order"
	"github.com/WEPublicGoods/wetask/pkg/pool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hibiken/asynq"
	"github.com/tinkler/moonmist/pkg/jsonz/cjson"
)

func parseFrom(t *asynq.Task) (*payload, error) {
	var p payload
	if err := cjson.Unmarshal(t.Payload(), &p); err != nil {
		return nil, fmt.Errorf("cjson.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	return &p, nil
}

func Handle(ctx context.Context, t *asynq.Task) error {
	p, err := parseFrom(t)
	if err != nil {
		return err
	}
	client, ok := pool.GetClient(ctx, p.NetworkName)
	if !ok {
		return fmt.Errorf("network %s is not support, %w", p.NetworkName, asynq.SkipRetry)
	}
	orderData, err := order.LimitOrderABI.Pack(p.LimitOrder)
	if err != nil {
		return fmt.Errorf("pack order data %v error:%s, %w", p.LimitOrder, err.Error(), asynq.SkipRetry)
	}
	callable, _, err := p.checkUpkeep(ctx, client, &bind.CallOpts{
		From: p.Keeper,
	}, orderData)
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
		performTx, err := p.performUpkeep(ctx, client, transactOpts, orderData)
		if err != nil {
			return err
		}
		_, err = client.WaitForReceipt(ctx, performTx.Hash())
		return err
	}

	return nil
}

func (p *payload) checkUpkeep(ctx context.Context, client eclient.Ethclient, opts *bind.CallOpts, checkData []byte) (callable bool, executeData []byte, err error) {
	executeData = make([]byte, 0)

	stub, err := client.GetClient(ctx)
	if err != nil {
		return false, nil, err
	}
	contract, err := com.NewAutomationCompatible(p.AutomationCompatibleAddress, stub)
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

func (p *payload) performUpkeep(ctx context.Context, client eclient.Ethclient, opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	stub, err := client.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	contract, err := com.NewAutomationCompatible(p.AutomationCompatibleAddress, stub)
	if err != nil {
		return nil, err
	}
	return contract.PerformUpkeep(opts, performData)
}
