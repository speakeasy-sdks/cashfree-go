# Orders

## Overview

Collection of APIs to create, accept payments and refund for an order.

### Available Operations

* [Create](#create) - Create Order
* [Get](#get) - Get Order

## Create

### Order
An order is an entity which has a amount and currency associated with it. It is something for which you want to collect payment for.
Use this API to create orders with Cashfree from your backend to get a `payment_sessions_id`. 
You can use the `payment_sessions_id` to create a transaction for the order.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/operations"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/callbacks"
	"net/http"
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
    res, err := s.Orders.Create(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerBankAccountNumber: cashfree.String("temporibus"),
                CustomerBankCode: cashfree.Float64(710.36),
                CustomerBankIfsc: cashfree.String("quis"),
                CustomerEmail: cashfree.String("veritatis"),
                CustomerID: "deserunt",
                CustomerName: cashfree.String("perferendis"),
                CustomerPhone: "ipsam",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: cashfree.String("2021-07-02T10:20:12+05:30"),
            OrderID: cashfree.String("repellendus"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: cashfree.String("sapiente"),
                PaymentMethods: cashfree.String("quo"),
                ReturnURL: cashfree.String("odit"),
            },
            OrderNote: cashfree.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: cashfree.Float64(8700.88),
                    Percentage: cashfree.Float64(9786.19),
                    VendorID: cashfree.String("molestiae"),
                },
                shared.VendorSplit{
                    Amount: cashfree.Float64(7991.59),
                    Percentage: cashfree.Float64(8009.11),
                    VendorID: cashfree.String("esse"),
                },
                shared.VendorSplit{
                    Amount: cashfree.Float64(5204.78),
                    Percentage: cashfree.Float64(7805.29),
                    VendorID: cashfree.String("dolorum"),
                },
                shared.VendorSplit{
                    Amount: cashfree.Float64(1182.74),
                    Percentage: cashfree.Float64(7206.33),
                    VendorID: cashfree.String("officia"),
                },
            },
            OrderTags: map[string]string{
                "fugit": "deleniti",
                "hic": "optio",
                "totam": "beatae",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "commodi",
                TerminalPhoneNo: "molestiae",
                TerminalType: "modi",
            },
        },
        XAPIVersion: "qui",
        XIdempotencyKey: cashfree.String("impedit"),
        XRequestID: cashfree.String("cum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOrderRequest](../../models/operations/createorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**


## Get

Use this API to fetch the order that was created at Cashfree's using the `order_id`. 
## When to use this API
- To check the status of your order
- Once the order is PAID
- Once your customer returns to `return_url`


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
    res, err := s.Orders.Get(ctx, operations.GetOrderRequest{
        OrderID: "esse",
        XAPIVersion: "ipsum",
        XRequestID: cashfree.String("excepturi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetOrderRequest](../../models/operations/getorderrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetOrderResponse](../../models/operations/getorderresponse.md), error**

