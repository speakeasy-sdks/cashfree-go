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


    var cfPaymentID int64 = 158773

    var orderID string = "<value>"

    var xAPIVersion string = "2022-09-01"

    var xRequestID *string = cashfreego.String("<value>")

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
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetPaymentbyIDResponse](../../pkg/models/operations/getpaymentbyidresponse.md), error**
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
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## GetforOrder

Use this API to view all payment details for an order.

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


    var orderID string = "<value>"

    var xAPIVersion string = "2022-09-01"

    var xRequestID *string = cashfreego.String("<value>")

    ctx := context.Background()
    res, err := s.Payments.GetforOrder(ctx, orderID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
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
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetPaymentsforOrderResponse](../../pkg/models/operations/getpaymentsfororderresponse.md), error**
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
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

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
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"context"
	"log"
)

func main() {
    s := cashfreego.New()


    var xAPIVersion string = "2022-09-01"

    orderPayRequest := &shared.OrderPayRequest{
        OfferID: cashfreego.String("faa6cc05-d1e2-401c-b0cf-0c9db3ff0f0b"),
        PaymentMethod: shared.CreateOrderPayRequestPaymentMethodPaylaterPaymentMethod(
                shared.PaylaterPaymentMethod{
                    Paylater: shared.Paylater{
                        Channel: cashfreego.String("link"),
                        Phone: cashfreego.String("7789112345"),
                        Provider: shared.PaylaterProviderKotak.ToPointer(),
                    },
                },
        ),
        PaymentSessionID: "session__CvcEmNKDkmERQrxnx39ibhJ3Ii034pjc8ZVxf3qcgEXCWlgDDlHRgz2XYZCqpajDQSXMMtCusPgOIxYP2LZx0-05p39gC2Vgmq1RAj--gcn",
    }

    var xRequestID *string = cashfreego.String("<value>")

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
| `orderPayRequest`                                                                                                          | [*shared.OrderPayRequest](../../pkg/models/shared/orderpayrequest.md)                                                      | :heavy_minus_sign:                                                                                                         | Request body to create a transaction at cashfree using `payment_session_id`                                                |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.OrderPayResponse](../../pkg/models/operations/orderpayresponse.md), error**
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
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## PreauthorizeOrder

Use this API to capture or void a preauthorized payment

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
    res, err := s.Payments.PreauthorizeOrder(ctx, operations.CapturePreauthorizationRequest{
        AuthorizationRequest: &shared.AuthorizationRequest{
            Action: shared.AuthorizationRequestActionCapture.ToPointer(),
            Amount: cashfreego.Float64(100),
        },
        OrderID: "<value>",
        XAPIVersion: "2022-09-01",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CapturePreauthorizationRequest](../../pkg/models/operations/capturepreauthorizationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.CapturePreauthorizationResponse](../../pkg/models/operations/capturepreauthorizationresponse.md), error**
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
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## Submit

If you accept OTP on your own page, you can use the below API to send OTP to Cashfree.

### Example Usage

```go
package main

import(
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"context"
	"log"
)

func main() {
    s := cashfreego.New()


    var paymentID string = "<value>"

    var xAPIVersion string = "2022-09-01"

    otpRequest := &shared.OTPRequest{
        Action: shared.OTPRequestActionResendOtp,
        Otp: "<value>",
    }

    var xRequestID *string = cashfreego.String("<value>")

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
| `otpRequest`                                                                                                               | [*shared.OTPRequest](../../pkg/models/shared/otprequest.md)                                                                | :heavy_minus_sign:                                                                                                         | Request body to submit/resend headless OTP. To use this API make sure you have headless OTP enabled for your account       |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.SubmitOTPRequestResponse](../../pkg/models/operations/submitotprequestresponse.md), error**
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
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |
