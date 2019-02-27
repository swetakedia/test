---
title: TestHorizon Development Guide
---
## TestHorizon Development Guide

This document describes how to build TestHorizon from source, so that you can test and edit the code locally to develop bug fixes and new features.

If you are just starting with TestHorizon and want to try it out, consider the [Quickstart Guide](quickstart.md) instead. For information about administrating a TestHorizon instance in production, check out the [Administration Guide](admin.md).

## Building TestHorizon
Building TestHorizon requires the following developer tools:

- A [Unix-like](https://en.wikipedia.org/wiki/Unix-like) operating system with the common core commands (cp, tar, mkdir, bash, etc.)
- Golang 1.9 or later
- [git](https://git-scm.com/) (to check out TestHorizon's source code)
- [go-dep](https://golang.github.io/dep/) (package manager for Go)
- [mercurial](https://www.mercurial-scm.org/) (needed for `go-dep`)

1. Set your [GOPATH](https://github.com/golang/go/wiki/GOPATH) environment variable, if you haven't already. The default `GOPATH` is `$HOME/go`.
2. Clone the [Test Go](https://github.com/test/go) monorepo:  `go get github.com/test/go`. You should see the repository present at `$GOPATH/src/github.com/test/go`.
3. Enter the source dir: `cd $GOPATH/src/github.com/test/go`, and download external dependencies: `dep ensure -v`. You should see the downloaded third party dependencies in `$GOPATH/pkg`.
4. Compile the TestHorizon binary: `cd $GOPATH; go install github.com/test/go/services/testhorizon`. You should see the resulting `testhorizon` executable in `$GOPATH/bin`.
5. Add Go binaries to your PATH in your `bashrc` or equivalent, for easy access: `export PATH=${GOPATH//://bin:}/bin:$PATH`

Open a new terminal. Confirm everything worked by running `testhorizon --help` successfully. You should see an informative message listing the command line options supported by TestHorizon.

## Set up TestHorizon's database
TestHorizon uses a Postgres database backend to store test fixtures and record information ingested from an associated Test Core. To set this up:
1. Install [PostgreSQL](https://www.postgresql.org/).
2. Run `createdb testhorizon_dev` to initialise an empty database for TestHorizon's use.
3. Run `testhorizon db init --db-url postgres://localhost/testhorizon_dev` to install TestHorizon's database schema.

### Database problems?
1. Depending on your installation's defaults, you may need to configure a Postgres DB user with appropriate permissions for TestHorizon to access the database you created. Refer to the [Postgres documentation](https://www.postgresql.org/docs/current/sql-createuser.html) for details. Note: Remember to restart the Postgres server after making any changes to `pg_hba.conf` (the Postgres configuration file), or your changes won't take effect!
2. Make sure you pass the appropriate database name and user (and port, if using something non-standard) to TestHorizon using `--db-url`. One way is to use a Postgres URI with the following form: `postgres://USERNAME:PASSWORD@localhost:PORT/DB_NAME`.
3. If you get the error `connect failed: pq: SSL is not enabled on the server`, add `?sslmode=disable` to the end of the Postgres URI to allow connecting without SSL.
4. If your server is responding strangely, and you've exhausted all other options, reboot the machine. On some systems `service postgresql restart` or equivalent may not fully reset the state of the server.

## Run tests
At this point you should be able to run TestHorizon's unit tests:
```bash
cd $GOPATH/src/github.com/test/go/services/testhorizon
bash ../../support/scripts/run_tests
```

## Set up Test Core
TestHorizon provides an API to the Test network. It does this by ingesting data from an associated `test-core` instance. Thus, to run a full TestHorizon instance requires a `test-core` instance to be configured, up to date with the network state, and accessible to TestHorizon. TestHorizon accesses `test-core` through both an HTTP endpoint and by connecting directly to the `test-core` Postgres database.

The simplest way to set up Test Core is using the [Test Quickstart Docker Image](https://github.com/test/docker-test-core-testhorizon). This is a Docker container that provides both `test-core` and `testhorizon`, pre-configured for testing.

1. Install [Docker](https://www.docker.com/get-started).
2. Verify your Docker installation works: `docker run hello-world`
3. Create a local directory that the container can use to record state. This is helpful because it can take a few minutes to sync a new `test-core` with enough data for testing, and because it allows you to inspect and modify the configuration if needed. Here, we create a directory called `test` to use as the persistent volume: `cd $HOME; mkdir test`
4. Download and run the Test Quickstart container:

```bash
docker run --rm -it -p "8000:8000" -p "11626:11626" -p "11625:11625" -p"8002:5432" -v $HOME/test:/opt/test --name test test/quickstart --testnet
```

In this example we run the container in interactive mode. We map the container's TestHorizon HTTP port (`8000`), the `test-core` HTTP port (`11626`), and the `test-core` peer node port (`11625`) from the container to the corresponding ports on `localhost`. Importantly, we map the container's `postgresql` port (`5432`) to a custom port (`8002`) on `localhost`, so that it doesn't clash with our local Postgres install.
The `-v` option mounts the `test` directory for use by the container. See the [Quickstart Image documentation](https://github.com/test/docker-test-core-testhorizon) for a detailed explanation of these options.

5. The container is running both a `test-core` and a `testhorizon` instance. Log in to the container and stop TestHorizon:
```bash
docker exec -it test /bin/bash
supervisorctl
stop testhorizon
```

## Check Test Core status
Test Core takes some time to synchronise with the rest of the network. The default configuration will pull roughly a couple of day's worth of ledgers, and may take 15 - 30 minutes to catch up. Logs are stored in the container at `/var/log/supervisor`. You can check the progress by monitoring logs with `supervisorctl`:
```bash
docker exec -it test /bin/bash
supervisorctl tail -f test-core
```

You can also check status by looking at the HTTP endpoint, e.g. by visiting http://localhost:11626 in your browser.

## Connect TestHorizon to Test Core
You can connect TestHorizon to `test-core` at any time, but TestHorizon will not begin ingesting data until `test-core` has completed its catch-up process.

Now run your development version of TestHorizon (which is outside of the container), pointing it at the `test-core` running inside the container:

```bash
testhorizon --db-url="postgres://localhost/testhorizon_dev" --test-core-db-url="postgres://test:postgres@localhost:8002/core" --test-core-url="http://localhost:11626" --port 8001 --network-passphrase "Test SDF Network ; September 2015" --ingest
```

If all is well, you should see ingest logs written to standard out. You can test your TestHorizon instance with a query like: http://localhost:8001/transactions?limit=10&order=asc. Use the [Test Laboratory](https://www.test.org/laboratory/) to craft other queries to try out,
and read about the available endpoints and see examples in the [TestHorizon API reference](https://www.test.org/developers/testhorizon/reference/).

## The development cycle
Congratulations! You can now run the full development cycle to build and test your code.
1. Write code + tests
2. Run tests
3. Compile TestHorizon: `go install github.com/test/go/services/testhorizon`
4. Run TestHorizon (pointing at your running `test-core`)
5. Try TestHorizon queries

Check out the [Test Contributing Guide](https://github.com/test/docs/blob/master/CONTRIBUTING.md) to see how to contribute your work to the Test repositories. Once you've got something that works, open a pull request, linking to the issue that you are resolving with your contribution. We'll get back to you as quickly as we can.
