package resourceadapter

import (
	"context"

	. "github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/httpx"
	"github.com/test/go/services/testhorizon/internal/txsub"
	"github.com/test/go/support/render/hal"
)

// Populate fills out the details
func PopulateTransactionSuccess(ctx context.Context, dest *TransactionSuccess, result txsub.Result) {
	dest.Hash = result.Hash
	dest.Ledger = result.LedgerSequence
	dest.Env = result.EnvelopeXDR
	dest.Result = result.ResultXDR
	dest.Meta = result.ResultMetaXDR

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Transaction = lb.Link("/transactions", result.Hash)
	return
}
