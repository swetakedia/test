package resourceadapter

import (
	"testing"

	"github.com/test/go/protocols/testhorizon/operations"
	"github.com/test/go/services/testhorizon/internal/db2/history"
	"github.com/test/go/support/test"
	"github.com/stretchr/testify/assert"
)

// TestPopulateOperation_Successful tests operation object population.
func TestPopulateOperation_Successful(t *testing.T) {
	ctx, _ := test.ContextWithLogBuffer()

	var (
		dest   operations.Base
		row    history.Operation
		val    bool
		ledger = history.Ledger{}
	)

	dest = operations.Base{}
	row = history.Operation{TransactionSuccessful: nil}

	PopulateBaseOperation(ctx, &dest, row, ledger)
	assert.True(t, dest.TransactionSuccessful)

	dest = operations.Base{}
	val = true
	row = history.Operation{TransactionSuccessful: &val}

	PopulateBaseOperation(ctx, &dest, row, ledger)
	assert.True(t, dest.TransactionSuccessful)

	dest = operations.Base{}
	val = false
	row = history.Operation{TransactionSuccessful: &val}

	PopulateBaseOperation(ctx, &dest, row, ledger)
	assert.False(t, dest.TransactionSuccessful)
}
