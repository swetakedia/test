---
title: TestHorizon Administration Guide
---
## TestHorizon Administration Guide

TestHorizon is responsible for providing an HTTP API to data in the Test network. It ingests and re-serves the data produced by the test network in a form that is easier to consume than the performance-oriented data representations used by test-core.

This document describes how to administer a **production** TestHorizon instance. If you are just starting with TestHorizon and want to try it out, consider the [Quickstart Guide](quickstart.md) instead. For information about developing on the TestHorizon codebase, check out the [Development Guide](developing.md).

## Why run TestHorizon?

The Test Development Foundation runs two TestHorizon servers, one for the public network and one for the test network, free for anyone's use at https://testhorizon.test.org and https://testhorizon-testnet.test.org.  These servers should be fine for development and small scale projects, but it is not recommended that you use them for production services that need strong reliability.  By running TestHorizon within your own infrastructure provides a number of benefits:

  - Multiple instances can be run for redundancy and scalability.
  - Request rate limiting can be disabled.
  - Full operational control without dependency on the Test Development Foundations operations.

## Prerequisites

TestHorizon is dependent upon a test-core server.  TestHorizon needs access to both the SQL database and the HTTP API that is published by test-core. See [the administration guide](https://www.test.org/developers/test-core/learn/admin.html
) to learn how to set up and administer a test-core server.  Secondly, TestHorizon is dependent upon a postgres server, which it uses to store processed core data for ease of use. TestHorizon requires postgres version >= 9.3.

In addition to the two prerequisites above, you may optionally install a redis server to be used for rate limiting requests.

## Installing

