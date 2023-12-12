# PGReconciliation
(*PGReconciliation*)

## Overview

Transac1tio1n reconciliation

### Available Operations

* [Get](#get) - PG Reconciliation

## Get

Use this API to get the payment gateway reconciliation details with date range.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"context"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/operations"
	"log"
)

func main() {
    s := cashfreego.New(
        cashfreego.WithSecurity(shared.Security{
            Option1: &shared.SecurityOption1{
                XClientID: "<YOUR_API_KEY_HERE>",
                XClientSecret: "<YOUR_API_KEY_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.PGReconciliation.Get(ctx, operations.GetPGReconciliationRequest{
        FetchPGReconRequest: &shared.FetchPGReconRequest{
            Filters: shared.Filters{
                EndDate: "string",
                StartDate: "string",
            },
            Pagination: shared.Pagination{
                Limit: 700347,
            },
        },
        XAPIVersion: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchPGRecon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPGReconciliationRequest](../../pkg/models/operations/getpgreconciliationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.GetPGReconciliationResponse](../../pkg/models/operations/getpgreconciliationresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.SDKError            | 400-600                       | */*                           |
