# Refunds
(*Refunds*)

## Overview

Collection of APIs handle refunds.

### Available Operations

* [Create](#create) - Create Refund
* [Get](#get) - Get Refund
* [GetAllforOrder](#getallfororder) - Get All Refunds for an Order

## Create

Use this API to initiate refunds.

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
    res, err := s.Refunds.Create(ctx, operations.CreateRefundRequest{
        CreateRefundRequest: &shared.CreateRefundRequest{
            RefundAmount: 9883.74,
            RefundID: "sapiente",
            RefundNote: cashfreego.String("architecto"),
            RefundSpeed: shared.CreateRefundRequestRefundSpeedInstant.ToPointer(),
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: cashfreego.Float64(2088.76),
                    Percentage: cashfreego.Float64(6350.59),
                    VendorID: cashfreego.String("consequuntur"),
                },
            },
        },
        OrderID: "repellat",
        XAPIVersion: "mollitia",
        XIdempotencyKey: cashfreego.String("occaecati"),
        XRequestID: cashfreego.String("numquam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateRefundRequest](../../models/operations/createrefundrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.CreateRefundResponse](../../models/operations/createrefundresponse.md), error**


## Get

Use this API to fetch a specific refund processed on your Cashfree Account.

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
    orderID := "commodi"
    refundID := "quam"
    xAPIVersion := "molestiae"
    xRequestID := "velit"

    ctx := context.Background()
    res, err := s.Refunds.Get(ctx, orderID, refundID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `orderID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Order or the invoice ID for which you want to view the refund details.                                                     |
| `refundID`                                                                                                                 | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Refund Id of the refund you want to fetch.                                                                                 |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetRefundResponse](../../models/operations/getrefundresponse.md), error**


## GetAllforOrder

Use this API to fetch all refunds processed against an order.

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
    orderID := "error"
    xAPIVersion := "quia"
    xRequestID := "quis"

    ctx := context.Background()
    res, err := s.Refunds.GetAllforOrder(ctx, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `orderID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Order or the invoice ID for which you want to view the refund details.                                                     |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetAllRefundsForOrderResponse](../../models/operations/getallrefundsfororderresponse.md), error**

