package test

import (
	"github.com/test/go/services/testhorizon/internal/test/scenarios"
)

func loadScenario(scenarioName string, includeTestHorizon bool) {
	testCorePath := scenarioName + "-core.sql"
	testhorizonPath := scenarioName + "-testhorizon.sql"

	if !includeTestHorizon {
		testhorizonPath = "blank-testhorizon.sql"
	}

	scenarios.Load(TestCoreDatabaseURL(), testCorePath)
	scenarios.Load(DatabaseURL(), testhorizonPath)
}
