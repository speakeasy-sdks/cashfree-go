# Refunds

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
    res, err := s.Refunds.Create(ctx, operations.CreateRefundRequest{
        CreateRefundRequest: &shared.CreateRefundRequest{
            RefundAmount: 7783.46,
            RefundID: "sequi",
            RefundNote: cashfree.String("tenetur"),
            RefundSpeed: shared.CreateRefundRequestRefundSpeedStandard.ToPointer(),
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: cashfree.Float64(8209.94),
                    Percentage: cashfree.Float64(135.71),
                    VendorID: cashfree.String("quasi"),
                },
                shared.VendorSplit{
                    Amount: cashfree.Float64(6228.46),
                    Percentage: cashfree.Float64(8379.45),
                    VendorID: cashfree.String("laborum"),
                },
                shared.VendorSplit{
                    Amount: cashfree.Float64(960.98),
                    Percentage: cashfree.Float64(9719.45),
                    VendorID: cashfree.String("voluptatibus"),
                },
            },
        },
        OrderID: "vero",
        XAPIVersion: "nihil",
        XIdempotencyKey: cashfree.String("praesentium"),
        XRequestID: cashfree.String("voluptatibus"),
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
    res, err := s.Refunds.Get(ctx, operations.GetRefundRequest{
        OrderID: "ipsa",
        RefundID: "omnis",
        XAPIVersion: "voluptate",
        XRequestID: cashfree.String("cum"),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetRefundRequest](../../models/operations/getrefundrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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
    res, err := s.Refunds.GetAllforOrder(ctx, operations.GetAllRefundsForOrderRequest{
        OrderID: "perferendis",
        XAPIVersion: "doloremque",
        XRequestID: cashfree.String("reprehenderit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetAllRefundsForOrderRequest](../../models/operations/getallrefundsfororderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetAllRefundsForOrderResponse](../../models/operations/getallrefundsfororderresponse.md), error**

