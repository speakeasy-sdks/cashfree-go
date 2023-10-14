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

    createOfferBackendRequest := &shared.CreateOfferBackendRequest{
        OfferDetails: shared.OfferDetails{
            CashbackDetails: &shared.CashbackDetails{
                CashbackType: shared.CashbackDetailsCashbackTypePercentage,
                CashbackValue: "Money blue shred",
                MaxCashbackAmount: "technology East",
            },
            DiscountDetails: &shared.DiscountDetails{
                DiscountType: shared.DiscountDetailsDiscountTypeFlat,
                DiscountValue: "Northwest",
                MaxDiscountAmount: "SUV quantify Polestar",
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
            OfferTncType: shared.OfferTncOfferTncTypeText,
            OfferTncValue: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
        },
        OfferValidations: shared.OfferValidations{
            MaxAllowed: 10,
            MinAmount: cashfreego.Float64(1),
            PaymentMethod: shared.CreateOfferValidationsPaymentMethodOfferWallet(
                    shared.OfferWallet{
                        App: &shared.OfferWalletWalletOffer{
                            Provider: cashfreego.String("paytm"),
                        },
                    },
            ),
        },
    }

    var xRequestID *string = "physical"

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
| `createOfferBackendRequest`                                                                                                | [*shared.CreateOfferBackendRequest](../../models/shared/createofferbackendrequest.md)                                      | :heavy_minus_sign:                                                                                                         | Request body to create an offer at Cashfree                                                                                |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.CreateOfferResponse](../../models/operations/createofferresponse.md), error**


## Get

Use this API to get offer by offer_id

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


    var offerID string = "female"

    var xAPIVersion string = "program"

    var xRequestID *string = "transmit"

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
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetOfferResponse](../../models/operations/getofferresponse.md), error**

