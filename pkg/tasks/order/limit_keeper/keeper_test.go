package limit_keeper

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

type mockBackend struct {
	nonce uint64
	err   error
}

func (m *mockBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	if m.err != nil {
		return 0, m.err
	}
	return m.nonce, nil
}

func (m *mockBackend) waitDone() {
	m.nonce++
}

func TestKeeper_GetNonce(t *testing.T) {
	addr := common.HexToAddress("0x1234")
	ctx := context.Background()

	t.Run("initial nonce from backend", func(t *testing.T) {
		backend := &mockBackend{nonce: 10}
		k := &keeper{address: addr}

		nonce, cleanup, err := k.GetNonce(ctx, backend)
		assert.NoError(t, err)
		assert.Equal(t, uint64(10), nonce)
		assert.NotNil(t, cleanup)
		cleanup()
	})

	t.Run("error from backend", func(t *testing.T) {
		backend := &mockBackend{err: errors.New("backend error")}
		k := &keeper{address: addr}

		_, cleanup, err := k.GetNonce(ctx, backend)
		assert.Error(t, err)
		assert.Nil(t, cleanup)
	})

	t.Run("concurrent nonce generation", func(t *testing.T) {
		backend := &mockBackend{nonce: 100}
		k := &keeper{address: addr}

		var wg sync.WaitGroup
		results := make(chan uint64, 100)

		// Start 100 concurrent goroutines
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				nonce, cleanup, err := k.GetNonce(ctx, backend)
				if err != nil {
					t.Error(err)
					return
				}
				results <- nonce
				backend.waitDone()
				cleanup()
			}()
		}

		wg.Wait()
		close(results)

		// Verify we got unique nonces
		seen := make(map[uint64]bool)
		for nonce := range results {
			if seen[nonce] {
				t.Errorf("duplicate nonce: %d", nonce)
			}
			seen[nonce] = true
		}
	})

	t.Run("cleanup function", func(t *testing.T) {
		backend := &mockBackend{nonce: 50}
		k := &keeper{address: addr}

		// Get initial nonce
		nonce1, cleanup1, err := k.GetNonce(ctx, backend)
		assert.NoError(t, err)
		assert.Equal(t, uint64(50), nonce1)
		defer func() {
			backend.waitDone()
			cleanup1()
		}()

		// Get second nonce
		nonce2, cleanup2, err := k.GetNonce(ctx, backend)
		assert.NoError(t, err)
		assert.Equal(t, uint64(51), nonce2)

		// Cleanup second nonce
		backend.waitDone()
		cleanup2()
	})

	t.Run("stress test", func(t *testing.T) {
		backend := &mockBackend{nonce: 1000}
		k := &keeper{address: addr}

		var wg sync.WaitGroup
		start := make(chan struct{})

		// Start 1000 goroutines that will all try to get nonces at the same time
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-start
				_, cleanup, err := k.GetNonce(ctx, backend)
				if err != nil {
					t.Error(err)
					return
				}
				time.Sleep(time.Millisecond) // Simulate some work
				backend.waitDone()
				cleanup()
			}()
		}

		// Start all goroutines at once
		close(start)
		wg.Wait()
	})
}
