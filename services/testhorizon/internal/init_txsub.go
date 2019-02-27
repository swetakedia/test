package testhorizon

import (
	"net/http"

	"github.com/test/go/services/testhorizon/internal/db2/core"
	"github.com/test/go/services/testhorizon/internal/db2/history"
	"github.com/test/go/services/testhorizon/internal/txsub"
	results "github.com/test/go/services/testhorizon/internal/txsub/results/db"
	"github.com/test/go/services/testhorizon/internal/txsub/sequence"
)

func initSubmissionSystem(app *App) {
	cq := &core.Q{Session: app.CoreSession(nil)}

	app.submitter = &txsub.System{
		Pending:         txsub.NewDefaultSubmissionList(),
		Submitter:       txsub.NewDefaultSubmitter(http.DefaultClient, app.config.TestCoreURL),
		SubmissionQueue: sequence.NewManager(),
		Results: &results.DB{
			Core:    cq,
			History: &history.Q{Session: app.TestHorizonSession(nil)},
		},
		Sequences:         cq.SequenceProvider(),
		NetworkPassphrase: app.config.NetworkPassphrase,
	}
}

func init() {
	appInit.Add("txsub", initSubmissionSystem, "app-context", "log", "testhorizon-db", "core-db")
}
