# Offers

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
    res, err := s.Offers.Create(ctx, operations.CreateOfferRequest{
        CreateOfferBackendRequest: &shared.CreateOfferBackendRequest{
            OfferDetails: shared.OfferDetails{
                CashbackDetails: &shared.CashbackDetails{
                    CashbackType: shared.CashbackDetailsCashbackTypePercentage,
                    CashbackValue: "tempora",
                    MaxCashbackAmount: "suscipit",
                },
                DiscountDetails: &shared.DiscountDetails{
                    DiscountType: shared.DiscountDetailsDiscountTypeFlat,
                    DiscountValue: "minus",
                    MaxDiscountAmount: "placeat",
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
                MinAmount: cashfree.Float64(1),
                PaymentMethod: shared.OfferValidationsPaymentMethod{},
            },
        },
        XAPIVersion: "voluptatum",
        XRequestID: cashfree.String("iusto"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OfferEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOfferRequest](../../models/operations/createofferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
    res, err := s.Offers.Get(ctx, operations.GetOfferRequest{
        OfferID: "excepturi",
        XAPIVersion: "nisi",
        XRequestID: cashfree.String("recusandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OfferEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetOfferRequest](../../models/operations/getofferrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetOfferResponse](../../models/operations/getofferresponse.md), error**

