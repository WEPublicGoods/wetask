package pool

import (
	"context"
	"errors"

	"github.com/WEPublicGoods/wetask/pkg/eth/eclient"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type contextKey string

const clientsKey contextKey = "pool"
const walletsKey contextKey = "wallet"

func WithPool(ctx context.Context, clients []eclient.Ethclient, wallet *keystore.KeyStore) context.Context {
	if len(clients) == 0 {
		panic("no client")
	}
	if wallet == nil {
		panic("no wallet")
	}
	clientMap := make(map[string]eclient.Ethclient, len(clients))
	for _, c := range clients {
		clientMap[c.Network()] = c
	}
	ctx = context.WithValue(ctx, clientsKey, clientMap)
	ctx = context.WithValue(ctx, walletsKey, wallet)
	return ctx
}

func GetClient(ctx context.Context, networkName string) (eclient.Ethclient, bool) {
	i := ctx.Value(clientsKey)
	clientMap, ok := i.(map[string]eclient.Ethclient)
	if !ok {
		return nil, false
	}
	c, ok := clientMap[networkName]
	return c, ok
}

func GetAccount(ctx context.Context, keeper common.Address) (*accounts.Account, error) {
	i := ctx.Value(walletsKey)
	wallet, ok := i.(*keystore.KeyStore)
	if !ok {
		return nil, errors.New("context is invalid")
	}
	account, err := wallet.Find(accounts.Account{Address: keeper})
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func GetSignedTransactOpts(ctx context.Context, networkName string, keeper common.Address) (*bind.TransactOpts, error) {
	i := ctx.Value(walletsKey)
	wallet, ok := i.(*keystore.KeyStore)
	if !ok {
		return nil, errors.New("context is invalid")
	}
	client, ok := GetClient(ctx, networkName)
	if !ok {
		return nil, errors.New("network " + networkName + " is not support")
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	account, err := wallet.Find(accounts.Account{Address: keeper})
	if err != nil {
		return nil, err
	}
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, t, chainId)
		},
	}, nil
}
