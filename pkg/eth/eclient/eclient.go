package eclient

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"
)

var ErrInvalidRPCs = errors.New("invalid web3 rpcs")

type Ethclient interface {
	ChainID(ctx context.Context) (*big.Int, error)
	Network() string
	GetClient(ctx context.Context) (bind.ContractBackend, error)
	WaitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	UrgeReceipt(ctx context.Context, send func() (common.Hash, error), maxIncreaseTimes int) (*types.Receipt, error)
}

type EthclientPool struct {
	networkName string
	rpc         []string
}

func NewEthclientPool(networkName string, rpc ...string) *EthclientPool {
	if len(rpc) == 0 {
		panic("set at lease one rpc")
	}
	return &EthclientPool{networkName: networkName, rpc: rpc}
}

func (cli *EthclientPool) getRawClient(ctx context.Context) (*ethclient.Client, error) {
	// TODO pool select
	client, err := ethclient.DialContext(ctx, cli.rpc[0])
	if err == nil {
		return client, err
	}

	return nil, ErrInvalidRPCs
}

func (cli *EthclientPool) Network() string {
	return cli.networkName
}

func (cli *EthclientPool) GetClient(ctx context.Context) (bind.ContractBackend, error) {
	// TODO pool select
	client, err := cli.getRawClient(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (cli *EthclientPool) ChainID(ctx context.Context) (*big.Int, error) {
	c, err := cli.getRawClient(ctx)
	if err != nil {
		return nil, err
	}
	return c.ChainID(ctx)
}

func (cli *EthclientPool) WaitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	c, err := cli.getRawClient(ctx)
	if err != nil {
		return nil, err
	}
	for {
		receipt, err := c.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				if ctx.Err() == nil {
					time.Sleep(1 * time.Second)
					continue
				} else {
					return nil, ctx.Err()
				}
			}
			return nil, err
		}
		if receipt != nil && receipt.Status == types.ReceiptStatusFailed {
			return nil, fmt.Errorf("transaction %s failed, %w", txHash.Hex(), asynq.SkipRetry)
		}
		if receipt != nil && receipt.BlockNumber.Cmp(big.NewInt(0)) > 0 {
			return receipt, nil
		}
	}
}

func (cli *EthclientPool) UrgeReceipt(ctx context.Context, send func() (common.Hash, error), maxIncreaseTimes int) (*types.Receipt, error) {
	c, err := cli.getRawClient(ctx)
	if err != nil {
		return nil, err
	}
	for {
		txHash, err := send()
		if err != nil {
			return nil, err
		}
		receipt, err := c.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				if ctx.Err() == nil {
					time.Sleep(1 * time.Second)
					if maxIncreaseTimes == 0 {
						return nil, fmt.Errorf("failed after %d times to increase gas, %w", maxIncreaseTimes, asynq.SkipRetry)
					}
					maxIncreaseTimes--
					continue
				} else {
					return nil, ctx.Err()
				}
			}
			return nil, err
		}
		if receipt != nil && receipt.Status == types.ReceiptStatusFailed {
			return nil, fmt.Errorf("transaction %s failed, %w", txHash.Hex(), asynq.SkipRetry)
		}
		if receipt != nil && receipt.BlockNumber.Cmp(big.NewInt(0)) > 0 {
			return receipt, nil
		}
	}
}
