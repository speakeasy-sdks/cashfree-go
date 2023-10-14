# PaymentLinks
(*PaymentLinks*)

## Overview

Collection of APIs handle payment links.

### Available Operations

* [Cancel](#cancel) - Cancel Payment Link
* [Create](#create) - Create Payment Link
* [Fetch](#fetch) - Fetch Payment Link Details
* [GetOrders](#getorders) - Get Orders for a Payment Link

## Cancel

Use this API to cancel a payment link. No further payments can be done against a cancelled link. Only a link in ACTIVE status can be cancelled.

### Example Usage

```go
package main

import(
	"context"
	"log"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
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


    var linkID string = "Clifton"

    var xAPIVersion string = "powder"

    var xIdempotencyKey *string = "Avon teal"

    var xRequestID *string = "Northwest"

    ctx := context.Background()
    res, err := s.PaymentLinks.Cancel(ctx, linkID, xAPIVersion, xIdempotencyKey, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkCancelledResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `linkID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The payment link ID which you want to cancel.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `xAPIVersion`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | API version to be used. Format is in YYYY-MM-DD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `xIdempotencyKey`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Idempotency works by saving the resulting status code and body of the first request made for any given idempotency key, regardless of whether it succeeded or failed. Subsequent requests with the same key return the same result, including 500 errors.<br/><br/>Currently supported on all POST calls that uses x-client-id & x-client-secret. To use enable, pass x-idempotency-key in the request header. The value of this header must be unique to each operation you are trying to do. One example can be to use the same order_id that you pass while creating orders <br/> |
| `xRequestID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |


### Response

**[*operations.CancelPaymentLinkResponse](../../models/operations/cancelpaymentlinkresponse.md), error**


## Create

Use this API to create a new payment link. The created payment link url will be available in the API response parameter link_url.

### Example Usage

```go
package main

import(
	"context"
	"log"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
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


    var xAPIVersion string = "online"

    createLinkRequest := &shared.CreateLinkRequest{
        CustomerDetails: shared.LinkCustomerDetailsEntity{
            CustomerPhone: "Extended South",
        },
        LinkAmount: 9967.06,
        LinkCurrency: "abnormally deposit evolve",
        LinkID: "fuchsia Gasoline Screen",
        LinkMeta: &shared.LinkMetaEntity{},
        LinkNotes: map[string]string{
            "key_1": "value_1",
            "key_2": "value_2",
        },
        LinkNotify: &shared.LinkNotifyEntity{},
        LinkPurpose: "physical Ameliorated",
    }

    var xIdempotencyKey *string = "after"

    var xRequestID *string = "overriding"

    ctx := context.Background()
    res, err := s.PaymentLinks.Create(ctx, xAPIVersion, createLinkRequest, xIdempotencyKey, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `xAPIVersion`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | API version to be used. Format is in YYYY-MM-DD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `createLinkRequest`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [*shared.CreateLinkRequest](../../models/shared/createlinkrequest.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Request body to create a payment link at cashfree                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `xIdempotencyKey`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Idempotency works by saving the resulting status code and body of the first request made for any given idempotency key, regardless of whether it succeeded or failed. Subsequent requests with the same key return the same result, including 500 errors.<br/><br/>Currently supported on all POST calls that uses x-client-id & x-client-secret. To use enable, pass x-idempotency-key in the request header. The value of this header must be unique to each operation you are trying to do. One example can be to use the same order_id that you pass while creating orders <br/> |
| `xRequestID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |


### Response

**[*operations.CreatePaymentLinkResponse](../../models/operations/createpaymentlinkresponse.md), error**


## Fetch

Use this API to view all details and status of a payment link.

### Example Usage

```go
package main

import(
	"context"
	"log"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
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


    var linkID string = "grateful"

    var xAPIVersion string = "tomorrow"

    var xRequestID *string = "Account"

    ctx := context.Background()
    res, err := s.PaymentLinks.Fetch(ctx, linkID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `linkID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The payment link ID for which you want to view the details.                                                                |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.FetchPaymentLinkDetailsResponse](../../models/operations/fetchpaymentlinkdetailsresponse.md), error**


## GetOrders

Use this API to view all order details for a payment link.

### Example Usage

```go
package main

import(
	"context"
	"log"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
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


    var linkID string = "violet"

    var xAPIVersion string = "delicious"

    var xRequestID *string = "Account"

    ctx := context.Background()
    res, err := s.PaymentLinks.GetOrders(ctx, linkID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkOrdersResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `linkID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The payment link ID for which you want to view the order details.                                                          |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetPaymentLinkOrdersResponse](../../models/operations/getpaymentlinkordersresponse.md), error**

