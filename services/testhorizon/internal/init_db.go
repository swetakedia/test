package testhorizon

import (
	"github.com/test/go/services/testhorizon/internal/db2/core"
	"github.com/test/go/services/testhorizon/internal/db2/history"
	"github.com/test/go/support/db"
	"github.com/test/go/support/log"
)

func initTestHorizonDb(app *App) {
	session, err := db.Open("postgres", app.config.DatabaseURL)
	if err != nil {
		log.Panic(err)
	}

	// Make sure MaxIdleConns is equal MaxOpenConns. In case of high variance
	// in number of requests closing and opening connections may slow down TestHorizon.
	session.DB.SetMaxIdleConns(app.config.MaxDBConnections)
	session.DB.SetMaxOpenConns(app.config.MaxDBConnections)
	app.historyQ = &history.Q{session}
}

func initCoreDb(app *App) {
	session, err := db.Open("postgres", app.config.TestCoreDatabaseURL)
	if err != nil {
		log.Panic(err)
	}

	// Make sure MaxIdleConns is equal MaxOpenConns. In case of high variance
	// in number of requests closing and opening connections may slow down TestHorizon.
	session.DB.SetMaxIdleConns(app.config.MaxDBConnections)
	session.DB.SetMaxOpenConns(app.config.MaxDBConnections)
	app.coreQ = &core.Q{session}
}

func init() {
	appInit.Add("testhorizon-db", initTestHorizonDb, "app-context", "log")
	appInit.Add("core-db", initCoreDb, "app-context", "log")
}
