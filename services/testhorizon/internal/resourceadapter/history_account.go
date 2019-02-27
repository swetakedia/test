package resourceadapter

import (
	"context"

	. "github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/db2/history"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
