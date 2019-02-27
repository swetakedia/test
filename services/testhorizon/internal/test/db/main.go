// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on testhorizon and so should be able to be imported by
// any testhorizon package.
package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
	db "github.com/test/go/support/db/dbtest"
)

var (
	coreDB     *sqlx.DB
	coreUrl    *string
	testhorizonDB  *sqlx.DB
	testhorizonUrl *string
)

// TestHorizon returns a connection to the testhorizon test database
func TestHorizon(t *testing.T) *sqlx.DB {
	if testhorizonDB != nil {
		return testhorizonDB
	}
	postgres := db.Postgres(t)
	testhorizonUrl = &postgres.DSN
	testhorizonDB = postgres.Open()
	return testhorizonDB
}

// TestHorizonURL returns the database connection the url any test
// use when connecting to the history/testhorizon database
func TestHorizonURL() string {
	if testhorizonUrl == nil {
		log.Panic(fmt.Errorf("TestHorizon not initialized"))
	}
	return *testhorizonUrl
}

// TestCore returns a connection to the test core test database
func TestCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// TestCoreURL returns the database connection the url any test
// use when connecting to the test-core database
func TestCoreURL() string {
	if coreUrl == nil {
		log.Panic(fmt.Errorf("TestCore not initialized"))
	}
	return *coreUrl
}
