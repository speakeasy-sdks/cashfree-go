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
    xAPIVersion := "Gasoline"
    eligibilityOffersRequest := &shared.EligibilityOffersRequest{
        Filters: &shared.OfferFilters{
            OfferType: []shared.OfferType{
                shared.OfferTypeNoCostEmi,
            },
        },
        Queries: shared.OfferQueries{
            Amount: cashfreego.Float64(100),
            OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
        },
    }
    xRequestID := "Southwest"

    ctx := context.Background()
    res, err := s.Eligibility.GetAllOffers(ctx, xAPIVersion, eligibilityOffersRequest, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleOffersEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `eligibilityOffersRequest`                                                                                                 | [*shared.EligibilityOffersRequest](../../models/shared/eligibilityoffersrequest.md)                                        | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for offers                                                                           |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


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
    xAPIVersion := "protocol"
    eligibilityCardlessEMIRequest := &shared.EligibilityCardlessEMIRequest{
        Queries: shared.CardlessEMIQueries{
            Amount: cashfreego.Float64(100),
            CustomerDetails: &shared.CustomerDetailsCardlessEMI{
                CustomerPhone: "9898989898",
            },
            OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
        },
    }
    xRequestID := "navigate"

    ctx := context.Background()
    res, err := s.Eligibility.GetCardlessEMI(ctx, xAPIVersion, eligibilityCardlessEMIRequest, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleCardlessEMIEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `eligibilityCardlessEMIRequest`                                                                                            | [*shared.EligibilityCardlessEMIRequest](../../models/shared/eligibilitycardlessemirequest.md)                              | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for cardlessemi                                                                      |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


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
    xAPIVersion := "visualize"
    eligibilityCardlessEMIRequest := &shared.EligibilityCardlessEMIRequest{
        Queries: shared.CardlessEMIQueries{
            Amount: cashfreego.Float64(100),
            CustomerDetails: &shared.CustomerDetailsCardlessEMI{
                CustomerPhone: "9898989898",
            },
            OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
        },
    }
    xRequestID := "concept"

    ctx := context.Background()
    res, err := s.Eligibility.GetPaylaterMethods(ctx, xAPIVersion, eligibilityCardlessEMIRequest, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.EligiblePaylaters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `eligibilityCardlessEMIRequest`                                                                                            | [*shared.EligibilityCardlessEMIRequest](../../models/shared/eligibilitycardlessemirequest.md)                              | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for paylater                                                                         |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetEligibilityPaylaterResponse](../../models/operations/geteligibilitypaylaterresponse.md), error**

