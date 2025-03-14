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

	input := LimitOrderExecuteInput{
		Order:             order,
		TokenIn:           common.HexToAddress("0x4444444444444444444444444444444444444444"),
		TokenOut:          common.HexToAddress("0x5555555555555555555555555555555555555555"),
		Routes:            routes,
		AmountIn:          big.NewInt(10000),
		AmountOutMin:      big.NewInt(9500),
		AmountOutExpected: big.NewInt(10000),
		FeeReceiver:       common.HexToAddress("0x6666666666666666666666666666666666666666"),
	}

	// Pack the input
	packed, err := LimitOrderExecuteInputABI.Pack(input)
	assert.NoError(t, err)
	assert.NotEmpty(t, packed)

	// Unpack the packed data
	_, err = LimitOrderExecuteInputABI.Unpack(packed)
	assert.NoError(t, err)

	// Define anonymous struct type to match ABI unpacking with JSON tags
	// type unpackedStruct struct {
	// 	Order struct {
	// 		Account    common.Address `json:"account"`
	// 		Index      *big.Int       `json:"index"`
	// 		OrderType  *big.Int       `json:"orderType"`
	// 		ExecuteFee *big.Int       `json:"executeFee"`
	// 	} `json:"order"`
	// 	Routes []struct {
	// 		DexId        uint16         `json:"dexId"`
	// 		TokenIn      common.Address `json:"tokenIn"`
	// 		TokenOut     common.Address `json:"tokenOut"`
	// 		AmountIn     *big.Int       `json:"amountIn"`
	// 		AmountOutMin *big.Int       `json:"amountOutMin"`
	// 		ExtraData    []uint8        `json:"extraData"`
	// 	} `json:"routes"`
	// 	TokenIn           common.Address `json:"tokenIn"`
	// 	TokenOut          common.Address `json:"tokenOut"`
	// 	AmountIn          *big.Int       `json:"amountIn"`
	// 	AmountOutMin      *big.Int       `json:"amountOutMin"`
	// 	AmountOutExpected *big.Int       `json:"amountOutExpected"`
	// 	FeeReceiver       common.Address `json:"feeReceiver"`
	// }

	// Get the unpacked struct
	// unpackedInput := unpacked[0].(unpackedStruct)

	// Compare original and unpacked values
	// assert.Equal(t, input.Order.Account, unpackedInput.Order.Account)
	// assert.Equal(t, input.Order.Index, unpackedInput.Order.Index)
	// assert.Equal(t, input.Order.OrderType, unpackedInput.Order.OrderType)
	// assert.Equal(t, input.Order.ExecuteFee, unpackedInput.Order.ExecuteFee)

	// assert.Equal(t, input.TokenIn, unpackedInput.TokenIn)
	// assert.Equal(t, input.TokenOut, unpackedInput.TokenOut)
	// assert.Equal(t, input.AmountIn, unpackedInput.AmountIn)
	// assert.Equal(t, input.AmountOutMin, unpackedInput.AmountOutMin)
	// assert.Equal(t, input.AmountOutExpected, unpackedInput.AmountOutExpected)
	// assert.Equal(t, input.FeeReceiver, unpackedInput.FeeReceiver)

	// assert.Len(t, unpackedInput.Routes, 1)
	// assert.Equal(t, input.Routes[0].DexId, unpackedInput.Routes[0].DexId)
	// assert.Equal(t, input.Routes[0].TokenIn, unpackedInput.Routes[0].TokenIn)
	// assert.Equal(t, input.Routes[0].TokenOut, unpackedInput.Routes[0].TokenOut)
	// assert.Equal(t, input.Routes[0].AmountIn, unpackedInput.Routes[0].AmountIn)
	// assert.Equal(t, input.Routes[0].AmountOutMin, unpackedInput.Routes[0].AmountOutMin)
	// assert.Equal(t, input.Routes[0].ExtraData, unpackedInput.Routes[0].ExtraData)
}
