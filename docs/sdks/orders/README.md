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
    xAPIVersion := "ab"
    createOrderBackendRequest := &shared.CreateOrderBackendRequest{
        CustomerDetails: shared.CustomerDetails{
            CustomerBankAccountNumber: cashfree.String("quis"),
            CustomerBankCode: cashfree.Float64(871.29),
            CustomerBankIfsc: cashfree.String("deserunt"),
            CustomerEmail: cashfree.String("perferendis"),
            CustomerID: "ipsam",
            CustomerName: cashfree.String("repellendus"),
            CustomerPhone: "sapiente",
        },
        OrderAmount: 10.15,
        OrderCurrency: "INR",
        OrderExpiryTime: cashfree.String("2021-07-02T10:20:12+05:30"),
        OrderID: cashfree.String("quo"),
        OrderMeta: &shared.OrderMeta{
            NotifyURL: cashfree.String("odit"),
            PaymentMethods: cashfree.String("at"),
            ReturnURL: cashfree.String("at"),
        },
        OrderNote: cashfree.String("Test order"),
        OrderSplits: []shared.VendorSplit{
            shared.VendorSplit{
                Amount: cashfree.Float64(4736.08),
                Percentage: cashfree.Float64(7991.59),
                VendorID: cashfree.String("quod"),
            },
            shared.VendorSplit{
                Amount: cashfree.Float64(4614.79),
                Percentage: cashfree.Float64(5204.78),
                VendorID: cashfree.String("porro"),
            },
            shared.VendorSplit{
                Amount: cashfree.Float64(6788.8),
                Percentage: cashfree.Float64(1182.74),
                VendorID: cashfree.String("nam"),
            },
            shared.VendorSplit{
                Amount: cashfree.Float64(6399.21),
                Percentage: cashfree.Float64(5820.2),
                VendorID: cashfree.String("fugit"),
            },
        },
        OrderTags: map[string]string{
            "hic": "optio",
            "totam": "beatae",
            "commodi": "molestiae",
        },
        Terminal: &shared.TerminalDetails{
            TerminalID: "modi",
            TerminalPhoneNo: "qui",
            TerminalType: "impedit",
        },
    }
    xIdempotencyKey := "cum"
    xRequestID := "esse"

    ctx := context.Background()
    res, err := s.Orders.Create(ctx, xAPIVersion, createOrderBackendRequest, xIdempotencyKey, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `xAPIVersion`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | API version to be used. Format is in YYYY-MM-DD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `createOrderBackendRequest`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [*shared.CreateOrderBackendRequest](../../models/shared/createorderbackendrequest.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Request body to create an order at cashfree                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| `xIdempotencyKey`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Idempotency works by saving the resulting status code and body of the first request made for any given idempotency key, regardless of whether it succeeded or failed. Subsequent requests with the same key return the same result, including 500 errors.<br/><br/>Currently supported on all POST calls that uses x-client-id & x-client-secret. To use enable, pass x-idempotency-key in the request header. The value of this header must be unique to each operation you are trying to do. One example can be to use the same order_id that you pass while creating orders <br/> |
| `xRequestID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |


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
    orderID := "ipsum"
    xAPIVersion := "excepturi"
    xRequestID := "aspernatur"

    ctx := context.Background()
    res, err := s.Orders.Get(ctx, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `orderID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The order or invoice ID for which you want to view the order  details.                                                     |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetOrderResponse](../../models/operations/getorderresponse.md), error**

