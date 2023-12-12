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
    res, err := s.Refunds.Create(ctx, operations.CreateRefundRequest{
        CreateRefundRequest: &shared.CreateRefundRequest{
            RefundAmount: 4865.89,
            RefundID: "string",
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{},
            },
        },
        OrderID: "string",
        XAPIVersion: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateRefundRequest](../../pkg/models/operations/createrefundrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateRefundResponse](../../pkg/models/operations/createrefundresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.APIError502         | 502                           | application/json              |
| sdkerrors.SDKError            | 400-600                       | */*                           |

## Get

Use this API to fetch a specific refund processed on your Cashfree Account.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"context"
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


    var orderID string = "string"

    var refundID string = "string"

    var xAPIVersion string = "string"

    var xRequestID *string = "string"

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
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetRefundResponse](../../pkg/models/operations/getrefundresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.APIError502         | 502                           | application/json              |
| sdkerrors.SDKError            | 400-600                       | */*                           |

## GetAllforOrder

Use this API to fetch all refunds processed against an order.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"context"
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


    var orderID string = "string"

    var xAPIVersion string = "string"

    var xRequestID *string = "string"

    ctx := context.Background()
    res, err := s.Refunds.GetAllforOrder(ctx, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
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
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetAllRefundsForOrderResponse](../../pkg/models/operations/getallrefundsfororderresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.APIError502         | 502                           | application/json              |
| sdkerrors.SDKError            | 400-600                       | */*                           |
