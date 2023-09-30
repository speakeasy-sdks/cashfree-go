<!-- Start SDK Example Usage -->


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
<!-- End SDK Example Usage -->