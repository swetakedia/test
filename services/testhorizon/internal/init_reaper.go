package testhorizon

import (
	"github.com/test/go/services/testhorizon/internal/reap"
)

func initReaper(app *App) {
	app.reaper = reap.New(app.config.HistoryRetentionCount, app.TestHorizonSession(nil))
}

func init() {
	appInit.Add("reaper", initReaper, "app-context", "log", "testhorizon-db")
}
