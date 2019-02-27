package testhorizon

import (
	"encoding/json"
	"testing"

	"github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/test"
)

func TestRootAction(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	server := test.NewStaticMockServer(`{
			"info": {
				"network": "test",
				"build": "test-core",
				"ledger": {
					"version": 3
				},
				"protocol_version": 4
			}
		}`)
	defer server.Close()

	ht.App.testhorizonVersion = "test-testhorizon"
	ht.App.config.TestCoreURL = server.URL
	ht.App.config.NetworkPassphrase = "test"
	ht.App.UpdateTestCoreInfo()

	w := ht.Get("/")

	if ht.Assert.Equal(200, w.Code) {
		var actual testhorizon.Root
		err := json.Unmarshal(w.Body.Bytes(), &actual)
		ht.Require.NoError(err)
		ht.Assert.Equal("test-testhorizon", actual.TestHorizonVersion)
		ht.Assert.Equal("test-core", actual.TestCoreVersion)
		ht.Assert.Equal(int32(4), actual.CoreSupportedProtocolVersion)
		ht.Assert.Equal(int32(3), actual.CurrentProtocolVersion)
	}
}
