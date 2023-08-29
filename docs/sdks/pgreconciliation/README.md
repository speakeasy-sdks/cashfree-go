# PGReconciliation

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
	"context"
	"log"
	"github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/operations"
)

func main() {
    s := cashfree.New(
        cashfree.WithSecurity(shared.Security{
            Option1: &shared.SecurityOption1{
                XClientID: "",
                XClientSecret: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.PGReconciliation.Get(ctx, operations.GetPGReconciliationRequest{
        AcceptMedia: cashfree.String("aspernatur"),
        FetchPGReconRequest: &shared.FetchPGReconRequest{
            Filters: shared.FetchPGReconRequestFilters{
                EndDate: "perferendis",
                StartDate: "ad",
            },
            Pagination: shared.FetchPGReconRequestPagination{
                Cursor: cashfree.String("natus"),
                Limit: 149675,
            },
        },
        XAPIVersion: "iste",
        XIdempotencyKey: cashfree.String("dolor"),
        XRequestID: cashfree.String("natus"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPGReconciliationRequest](../../models/operations/getpgreconciliationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetPGReconciliationResponse](../../models/operations/getpgreconciliationresponse.md), error**

