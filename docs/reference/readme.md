---
title: Overview
---

The Go SDK contains packages for interacting with most aspects of the test ecosystem.  In addition to generally useful, low-level packages such as [`keypair`](https://godoc.org/github.com/test/go/keypair) (used for creating test-compliant public/secret key pairs), the Go SDK also contains code for the server applications and client tools written in go.

## Godoc reference

The most accurate and up-to-date reference information on the Go SDK is found within godoc.  The godoc.org service automatically updates the documentation for the Go SDK everytime github is updated.  The godoc for all of our packages can be found at (https://godoc.org/github.com/test/go).

## Client Packages

The Go SDK contains packages for interacting with the various test services:

- [`testhorizon`](https://godoc.org/github.com/test/go/clients/testhorizon) provides client access to a testhorizon server, allowing you to load account information, stream payments, post transactions and more.
- [`testtoml`](https://godoc.org/github.com/test/go/clients/testtoml) provides the ability to resolve Test.toml files from the internet.  You can read about [Test.toml concepts here](../../guides/concepts/test-toml.md).
- [`federation`](https://godoc.org/github.com/test/go/clients/federation) makes it easy to resolve a test addresses (e.g. `scott*test.org`) into a test account ID suitable for use within a transaction.

