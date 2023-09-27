# Payments
(*Payments*)

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
    cfPaymentID := 697631
    orderID := "architecto"
    xAPIVersion := "ipsa"
    xRequestID := "reiciendis"

    ctx := context.Background()
    res, err := s.Payments.Payment(ctx, cfPaymentID, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `cfPaymentID`                                                                                                              | *int64*                                                                                                                    | :heavy_check_mark:                                                                                                         | Cashfree payment ID to view the payment details of an order.                                                               |
| `orderID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The order or invoice ID for which you want to view the payment details.                                                    |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


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
    orderID := "est"
    xAPIVersion := "mollitia"
    xRequestID := "laborum"

    ctx := context.Background()
    res, err := s.Payments.GetforOrder(ctx, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `orderID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Order or the invoice ID for which you want to view the payment details.                                                    |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


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
    xAPIVersion := "dolores"
    orderPayRequest := &shared.OrderPayRequest{
        OfferID: cashfreego.String("faa6cc05-d1e2-401c-b0cf-0c9db3ff0f0b"),
        PaymentMethod: shared.OrderPayRequestPaymentMethod{},
        PaymentSessionID: "session__CvcEmNKDkmERQrxnx39ibhJ3Ii034pjc8ZVxf3qcgEXCWlgDDlHRgz2XYZCqpajDQSXMMtCusPgOIxYP2LZx0-05p39gC2Vgmq1RAj--gcn",
        SaveInstrument: cashfreego.Bool(false),
    }
    xRequestID := "dolorem"

    ctx := context.Background()
    res, err := s.Payments.PayOrder(ctx, xAPIVersion, orderPayRequest, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OrderPayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `orderPayRequest`                                                                                                          | [*shared.OrderPayRequest](../../models/shared/orderpayrequest.md)                                                          | :heavy_minus_sign:                                                                                                         | Request body to create a transaction at cashfree using `payment_session_id`                                                |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


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
    res, err := s.Payments.PreauthorizeOrder(ctx, operations.CapturePreauthorizationRequest{
        AuthorizationRequest: &shared.AuthorizationRequest{
            Action: shared.AuthorizationRequestActionCapture.ToPointer(),
            Amount: cashfreego.Float64(1289.26),
        },
        OrderID: "nobis",
        XAPIVersion: "enim",
        XIdempotencyKey: cashfreego.String("omnis"),
        XRequestID: cashfreego.String("nemo"),
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
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


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
    paymentID := "minima"
    xAPIVersion := "excepturi"
    otpRequest := &shared.OTPRequest{
        Action: shared.OTPRequestActionSubmitOtp,
        Otp: "iure",
    }
    xRequestID := "culpa"

    ctx := context.Background()
    res, err := s.Payments.Submit(ctx, paymentID, xAPIVersion, otpRequest, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OTPResponseEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `paymentID`                                                                                                                | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The order or invoice ID for which you want to view the payment details.                                                    |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `otpRequest`                                                                                                               | [*shared.OTPRequest](../../models/shared/otprequest.md)                                                                    | :heavy_minus_sign:                                                                                                         | Request body to submit/resend headless OTP. To use this API make sure you have headless OTP enabled for your account       |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.SubmitOTPRequestResponse](../../models/operations/submitotprequestresponse.md), error**

