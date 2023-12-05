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
				XClientID:     "",
				XClientSecret: "",
			},
		}),
	)

	var customerID string = "string"

	var instrumentID string = "string"

	var xAPIVersion string = "string"

	var xRequestID *string = "string"

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