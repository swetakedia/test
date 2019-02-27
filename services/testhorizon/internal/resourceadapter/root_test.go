package resourceadapter

import (
	"context"
	"net/url"
	"testing"

	"github.com/test/go/clients/testhorizon"
	"github.com/test/go/services/testhorizon/internal/ledger"
	"github.com/stretchr/testify/assert"
)

func TestPopulateRoot(t *testing.T) {
	res := &testhorizon.Root{}
	PopulateRoot(context.Background(),
		res,
		ledger.State{CoreLatest: 1, HistoryLatest: 3, HistoryElder: 2},
		"hVersion",
		"cVersion",
		"passphrase",
		100,
		101,
		urlMustParse(t, "https://friendbot.example.com"))

	assert.Equal(t, int32(1), res.CoreSequence)
	assert.Equal(t, int32(2), res.HistoryElderSequence)
	assert.Equal(t, int32(3), res.TestHorizonSequence)
	assert.Equal(t, "hVersion", res.TestHorizonVersion)
	assert.Equal(t, "cVersion", res.TestCoreVersion)
	assert.Equal(t, "passphrase", res.NetworkPassphrase)
	assert.Equal(t, "https://friendbot.example.com/{?addr}", res.Links.Friendbot.Href)

	// Without testbot
	res = &testhorizon.Root{}
	PopulateRoot(context.Background(),
		res,
		ledger.State{CoreLatest: 1, HistoryLatest: 3, HistoryElder: 2},
		"hVersion",
		"cVersion",
		"passphrase",
		100,
		101,
		nil)

	assert.Equal(t, int32(1), res.CoreSequence)
	assert.Equal(t, int32(2), res.HistoryElderSequence)
	assert.Equal(t, int32(3), res.TestHorizonSequence)
	assert.Equal(t, "hVersion", res.TestHorizonVersion)
	assert.Equal(t, "cVersion", res.TestCoreVersion)
	assert.Equal(t, "passphrase", res.NetworkPassphrase)
	assert.Empty(t, res.Links.Friendbot)
}

func urlMustParse(t *testing.T, s string) *url.URL {
	if u, err := url.Parse(s); err != nil {
		t.Fatalf("Unable to parse URL: %s/%v", s, err)
		return nil
	} else {
		return u
	}
}
