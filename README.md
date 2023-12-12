<div align="center">
    <img src="https://github.com/speakeasy-sdks/cashfree-go/assets/6267663/fcd70572-8452-4083-8a59-32293f2ed8df" width="300">
    <h1>Go SDK</h1>
   <p>Payments infrastructure for India</p>
   <a href="https://docs.cashfree.com/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/cashfree-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/cashfree-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/cashfree-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [TokenVault](docs/sdks/tokenvault/README.md)

* [DeleteSavedInstrument](docs/sdks/tokenvault/README.md#deletesavedinstrument) - Delete Saved Instrument
* [FetchSavedInstrument](docs/sdks/tokenvault/README.md#fetchsavedinstrument) - Fetch Single Saved Instrument
* [FetchSavedInstrumentCryptogram](docs/sdks/tokenvault/README.md#fetchsavedinstrumentcryptogram) - Fetch cryptogram for saved instrument
* [GetAllSavedInstruments](docs/sdks/tokenvault/README.md#getallsavedinstruments) - Fetch All Saved Instruments

### [Eligibility](docs/sdks/eligibility/README.md)

* [GetAllOffers](docs/sdks/eligibility/README.md#getalloffers) - Get eligible Offers
* [GetCardlessEMI](docs/sdks/eligibility/README.md#getcardlessemi) - Get eligible Cardless EMI
* [GetPaylaterMethods](docs/sdks/eligibility/README.md#getpaylatermethods) - Get eligible Paylater

### [PaymentLinks](docs/sdks/paymentlinks/README.md)

* [Cancel](docs/sdks/paymentlinks/README.md#cancel) - Cancel Payment Link
* [Create](docs/sdks/paymentlinks/README.md#create) - Create Payment Link
* [Fetch](docs/sdks/paymentlinks/README.md#fetch) - Fetch Payment Link Details
* [GetOrders](docs/sdks/paymentlinks/README.md#getorders) - Get Orders for a Payment Link

### [Offers](docs/sdks/offers/README.md)

* [Create](docs/sdks/offers/README.md#create) - Create Offer
* [Get](docs/sdks/offers/README.md#get) - Get Offer by ID

### [Orders](docs/sdks/orders/README.md)

* [Create](docs/sdks/orders/README.md#create) - Create Order
* [Get](docs/sdks/orders/README.md#get) - Get Order

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

### [PGReconciliation](docs/sdks/pgreconciliation/README.md)

* [Get](docs/sdks/pgreconciliation/README.md#get) - PG Reconciliation

### [SoftPOS](docs/sdks/softpos/README.md)

* [TerminalStatus](docs/sdks/softpos/README.md#terminalstatus) - Get terminal status using phone number
* [CreateTerminals](docs/sdks/softpos/README.md#createterminals) - Create Terminal
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/json              |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.APIError404         | 404                           | application/json              |
| sdkerrors.APIError409         | 409                           | application/json              |
| sdkerrors.IdempotencyError    | 422                           | application/json              |
| sdkerrors.RateLimitError      | 429                           | application/json              |
| sdkerrors.APIError            | 500                           | application/json              |
| sdkerrors.APIError502         | 502                           | application/json              |
| sdkerrors.SDKError            | 400-600                       | */*                           |

### Example

```go
package main

import (
	"context"
	"errors"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/sdkerrors"
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

	var customerID string = "string"

	var instrumentID string = "string"

	var xAPIVersion string = "string"

	var xRequestID *string = "string"

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSavedInstrument(ctx, customerID, instrumentID, xAPIVersion, xRequestID)
	if err != nil {

		var e *sdkerrors.BadRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AuthenticationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError404
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError409
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.IdempotencyError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.RateLimitError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError502
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://sandbox.cashfree.com/pg` | None |
| 1 | `https://api.cashfree.com/pg` | None |

#### Example

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
		cashfreego.WithServerIndex(1),
		cashfreego.WithSecurity(shared.Security{
			Option1: &shared.SecurityOption1{
				XClientID:     "<YOUR_API_KEY_HERE>",
				XClientSecret: "<YOUR_API_KEY_HERE>",
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


### Override Server URL Per-Client

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
		cashfreego.WithServerURL("https://sandbox.cashfree.com/pg"),
		cashfreego.WithSecurity(shared.Security{
			Option1: &shared.SecurityOption1{
				XClientID:     "<YOUR_API_KEY_HERE>",
				XClientSecret: "<YOUR_API_KEY_HERE>",
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

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
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"log"
	"pkg/models/operations"
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

	var customerID string = "string"

	var instrumentID string = "string"

	var xAPIVersion string = "string"

	var xRequestID *string = "string"

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSavedInstrument(ctx, customerID, instrumentID, xAPIVersion, xRequestID, operations.WithRetries(
		utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}

	if res.FetchAllSavedInstruments != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	cashfreego "github.com/speakeasy-sdks/cashfree-go"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"log"
)

func main() {
	s := cashfreego.New(
		cashfreego.WithRetryConfig(
			utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		cashfreego.WithSecurity(shared.Security{
			Option1: &shared.SecurityOption1{
				XClientID:     "<YOUR_API_KEY_HERE>",
				XClientSecret: "<YOUR_API_KEY_HERE>",
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
<!-- End Retries [retries] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports multiple security scheme combinations globally. You can choose from one of the alternatives by using the `WithSecurity` option when initializing the SDK client instance. The selected option will be used by default to authenticate with the API for all operations that support it.

#### Option1

All of the following schemes must be satisfied to use the `Option1` alternative:

| Name            | Type            | Scheme          |
| --------------- | --------------- | --------------- |
| `XClientID`     | apiKey          | API key         |
| `XClientSecret` | apiKey          | API key         |

Example:
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

#### Option2

All of the following schemes must be satisfied to use the `Option2` alternative:

| Name             | Type             | Scheme           |
| ---------------- | ---------------- | ---------------- |
| `XClientID`      | apiKey           | API key          |
| `XPartnerAPIKey` | apiKey           | API key          |

Example:
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
			Option2: &shared.SecurityOption2{
				XClientID:      "<YOUR_API_KEY_HERE>",
				XPartnerAPIKey: "<YOUR_API_KEY_HERE>",
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

#### Option3

All of the following schemes must be satisfied to use the `Option3` alternative:

| Name                     | Type                     | Scheme                   |
| ------------------------ | ------------------------ | ------------------------ |
| `XClientID`              | apiKey                   | API key                  |
| `XClientSignatureHeader` | apiKey                   | API key                  |

Example:
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
			Option3: &shared.SecurityOption3{
				XClientID:              "<YOUR_API_KEY_HERE>",
				XClientSignatureHeader: "<YOUR_API_KEY_HERE>",
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

#### Option4

All of the following schemes must be satisfied to use the `Option4` alternative:

| Name                 | Type                 | Scheme               |
| -------------------- | -------------------- | -------------------- |
| `XPartnerAPIKey`     | apiKey               | API key              |
| `XPartnerMerchantID` | apiKey               | API key              |

Example:
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
			Option4: &shared.SecurityOption4{
				XPartnerAPIKey:     "<YOUR_API_KEY_HERE>",
				XPartnerMerchantID: "<YOUR_API_KEY_HERE>",
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
