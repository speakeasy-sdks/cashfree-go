<div align="center">
    <img src="https://github.com/speakeasy-sdks/cashfree-go/assets/6267663/fcd70572-8452-4083-8a59-32293f2ed8df" width="300">
    <h1>Go SDK</h1>
   <p>Payments infrastructure for India</p>
   <a href="https://docs.cashfree.com/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/cashfree-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/cashfree-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/cashfree-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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

	var xAPIVersion string = "string"

	eligibilityOffersRequest := &shared.EligibilityOffersRequest{
		Filters: &shared.OfferFilters{
			OfferType: []shared.OfferType{
				shared.OfferTypeNoCostEmi,
			},
		},
		Queries: shared.OfferQueries{
			Amount:  cashfreego.Float64(100),
			OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
		},
	}

	var xRequestID *string = "string"

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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Eligibility](docs/sdks/eligibility/README.md)

* [GetAllOffers](docs/sdks/eligibility/README.md#getalloffers) - Get eligible Offers
* [GetCardlessEMI](docs/sdks/eligibility/README.md#getcardlessemi) - Get eligible Cardless EMI
* [GetPaylaterMethods](docs/sdks/eligibility/README.md#getpaylatermethods) - Get eligible Paylater

### [Offers](docs/sdks/offers/README.md)

* [Create](docs/sdks/offers/README.md#create) - Create Offer
* [Get](docs/sdks/offers/README.md#get) - Get Offer by ID

### [Orders](docs/sdks/orders/README.md)

* [Create](docs/sdks/orders/README.md#create) - Create Order
* [Get](docs/sdks/orders/README.md#get) - Get Order

### [PGReconciliation](docs/sdks/pgreconciliation/README.md)

* [Get](docs/sdks/pgreconciliation/README.md#get) - PG Reconciliation

### [PaymentLinks](docs/sdks/paymentlinks/README.md)

* [Cancel](docs/sdks/paymentlinks/README.md#cancel) - Cancel Payment Link
* [Create](docs/sdks/paymentlinks/README.md#create) - Create Payment Link
* [Fetch](docs/sdks/paymentlinks/README.md#fetch) - Fetch Payment Link Details
* [GetOrders](docs/sdks/paymentlinks/README.md#getorders) - Get Orders for a Payment Link

### [Payments](docs/sdks/payments/README.md)

* [Payment](docs/sdks/payments/README.md#payment) - Get Payment by ID
* [GetforOrder](docs/sdks/payments/README.md#getfororder) - Get Payments for an Order
* [PayOrder](docs/sdks/payments/README.md#payorder) - Order Pay
* [PreauthorizeOrder](docs/sdks/payments/README.md#preauthorizeorder) - Preauthorization
* [Submit](docs/sdks/payments/README.md#submit) - Submit or Resend OTP

### [Refunds](docs/sdks/refunds/README.md)

* [Create](docs/sdks/refunds/README.md#create) - Create Refund
* [Get](docs/sdks/refunds/README.md#get) - Get Refund
* [GetAllforOrder](docs/sdks/refunds/README.md#getallfororder) - Get All Refunds for an Order

### [Settlements](docs/sdks/settlements/README.md)

* [Fetch](docs/sdks/settlements/README.md#fetch) - Settlement Reconciliation
* [GetAll](docs/sdks/settlements/README.md#getall) - Get All Settlements
* [GetForOrder](docs/sdks/settlements/README.md#getfororder) - Get Settlements by Order ID

### [TokenVault](docs/sdks/tokenvault/README.md)

* [DeleteSavedInstrument](docs/sdks/tokenvault/README.md#deletesavedinstrument) - Delete Saved Instrument
* [FetchSavedInstrument](docs/sdks/tokenvault/README.md#fetchsavedinstrument) - Fetch Single Saved Instrument
* [FetchSavedInstrumentCryptogram](docs/sdks/tokenvault/README.md#fetchsavedinstrumentcryptogram) - Fetch cryptogram for saved instrument
* [GetAllSavedInstruments](docs/sdks/tokenvault/README.md#getallsavedinstruments) - Fetch All Saved Instruments

### [SoftPOS](docs/sdks/softpos/README.md)

* [TerminalStatus](docs/sdks/softpos/README.md#terminalstatus) - Get terminal status using phone number
* [CreateTerminals](docs/sdks/softpos/README.md#createterminals) - Create Terminal
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


## Example

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

	var xAPIVersion string = "string"

	eligibilityOffersRequest := &shared.EligibilityOffersRequest{
		Filters: &shared.OfferFilters{
			OfferType: []shared.OfferType{
				shared.OfferTypeNoCostEmi,
			},
		},
		Queries: shared.OfferQueries{
			Amount:  cashfreego.Float64(100),
			OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
		},
	}

	var xRequestID *string = "string"

	ctx := context.Background()
	res, err := s.Eligibility.GetAllOffers(ctx, xAPIVersion, eligibilityOffersRequest, xRequestID)
	if err != nil {

		var e *BadRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *AuthenticationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *ApiError404
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *ApiError409
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *IdempotencyError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *RateLimitError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *ApiError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *ApiError502
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://sandbox.cashfree.com/pg` | None |
| 1 | `https://api.cashfree.com/pg` | None |

For example:


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
		cashfreego.WithServerIndex(1),
	)

	var xAPIVersion string = "string"

	eligibilityOffersRequest := &shared.EligibilityOffersRequest{
		Filters: &shared.OfferFilters{
			OfferType: []shared.OfferType{
				shared.OfferTypeNoCostEmi,
			},
		},
		Queries: shared.OfferQueries{
			Amount:  cashfreego.Float64(100),
			OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
		},
	}

	var xRequestID *string = "string"

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


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


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
		cashfreego.WithServerURL("https://sandbox.cashfree.com/pg"),
	)

	var xAPIVersion string = "string"

	eligibilityOffersRequest := &shared.EligibilityOffersRequest{
		Filters: &shared.OfferFilters{
			OfferType: []shared.OfferType{
				shared.OfferTypeNoCostEmi,
			},
		},
		Queries: shared.OfferQueries{
			Amount:  cashfreego.Float64(100),
			OrderID: cashfreego.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
		},
	}

	var xRequestID *string = "string"

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
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
