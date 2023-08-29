# Settlements

## Overview

Collection of APIs handle settlements.

### Available Operations

* [Fetch](#fetch) - Settlement Reconciliation
* [GetAll](#getall) - Get All Settlements
* [GetForOrder](#getfororder) - Get Settlements by Order ID

## Fetch

Use this API to get settlement reconciliation details using Settlement ID, settlement UTR or date range.

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
    res, err := s.Settlements.Fetch(ctx, operations.GetSettlementReconciliationRequest{
        AcceptMedia: cashfree.String("ut"),
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    120196,
                    359444,
                    296140,
                    480894,
                },
                EndDate: cashfree.String("dicta"),
                SettlementUtrs: []string{
                    "enim",
                    "accusamus",
                    "commodi",
                },
                StartDate: cashfree.String("repudiandae"),
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Cursor: cashfree.String("quae"),
                Limit: 216822,
            },
        },
        XAPIVersion: "quidem",
        XIdempotencyKey: cashfree.String("molestias"),
        XRequestID: cashfree.String("excepturi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchSettlementRecon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetSettlementReconciliationRequest](../../models/operations/getsettlementreconciliationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetSettlementReconciliationResponse](../../models/operations/getsettlementreconciliationresponse.md), error**


## GetAll

Use this API to get all settlement details by specifying the settlement ID, settlement UTR or date range.

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
    res, err := s.Settlements.GetAll(ctx, operations.GetSettlementsRequest{
        AcceptMedia: cashfree.String("pariatur"),
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    508969,
                    523248,
                },
                EndDate: cashfree.String("voluptates"),
                SettlementUtrs: []string{
                    "repudiandae",
                },
                StartDate: cashfree.String("sint"),
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Cursor: cashfree.String("veritatis"),
                Limit: 929297,
            },
        },
        XAPIVersion: "incidunt",
        XIdempotencyKey: cashfree.String("enim"),
        XRequestID: cashfree.String("consequatur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchSettlement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetSettlementsRequest](../../models/operations/getsettlementsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetSettlementsResponse](../../models/operations/getsettlementsresponse.md), error**


## GetForOrder

Use this API to view all the settlements of a particular order.

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
    res, err := s.Settlements.GetForOrder(ctx, operations.GetSettlementsByOrderIDRequest{
        OrderID: "est",
        XAPIVersion: "quibusdam",
        XRequestID: cashfree.String("explicabo"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SettlementsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetSettlementsByOrderIDRequest](../../models/operations/getsettlementsbyorderidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetSettlementsByOrderIDResponse](../../models/operations/getsettlementsbyorderidresponse.md), error**

