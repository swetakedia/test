package testhorizon

import (
	"log"

	"github.com/test/go/services/testhorizon/internal/ingest"
)

func initIngester(app *App) {
	if !app.config.Ingest {
		return
	}

	if app.config.NetworkPassphrase == "" {
		log.Fatal("Cannot start ingestion without network passphrase.  Please confirm connectivity with test-core.")
	}

	app.ingester = ingest.New(
		app.config.NetworkPassphrase,
		app.config.TestCoreURL,
		app.CoreSession(nil),
		app.TestHorizonSession(nil),
		ingest.Config{
			EnableAssetStats:         app.config.EnableAssetStats,
			IngestFailedTransactions: app.config.IngestFailedTransactions,
		},
	)

	app.ingester.SkipCursorUpdate = app.config.SkipCursorUpdate
	app.ingester.HistoryRetentionCount = app.config.HistoryRetentionCount
}

func init() {
	appInit.Add("ingester", initIngester, "app-context", "log", "testhorizon-db", "core-db", "testCoreInfo")
}
