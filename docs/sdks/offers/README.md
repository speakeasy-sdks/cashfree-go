# Offers
(*Offers*)

## Overview

Collection of apis to get offers applicable for an order

### Available Operations

* [Create](#create) - Create Offer
* [Get](#get) - Get Offer by ID

## Create

Use this API to create offers with Cashfree from your backend

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


    var xAPIVersion string = "string"

    createOfferBackendRequest := &shared.CreateOfferBackendRequest{
        OfferDetails: shared.OfferDetails{
            CashbackDetails: &shared.CashbackDetails{
                CashbackType: shared.CashbackTypePercentage,
                CashbackValue: "20",
                MaxCashbackAmount: "150",
            },
            DiscountDetails: &shared.DiscountDetails{
                DiscountType: shared.DiscountTypeFlat,
                DiscountValue: "10",
                MaxDiscountAmount: "10",
            },
            OfferType: shared.OfferDetailsOfferTypeDiscountAndCashback,
        },
        OfferMeta: shared.OfferMeta{
            OfferCode: "CFTESTOFFER",
            OfferDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
            OfferEndTime: "2023-03-29T08:09:51Z",
            OfferStartTime: "2023-03-21T08:09:51Z",
            OfferTitle: "Test Offer",
        },
        OfferTnc: shared.OfferTnc{
            OfferTncType: shared.OfferTncTypeText,
            OfferTncValue: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
        },
        OfferValidations: shared.OfferValidations{
            MaxAllowed: 10,
            MinAmount: cashfreego.Float64(1),
            PaymentMethod: shared.CreatePaymentMethodOfferWallet(
                    shared.OfferWallet{},
            ),
        },
    }

    var xRequestID *string = cashfreego.String("string")

    ctx := context.Background()
    res, err := s.Offers.Create(ctx, xAPIVersion, createOfferBackendRequest, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OfferEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `createOfferBackendRequest`                                                                                                | [*shared.CreateOfferBackendRequest](../../pkg/models/shared/createofferbackendrequest.md)                                  | :heavy_minus_sign:                                                                                                         | Request body to create an offer at Cashfree                                                                                |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.CreateOfferResponse](../../pkg/models/operations/createofferresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## Get

Use this API to get offer by offer_id

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


    var offerID string = "string"

    var xAPIVersion string = "string"

    var xRequestID *string = cashfreego.String("string")

    ctx := context.Background()
    res, err := s.Offers.Get(ctx, offerID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OfferEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `offerID`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The offer ID for which you want to view the offer details.                                                                 |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetOfferResponse](../../pkg/models/operations/getofferresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |
