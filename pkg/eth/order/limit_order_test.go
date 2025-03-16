package order

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLimitOrderExecuteInput_PackUnpack(t *testing.T) {
	// Create test data
	order := Order{
		Account:    common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Index:      big.NewInt(1),
		OrderType:  big.NewInt(2),
		ExecuteFee: big.NewInt(100),
	}

	routes := []SwapRoute{
		{
			DexId:        1,
			TokenIn:      common.HexToAddress("0x2222222222222222222222222222222222222222"),
			TokenOut:     common.HexToAddress("0x3333333333333333333333333333333333333333"),
			AmountIn:     big.NewInt(1000),
			AmountOutMin: big.NewInt(900),
			ExtraData:    []byte{0x01, 0x02},
		},
	}

	// Pack the input
	packed, err := LimitOrderExecuteInputABI.Pack(
		order,
		false,
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		common.HexToAddress("0x5555555555555555555555555555555555555555"),
		big.NewInt(1),
		routes,
		big.NewInt(10000),
		big.NewInt(9500),
		big.NewInt(10000),
		common.HexToAddress("0x6666666666666666666666666666666666666666"),
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, packed)

	// Unpack the packed data
	unpacked, err := LimitOrderExecuteInputABI.Unpack(packed)
	assert.NoError(t, err)

	t.Logf("%v", unpacked)

}
