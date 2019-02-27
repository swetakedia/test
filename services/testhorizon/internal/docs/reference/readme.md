---
title: Overview
---

TestHorizon is an API server for the Test ecosystem.  It acts as the interface between [test-core](https://github.com/test/test-core) and applications that want to access the Test network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Test ecosystem](https://www.test.org/developers/guides/) for details of where TestHorizon fits in. You can also watch a [talk on TestHorizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Test.org developer Scott Fleckenstein:

[![TestHorizon: API webserver for the Test network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "TestHorizon: API webserver for the Test network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

TestHorizon provides a RESTful API to allow client applications to interact with the Test network. You can communicate with TestHorizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Test SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.test.org/developers/js-test-sdk/learn/index.html) for clients to use to interact with TestHorizon.

SDF runs a instance of TestHorizon that is connected to the test net: [https://testhorizon-testnet.test.org/](https://testhorizon-testnet.test.org/) and one that is connected to the public Test network:
[https://testhorizon.test.org/](https://testhorizon.test.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/test/js-test-sdk)
- [Java](https://github.com/test/java-test-sdk)
- [Go](https://github.com/test/go)

Community maintained libraries (in various states of completeness) for interacting with TestHorizon in other languages:<br>
- [Ruby](https://github.com/test/ruby-test-sdk)
- [Python](https://github.com/TestCN/py-test-base)
- [C#](https://github.com/elucidsoft/dotnet-test-sdk)
