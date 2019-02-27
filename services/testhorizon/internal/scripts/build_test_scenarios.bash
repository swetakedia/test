#! /usr/bin/env bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GOTOP="$( cd "$DIR/../../../../../../../.." && pwd )"
PACKAGES=$(find $GOTOP/src/github.com/test/go/services/testhorizon/internal/test/scenarios -iname '*.rb' -not -name '_common_accounts.rb')
#PACKAGES=$(find $GOTOP/src/github.com/test/go/services/testhorizon/internal/test/scenarios -iname 'failed_transactions.rb')

go install github.com/test/go/services/testhorizon

dropdb hayashi_scenarios --if-exists
createdb hayashi_scenarios

export TEST_CORE_DATABASE_URL="postgres://localhost/hayashi_scenarios?sslmode=disable"
export DATABASE_URL="postgres://localhost/testhorizon_scenarios?sslmode=disable"
export NETWORK_PASSPHRASE="Test SDF Network ; September 2015"
export TEST_CORE_URL="http://localhost:8080"
export SKIP_CURSOR_UPDATE="true"
export INGEST_FAILED_TRANSACTIONS=true

# run all scenarios
for i in $PACKAGES; do
  CORE_SQL="${i%.rb}-core.sql"
  HORIZON_SQL="${i%.rb}-testhorizon.sql"
  scc -r $i --allow-failed-transactions --dump-root-db > $CORE_SQL

  # load the core scenario
  psql $TEST_CORE_DATABASE_URL < $CORE_SQL

  # recreate testhorizon dbs
  dropdb testhorizon_scenarios --if-exists
  createdb testhorizon_scenarios

  # import the core data into testhorizon
  $GOTOP/bin/testhorizon db init
  $GOTOP/bin/testhorizon db init-asset-stats
  $GOTOP/bin/testhorizon db rebase

  # write testhorizon data to sql file
  pg_dump $DATABASE_URL \
    --clean --if-exists --no-owner --no-acl --inserts \
    | sed '/SET idle_in_transaction_session_timeout/d' \
    | sed '/SET row_security/d' \
    > $HORIZON_SQL
done


# commit new sql files to bindata
go generate github.com/test/go/services/testhorizon/internal/test/scenarios
# go test github.com/test/go/services/testhorizon/internal/ingest
