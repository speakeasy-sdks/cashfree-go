<!-- Start SDK Example Usage -->


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
                    shared.OfferTypeDiscountAndCashback,
                    shared.OfferTypeDiscountAndCashback,
                    shared.OfferTypeNoCostEmi,
                },
            },
            Queries: shared.OfferQueries{
                Amount: cashfree.Float64(100),
                OrderID: cashfree.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: "unde",
        XRequestID: cashfree.String("nulla"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleOffersEntities != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->