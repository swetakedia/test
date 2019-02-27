---
title: Stale History
---

A testhorizon server may be configured to reject historical requests when the history is known to be further out of date than the configured threshold.  In such cases, this error is returned.  To resolve this error (provided you are the testhorizon instance's operator) please ensure that the ingestion system is running correctly and importing new ledgers.

## Attributes

As with all errors TestHorizon returns, `stale_history` follows the [Problem Details for HTTP APIs](https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00) draft specification guide and thus has the following attributes:

| Attribute | Type   | Description                                                                                                                     |
| --------- | ----   | ------------------------------------------------------------------------------------------------------------------------------- |
| Type      | URL    | The identifier for the error.  This is a URL that can be visited in the browser.                                                |
| Title     | String | A short title describing the error.                                                                                             |
| Status    | Number | An HTTP status code that maps to the error.                                                                                     |
| Detail    | String | A more detailed description of the error.                                                                                       |
| Instance  | String | A token that uniquely identifies this request. Allows server administrators to correlate a client report with server log files  |

## Example

```shell
$ curl -X GET "https://testhorizon-testnet.test.org/transactions?cursor=1&order=desc"
{
  "type": "stale_history",
  "title": "Historical DB Is Too Stale",
  "status": 503,
  "detail": "This testhorizon instance is configured to reject client requests when it can determine that the history database is lagging too far behind the connected instance of test-core.  If you operate this server, please ensure that the ingestion system is properly running.",
  "instance": "testhorizon-testnet-001.prd.test001.internal.test-ops.com/ngUFNhn76T-078058"
}
```
