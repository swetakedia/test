package resourceadapter

import (
	. "github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/db2/core"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
