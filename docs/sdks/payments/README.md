# Payments

## Overview

Collection of APIs handle payments.

### Available Operations

* [Payment](#payment) - Get Payment by ID
* [GetforOrder](#getfororder) - Get Payments for an Order
* [PayOrder](#payorder) - Order Pay
* [PreauthorizeOrder](#preauthorizeorder) - Preauthorization
* [Submit](#submit) - Submit or Resend OTP

## Payment

Use this API to view payment details of an order for a payment ID.

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
    res, err := s.Payments.Payment(ctx, operations.GetPaymentbyIDRequest{
        CfPaymentID: 652790,
        OrderID: "dolorem",
        XAPIVersion: "culpa",
        XRequestID: cashfree.String("consequuntur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPaymentbyIDRequest](../../models/operations/getpaymentbyidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetPaymentbyIDResponse](../../models/operations/getpaymentbyidresponse.md), error**


## GetforOrder

Use this API to view all payment details for an order.

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
    res, err := s.Payments.GetforOrder(ctx, operations.GetPaymentsforOrderRequest{
        OrderID: "repellat",
        XAPIVersion: "mollitia",
        XRequestID: cashfree.String("occaecati"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPaymentsforOrder200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPaymentsforOrderRequest](../../models/operations/getpaymentsfororderrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetPaymentsforOrderResponse](../../models/operations/getpaymentsfororderresponse.md), error**


## PayOrder

Use this API when you have already created the orders and want
# Raj nadna
Cashfree to process the payment. To use this API S2S flag needs to be enabled
from the backend. In case you want to use the cards payment option the PCI
DSS flag is required, for more information send an email to "care@cashfree.com".


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
    res, err := s.Payments.PayOrder(ctx, operations.OrderPayRequest{
        OrderPayRequest: &shared.OrderPayRequest{
            OfferID: cashfree.String("faa6cc05-d1e2-401c-b0cf-0c9db3ff0f0b"),
            PaymentMethod: shared.OrderPayRequestPaymentMethod{},
            PaymentSessionID: "session__CvcEmNKDkmERQrxnx39ibhJ3Ii034pjc8ZVxf3qcgEXCWlgDDlHRgz2XYZCqpajDQSXMMtCusPgOIxYP2LZx0-05p39gC2Vgmq1RAj--gcn",
            SaveInstrument: cashfree.Bool(false),
        },
        XAPIVersion: "numquam",
        XRequestID: cashfree.String("commodi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrderPayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.OrderPayRequest](../../models/operations/orderpayrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.OrderPayResponse](../../models/operations/orderpayresponse.md), error**


## PreauthorizeOrder

Use this API to capture or void a preauthorized payment

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
    res, err := s.Payments.PreauthorizeOrder(ctx, operations.CapturePreauthorizationRequest{
        AuthorizationRequest: &shared.AuthorizationRequest{
            Action: shared.AuthorizationRequestActionCapture.ToPointer(),
            Amount: cashfree.Float64(4746.97),
        },
        OrderID: "velit",
        XAPIVersion: "error",
        XIdempotencyKey: cashfree.String("quia"),
        XRequestID: cashfree.String("quis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CapturePreauthorizationRequest](../../models/operations/capturepreauthorizationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CapturePreauthorizationResponse](../../models/operations/capturepreauthorizationresponse.md), error**


## Submit

If you accept OTP on your own page, you can use the below API to send OTP to Cashfree.

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
    res, err := s.Payments.Submit(ctx, operations.SubmitOTPRequestRequest{
        OTPRequest: &shared.OTPRequest{
            Action: shared.OTPRequestActionSubmitOtp,
            Otp: "laborum",
        },
        PaymentID: "animi",
        XAPIVersion: "enim",
        XRequestID: cashfree.String("odit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OTPResponseEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.SubmitOTPRequestRequest](../../models/operations/submitotprequestrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.SubmitOTPRequestResponse](../../models/operations/submitotprequestresponse.md), error**

