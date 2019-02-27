package resourceadapter

import (
	. "github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/db2/core"
)

func PopulateAccountThresholds(dest *AccountThresholds, row core.Account) {
	dest.LowThreshold = row.Thresholds[1]
	dest.MedThreshold = row.Thresholds[2]
	dest.HighThreshold = row.Thresholds[3]
}