To install TestHorizon, you have a choice: either downloading a [prebuilt release for your target architecture](https://github.com/test/go/releases) and operation system, or [building TestHorizon yourself](#Building).  When either approach is complete, you will find yourself with a directory containing a file named `testhorizon`.  This file is a native binary.

After building or unpacking TestHorizon, you simply need to copy the native binary into a directory that is part of your PATH.  Most unix-like systems have `/usr/local/bin` in PATH by default, so unless you have a preference or know better, we recommend you copy the binary there.

To test the installation, simply run `testhorizon --help` from a terminal.  If the help for TestHorizon is displayed, your installation was successful. Note: some shells, such as zsh, cache PATH lookups.  You may need to clear your cache  (by using `rehash` in zsh, for example) before trying to run `testhorizon --help`.


## Building

Should you decide not to use one of our prebuilt releases, you may instead build TestHorizon from source.  To do so, you need to install some developer tools:

- A unix-like operating system with the common core commands (cp, tar, mkdir, bash, etc.)
- A compatible distribution of Go (we officially support Go 1.9 and later)
- [go-dep](https://golang.github.io/dep/)
- [git](https://git-scm.com/)
- [mercurial](https://www.mercurial-scm.org/)


1. Set your [GOPATH](https://github.com/golang/go/wiki/GOPATH) environment variable, if you haven't already. The default `GOPATH` is `$HOME/go`.
2. Clone the Test Go monorepo:  `go get github.com/test/go`. You should see the repository cloned at `$GOPATH/src/github.com/test/go`.
3. Enter the source dir: `cd $GOPATH/src/github.com/test/go`, and download external dependencies: `dep ensure -v`. You should see the downloaded third party dependencies in `$GOPATH/pkg`.
4. Compile the TestHorizon binary: `cd $GOPATH; go install github.com/test/go/services/testhorizon`. You should see the `testhorizon` binary in `$GOPATH/bin`.
5. Add Go binaries to your PATH in your `bashrc` or equivalent, for easy access: `export PATH=${GOPATH//://bin:}/bin:$PATH`

Open a new terminal. Confirm everything worked by running `testhorizon --help` successfully.

Note:  Building directly on windows is not supported.


## Configuring

TestHorizon is configured using command line flags or environment variables.  To see the list of command line flags that are available (and their default values) for your version of TestHorizon, run:

`testhorizon --help`

As you will see if you run the command above, TestHorizon defines a large number of flags, however only three are required:

| flag                    | envvar                      | example                              |
|-------------------------|-----------------------------|--------------------------------------|
| `--db-url`              | `DATABASE_URL`              | postgres://localhost/testhorizon_testnet |
| `--test-core-db-url` | `TEST_CORE_DATABASE_URL` | postgres://localhost/core_testnet    |
| `--test-core-url`    | `TEST_CORE_URL`          | http://localhost:11626               |

`--db-url` specifies the TestHorizon database, and its value should be a valid [PostgreSQL Connection URI](http://www.postgresql.org/docs/9.2/static/libpq-connect.html#AEN38419).  `--test-core-db-url` specifies a test-core database which will be used to load data about the test ledger.  Finally, `--test-core-url` specifies the HTTP control port for an instance of test-core.  This URL should be associated with the test-core that is writing to the database at `--test-core-db-url`.

Specifying command line flags every time you invoke TestHorizon can be cumbersome, and so we recommend using environment variables.  There are many tools you can use to manage environment variables:  we recommend either [direnv](http://direnv.net/) or [dotenv](https://github.com/bkeepers/dotenv).  A template configuration that is compatible with dotenv can be found in the [TestHorizon git repo](https://github.com/test/go/blob/master/services/testhorizon/.env.template).



## Preparing the database

Before the TestHorizon server can be run, we must first prepare the TestHorizon database.  This database will be used for all of the information produced by TestHorizon, notably historical information about successful transactions that have occurred on the test network.  

To prepare a database for TestHorizon's use, first you must ensure the database is blank.  It's easiest to simply create a new database on your postgres server specifically for TestHorizon's use.  Next you must install the schema by running `testhorizon db init`.  Remember to use the appropriate command line flags or environment variables to configure TestHorizon as explained in [Configuring ](#Configuring).  This command will log any errors that occur.

### Postgres configuration

It is recommended to set `random_page_cost=1` in Postgres configuration if you are using SSD storage. With this setting Query Planner will make a better use of indexes, expecially for `JOIN` queries. We have noticed a huge speed improvement for some queries.

## Running

Once your TestHorizon database is configured, you're ready to run TestHorizon.  To run TestHorizon you simply run `testhorizon` or `testhorizon serve`, both of which start the HTTP server and start logging to standard out.  When run, you should see some output that similar to:

```
INFO[0000] Starting testhorizon on :8000                     pid=29013
```

The log line above announces that TestHorizon is ready to serve client requests. Note: the numbers shown above may be different for your installation.  Next we can confirm that TestHorizon is responding correctly by loading the root resource.  In the example above, that URL would be [http://127.0.0.1:8000/] and simply running `curl http://127.0.0.1:8000/` shows you that the root resource can be loaded correctly.


## Ingesting live test-core data

TestHorizon provides most of its utility through ingested data.  Your TestHorizon server can be configured to listen for and ingest transaction results from the connected test-core.  We recommend that within your infrastructure you run one (and only one) TestHorizon process that is configured in this way.   While running multiple ingestion processes will not corrupt the TestHorizon database, your error logs will quickly fill up as the two instances race to ingest the data from test-core.  We may develop a system that coordinates multiple TestHorizon processes in the future, but we would also be happy to include an external contribution that accomplishes this.

To enable ingestion, you must either pass `--ingest=true` on the command line or set the `INGEST` environment variable to "true".

### Ingesting historical data

To enable ingestion of historical data from test-core you need to run `testhorizon db backfill NUM_LEDGERS`. If you're running a full validator with published history archive, for example, you might want to ingest all of history. In this case your `NUM_LEDGERS` should be slightly higher than the current ledger id on the network. You can run this process in the background while your TestHorizon server is up. This continuously decrements the `history.elder_ledger` in your /metrics endpoint until `NUM_LEDGERS` is reached and the backfill is complete. 

### Managing storage for historical data

Over time, the recorded network history will grow unbounded, increasing storage used by the database. TestHorizon expands the data ingested from test-core and needs sufficient disk space. Unless you need to maintain a history archive you may configure TestHorizon to only retain a certain number of ledgers in the database. This is done using the `--history-retention-count` flag or the `HISTORY_RETENTION_COUNT` environment variable. Set the value to the number of recent ledgers you wish to keep around, and every hour the TestHorizon subsystem will reap expired data.  Alternatively, you may execute the command `testhorizon db reap` to force a collection.

### Surviving test-core downtime

TestHorizon tries to maintain a gap-free window into the history of the test-network.  This reduces the number of edge cases that TestHorizon-dependent software must deal with, aiming to make the integration process simpler.  To maintain a gap-free history, TestHorizon needs access to all of the metadata produced by test-core in the process of closing a ledger, and there are instances when this metadata can be lost.  Usually, this loss of metadata occurs because the test-core node went offline and performed a catchup operation when restarted.

To ensure that the metadata required by TestHorizon is maintained, you have several options: You may either set the `CATCHUP_COMPLETE` test-core configuration option to `true` or configure `CATCHUP_RECENT` to determine the amount of time your test-core can be offline without having to rebuild your TestHorizon database.

Unless your node is a full validator and archive publisher we _do not_ recommend using the `CATCHUP_COMPLETE` method, as this will force test-core to apply every transaction from the beginning of the ledger, which will take an ever increasing amount of time. Instead, we recommend you set the `CATCHUP_RECENT` config value. To do this, determine how long of a downtime you would like to survive (expressed in seconds) and divide by ten.  This roughly equates to the number of ledgers that occur within your desired grace period (ledgers roughly close at a rate of one every ten seconds).  With this value set, test-core will replay transactions for ledgers that are recent enough, ensuring that the metadata needed by TestHorizon is present.

### Correcting gaps in historical data

In the section above, we mentioned that TestHorizon _tries_ to maintain a gap-free window.  Unfortunately, it cannot directly control the state of test-core and [so gaps may form](https://www.test.org/developers/software/known-issues.html#gaps-detected) due to extended down time.  When a gap is encountered, TestHorizon will stop ingesting historical data and complain loudly in the log with error messages (log lines will include "ledger gap detected").  To resolve this situation, you must re-establish the expected state of the test-core database and purge historical data from TestHorizon's database.  We leave the details of this process up to the reader as it is dependent upon your operating needs and configuration, but we offer one potential solution:

We recommend you configure the HISTORY_RETENTION_COUNT in TestHorizon to a value less than or equal to the configured value for CATCHUP_RECENT in test-core.  Given this situation any downtime that would cause a ledger gap will require a downtime greater than the amount of historical data retained by TestHorizon.  To re-establish continuity:

1.  Stop TestHorizon.
2.  run `testhorizon db reap` to clear the historical database.
3.  Clear the cursor for TestHorizon by running `test-core -c "dropcursor?id=HORIZON"` (ensure capitilization is maintained).
4.  Clear ledger metadata from before the gap by running `test-core -c "maintenance?queue=true"`.
5.  Restart TestHorizon.    

## Managing Stale Historical Data

TestHorizon ingests ledger data from a connected instance of test-core.  In the event that test-core stops running (or if TestHorizon stops ingesting data for any other reason), the view provided by TestHorizon will start to lag behind reality.  For simpler applications, this may be fine, but in many cases this lag is unacceptable and the application should not continue operating until the lag is resolved.

To help applications that cannot tolerate lag, TestHorizon provides a configurable "staleness" threshold.  Given that enough lag has accumulated to surpass this threshold (expressed in number of ledgers), TestHorizon will only respond with an error: [`stale_history`](./errors/stale-history.md).  To configure this option, use either the `--history-stale-threshold` command line flag or the `HISTORY_STALE_THRESHOLD` environment variable.  NOTE:  non-historical requests (such as submitting transactions or finding payment paths) will not error out when the staleness threshold is surpassed.

## Monitoring

To ensure that your instance of TestHorizon is performing correctly we encourage you to monitor it, and provide both logs and metrics to do so.  

TestHorizon will output logs to standard out.  Information about what requests are coming in will be reported, but more importantly, warnings or errors will also be emitted by default.  A correctly running TestHorizon instance will not output any warning or error log entries.

Metrics are collected while a TestHorizon process is running and they are exposed at the `/metrics` path.  You can see an example at (https://testhorizon-testnet.test.org/metrics).

## I'm Stuck! Help!

If any of the above steps don't work or you are otherwise prevented from correctly setting up TestHorizon, please come to our community and tell us.  Either [post a question at our Stack Exchange](https://test.stackexchange.com/) or [chat with us on slack](http://slack.test.org/) to ask for help.
