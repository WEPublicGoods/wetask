package limit_keeper

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type backend interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

type keeper struct {
	address      common.Address
	currentNonce uint64
	activeCount  uint64
	mu           sync.Mutex
}

func (k *keeper) GetNonce(ctx context.Context, backend backend) (uint64, func(), error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	// If we have active nonces, increment and return
	if k.activeCount > 0 {
		nonce := k.currentNonce + 1
		k.currentNonce = nonce
		k.activeCount++
		return nonce, func() {
			k.mu.Lock()
			k.activeCount--
			k.mu.Unlock()
		}, nil
	}

	// Get fresh nonce from blockchain
	nonce, err := backend.PendingNonceAt(ctx, k.address)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	k.currentNonce = nonce
	k.activeCount++
	return nonce, func() {
		k.mu.Lock()
		k.activeCount--
		k.mu.Unlock()
	}, nil
}
