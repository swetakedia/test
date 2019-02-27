package ethereum

import (
	"math/big"

	"github.com/test/go/services/bifrost/common"
)

func (t Transaction) ValueToTest() string {
	valueEth := new(big.Rat)
	valueEth.Quo(new(big.Rat).SetInt(t.ValueWei), weiInEth)
	return valueEth.FloatString(common.TestAmountPrecision)
}
