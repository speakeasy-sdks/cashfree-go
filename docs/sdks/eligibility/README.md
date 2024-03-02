# Eligibility
(*Eligibility*)

### Available Operations

* [GetAllOffers](#getalloffers) - Get eligible Offers
* [GetCardlessEMI](#getcardlessemi) - Get eligible Cardless EMI
* [GetPaylaterMethods](#getpaylatermethods) - Get eligible Paylater

## GetAllOffers

Use this API to get eligible offers for an order or amount.

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


    var xAPIVersion string = "2022-09-01"

    eligibilityOffersRequest := &shared.EligibilityOffersRequest{
        Queries: shared.OfferQueries{
            Amount: cashfreego.Float64(100),
            OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
        },
    }

    var xRequestID *string = cashfreego.String("<value>")

    ctx := context.Background()
    res, err := s.Eligibility.GetAllOffers(ctx, xAPIVersion, eligibilityOffersRequest, xRequestID)
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
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `eligibilityOffersRequest`                                                                                                 | [*shared.EligibilityOffersRequest](../../pkg/models/shared/eligibilityoffersrequest.md)                                    | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for offers                                                                           |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetEligibilityOfferResponse](../../pkg/models/operations/geteligibilityofferresponse.md), error**
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

## GetCardlessEMI

Use this API to get eligible Cardless EMI Payment Methods for a customer on an order.

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


    var xAPIVersion string = "2022-09-01"

    eligibilityCardlessEMIRequest := &shared.EligibilityCardlessEMIRequest{
        Queries: shared.CardlessEMIQueries{
            Amount: cashfreego.Float64(100),
            OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
        },
    }

    var xRequestID *string = cashfreego.String("<value>")

    ctx := context.Background()
    res, err := s.Eligibility.GetCardlessEMI(ctx, xAPIVersion, eligibilityCardlessEMIRequest, xRequestID)
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
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `eligibilityCardlessEMIRequest`                                                                                            | [*shared.EligibilityCardlessEMIRequest](../../pkg/models/shared/eligibilitycardlessemirequest.md)                          | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for cardlessemi                                                                      |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetEligibilityCardlessEMIResponse](../../pkg/models/operations/geteligibilitycardlessemiresponse.md), error**
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

## GetPaylaterMethods

Use this API to get eligible Paylater Payment Methods for a customer on an order.

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


    var xAPIVersion string = "2022-09-01"

    eligibilityCardlessEMIRequest := &shared.EligibilityCardlessEMIRequest{
        Queries: shared.CardlessEMIQueries{
            Amount: cashfreego.Float64(100),
            OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
        },
    }

    var xRequestID *string = cashfreego.String("<value>")

    ctx := context.Background()
    res, err := s.Eligibility.GetPaylaterMethods(ctx, xAPIVersion, eligibilityCardlessEMIRequest, xRequestID)
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
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `eligibilityCardlessEMIRequest`                                                                                            | [*shared.EligibilityCardlessEMIRequest](../../pkg/models/shared/eligibilitycardlessemirequest.md)                          | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for paylater                                                                         |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetEligibilityPaylaterResponse](../../pkg/models/operations/geteligibilitypaylaterresponse.md), error**
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
