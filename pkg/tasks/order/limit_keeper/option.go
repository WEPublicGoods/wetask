package limit_keeper

import (
	"fmt"
	"math/big"

	"github.com/hibiken/asynq"
)

type (
	basefeeWiggleMultiplierOption big.Int
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
