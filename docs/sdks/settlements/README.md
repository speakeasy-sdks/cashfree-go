# Settlements
(*Settlements*)

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
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/operations"
)

func main() {
    s := cashfreego.New(
        cashfreego.WithSecurity(shared.Security{
            Option1: &shared.SecurityOption1{
                XClientID: "",
                XClientSecret: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Settlements.Fetch(ctx, operations.GetSettlementReconciliationRequest{
        AcceptMedia: cashfreego.String("Total Honda Southeast"),
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    324295,
                },
                EndDate: cashfreego.String("who Caesium sometimes"),
                SettlementUtrs: []string{
                    "lumen",
                },
                StartDate: cashfreego.String("auxiliary New cyan"),
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Cursor: cashfreego.String("connect"),
                Limit: 306165,
            },
        },
        XAPIVersion: "scarily",
        XIdempotencyKey: cashfreego.String("model microphone"),
        XRequestID: cashfreego.String("feed holistic"),
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
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


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
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/operations"
)

func main() {
    s := cashfreego.New(
        cashfreego.WithSecurity(shared.Security{
            Option1: &shared.SecurityOption1{
                XClientID: "",
                XClientSecret: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Settlements.GetAll(ctx, operations.GetSettlementsRequest{
        AcceptMedia: cashfreego.String("Van Reggae"),
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    651203,
                },
                EndDate: cashfreego.String("North PCI"),
                SettlementUtrs: []string{
                    "Program",
                },
                StartDate: cashfreego.String("scalable"),
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Cursor: cashfreego.String("violet"),
                Limit: 151768,
            },
        },
        XAPIVersion: "Fish",
        XIdempotencyKey: cashfreego.String("Devolved meh radian"),
        XRequestID: cashfreego.String("Assurance"),
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
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


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
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/operations"
)

func main() {
    s := cashfreego.New(
        cashfreego.WithSecurity(shared.Security{
            Option1: &shared.SecurityOption1{
                XClientID: "",
                XClientSecret: "",
            },
        }),
    )
    orderID := "Avon"
    xAPIVersion := "Coupe"
    xRequestID := "ASCII"

    ctx := context.Background()
    res, err := s.Settlements.GetForOrder(ctx, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.SettlementsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `orderID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Order ID for which you want to view the settlements.                                                                       |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetSettlementsByOrderIDResponse](../../models/operations/getsettlementsbyorderidresponse.md), error**

