package bitcoin

import (
	"math/big"

	"github.com/test/go/services/bifrost/common"
)

func (t Transaction) ValueToTest() string {
	valueSat := new(big.Int).SetInt64(t.ValueSat)
	valueBtc := new(big.Rat).Quo(new(big.Rat).SetInt(valueSat), satInBtc)
	return valueBtc.FloatString(common.TestAmountPrecision)
}
