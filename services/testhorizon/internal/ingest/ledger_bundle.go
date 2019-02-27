package ingest

import (
	"database/sql"
	"fmt"

	"github.com/test/go/services/testhorizon/internal/db2/core"
	"github.com/test/go/support/db"
	"github.com/test/go/support/errors"
)

// Load runs queries against `core` to fill in the records of the bundle.
func (lb *LedgerBundle) Load(db *db.Session) error {
	q := &core.Q{Session: db}
	// Load Header
	err := q.LedgerHeaderBySequence(&lb.Header, lb.Sequence)
	if err != nil {
		// Remove when TestHorizon is able to handle gaps in test-core DB.
		// More info:
		// * https://github.com/test/go/issues/335
		// * https://www.test.org/developers/software/known-issues.html#gaps-detected
		if err == sql.ErrNoRows {
			return errors.New(fmt.Sprintf("Gap detected in test-core database (ledger=%d). More information: https://www.test.org/developers/software/known-issues.html#gaps-detected", lb.Sequence))
		}
		return errors.Wrap(err, "failed to load header")
	}

	// Load transactions
	err = q.TransactionsByLedger(&lb.Transactions, lb.Sequence)
	if err != nil {
		return errors.Wrap(err, "failed to load transactions")
	}

	err = q.TransactionFeesByLedger(&lb.TransactionFees, lb.Sequence)
	if err != nil {
		return errors.Wrap(err, "failed to load transaction fees")
	}

	return nil
}
