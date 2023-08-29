# PaymentLinks

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
    res, err := s.PaymentLinks.Cancel(ctx, operations.CancelPaymentLinkRequest{
        LinkID: "laboriosam",
        XAPIVersion: "hic",
        XIdempotencyKey: cashfree.String("saepe"),
        XRequestID: cashfree.String("fuga"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkCancelledResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CancelPaymentLinkRequest](../../models/operations/cancelpaymentlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.PaymentLinks.Create(ctx, operations.CreatePaymentLinkRequest{
        CreateLinkRequest: &shared.CreateLinkRequest{
            CustomerDetails: shared.LinkCustomerDetailsEntity{
                CustomerEmail: cashfree.String("in"),
                CustomerName: cashfree.String("corporis"),
                CustomerPhone: "iste",
            },
            LinkAmount: 4370.32,
            LinkAutoReminders: cashfree.Bool(false),
            LinkCurrency: "saepe",
            LinkExpiryTime: cashfree.String("quidem"),
            LinkID: "architecto",
            LinkMeta: &shared.LinkMetaEntity{
                NotifyURL: cashfree.String("ipsa"),
                PaymentMethods: cashfree.String("reiciendis"),
                ReturnURL: cashfree.String("est"),
                UpiIntent: cashfree.Bool(false),
            },
            LinkMinimumPartialAmount: cashfree.Float64(6531.4),
            LinkNotes: map[string]string{
                "dolores": "dolorem",
                "corporis": "explicabo",
                "nobis": "enim",
            },
            LinkNotify: &shared.LinkNotifyEntity{
                SendEmail: cashfree.Bool(false),
                SendSms: cashfree.Bool(false),
            },
            LinkPartialPayments: cashfree.Bool(false),
            LinkPurpose: "omnis",
        },
        XAPIVersion: "nemo",
        XIdempotencyKey: cashfree.String("minima"),
        XRequestID: cashfree.String("excepturi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreatePaymentLinkRequest](../../models/operations/createpaymentlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.PaymentLinks.Fetch(ctx, operations.FetchPaymentLinkDetailsRequest{
        LinkID: "accusantium",
        XAPIVersion: "iure",
        XRequestID: cashfree.String("culpa"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.FetchPaymentLinkDetailsRequest](../../models/operations/fetchpaymentlinkdetailsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


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
    res, err := s.PaymentLinks.GetOrders(ctx, operations.GetPaymentLinkOrdersRequest{
        LinkID: "doloribus",
        XAPIVersion: "sapiente",
        XRequestID: cashfree.String("architecto"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkOrdersResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPaymentLinkOrdersRequest](../../models/operations/getpaymentlinkordersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetPaymentLinkOrdersResponse](../../models/operations/getpaymentlinkordersresponse.md), error**

