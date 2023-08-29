# Eligibility

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
    res, err := s.Eligibility.GetAllOffers(ctx, operations.GetEligibilityOfferRequest{
        EligibilityOffersRequest: &shared.EligibilityOffersRequest{
            Filters: &shared.OfferFilters{
                OfferType: []shared.OfferType{
                    shared.OfferTypeNoCostEmi,
                    shared.OfferTypeCashback,
                    shared.OfferTypeDiscountAndCashback,
                },
            },
            Queries: shared.OfferQueries{
                Amount: cashfree.Float64(100),
                OrderID: cashfree.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: "deserunt",
        XRequestID: cashfree.String("suscipit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleOffersEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetEligibilityOfferRequest](../../models/operations/geteligibilityofferrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetEligibilityOfferResponse](../../models/operations/geteligibilityofferresponse.md), error**


## GetCardlessEMI

Use this API to get eligible Cardless EMI Payment Methods for a customer on an order.

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
    res, err := s.Eligibility.GetCardlessEMI(ctx, operations.GetEligibilityCardlessEMIRequest{
        EligibilityCardlessEMIRequest: &shared.EligibilityCardlessEMIRequest{
            Queries: shared.CardlessEMIQueries{
                Amount: cashfree.Float64(100),
                CustomerDetails: &shared.CustomerDetailsCardlessEMI{
                    CustomerPhone: "9898989898",
                },
                OrderID: cashfree.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: "iure",
        XRequestID: cashfree.String("magnam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleCardlessEMIEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetEligibilityCardlessEMIRequest](../../models/operations/geteligibilitycardlessemirequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetEligibilityCardlessEMIResponse](../../models/operations/geteligibilitycardlessemiresponse.md), error**


## GetPaylaterMethods

Use this API to get eligible Paylater Payment Methods for a customer on an order.

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
    res, err := s.Eligibility.GetPaylaterMethods(ctx, operations.GetEligibilityPaylaterRequest{
        EligibilityCardlessEMIRequest: &shared.EligibilityCardlessEMIRequest{
            Queries: shared.CardlessEMIQueries{
                Amount: cashfree.Float64(100),
                CustomerDetails: &shared.CustomerDetailsCardlessEMI{
                    CustomerPhone: "9898989898",
                },
                OrderID: cashfree.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: "debitis",
        XRequestID: cashfree.String("ipsa"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligiblePaylaters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetEligibilityPaylaterRequest](../../models/operations/geteligibilitypaylaterrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetEligibilityPaylaterResponse](../../models/operations/geteligibilitypaylaterresponse.md), error**

