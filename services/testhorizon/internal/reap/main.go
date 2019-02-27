// Package reap contains the history reaping subsystem for testhorizon.  This system
// is designed to remove data from the history database such that it does not
// grow indefinitely.  The system can be configured with a number of ledgers to
// maintain at a minimum.
package reap

import (
	"time"

	"github.com/test/go/support/db"
)

// System represents the history reaping subsystem of testhorizon.
type System struct {
	TestHorizonDB      *db.Session
	RetentionCount uint

	nextRun time.Time
}

// New initializes the reaper, causing it to begin polling the test-core
// database for now ledgers and ingesting data into the testhorizon database.
func New(retention uint, testhorizon *db.Session) *System {
	r := &System{
		TestHorizonDB:      testhorizon,
		RetentionCount: retention,
	}

	r.nextRun = time.Now().Add(1 * time.Hour)
	return r
}
