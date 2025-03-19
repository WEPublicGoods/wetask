package limit_keeper

import (
	"fmt"
	"math/big"

	"github.com/hibiken/asynq"
)

type (
	basefeeWiggleMultiplierOption big.Int
	gasLimitMultiplierOption      float64
)

func (n basefeeWiggleMultiplierOption) String() string {
	return fmt.Sprintf("BasefeeWiggleMultiplier(%s)", ((*big.Int)(&n)).String())
}

func (n basefeeWiggleMultiplierOption) Type() asynq.OptionType { return asynq.OptionType(10) }

func (n basefeeWiggleMultiplierOption) Value() interface{} { return big.Int(n) }

// default basefee wiggle multiplier is 2
func BasefeeWiggleMultiplier(n *big.Int) asynq.Option {
	if n == nil {
		n = big.NewInt(2)
	}
	return basefeeWiggleMultiplierOption(*n)
}

func (n gasLimitMultiplierOption) String() string {
	return fmt.Sprintf("GasLimitMultiplier(%f)", float64(n))
}

func (n gasLimitMultiplierOption) Type() asynq.OptionType { return asynq.OptionType(11) }

func (n gasLimitMultiplierOption) Value() interface{} { return float64(n) }

// default gas limit multiplier is 1.0
func GasLimitMultiplier(n float64) asynq.Option {
	if n == 0 {
		n = 1.0
	}
	return gasLimitMultiplierOption(n)
}
