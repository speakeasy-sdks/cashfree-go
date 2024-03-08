<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"log"
)

func main() {
	s := cashfreego.New(
		cashfreego.WithSecurity(shared.Security{
			Option1: &shared.SecurityOption1{
				XClientID:     "<YOUR_API_KEY_HERE>",
				XClientSecret: "<YOUR_API_KEY_HERE>",
			},
		}),
	)

	var customerID string = "<value>"

	var instrumentID string = "<value>"

	var xAPIVersion string = "2022-09-01"

	var xRequestID *string = cashfreego.String("<value>")

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSavedInstrument(ctx, customerID, instrumentID, xAPIVersion, xRequestID)
	if err != nil {
		log.Fatal(err)
	}
	if res.FetchAllSavedInstruments != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->