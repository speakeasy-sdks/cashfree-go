# TokenVault

## Overview

Cashfree's token Vault helps you save cards and tokenize them in a PCI complaint manner. We support creation of network tokens which can be used across acquiring banks

### Available Operations

* [DeleteSavedInstrument](#deletesavedinstrument) - Delete Saved Instrument
* [FetchSavedInstrument](#fetchsavedinstrument) - Fetch Single Saved Instrument
* [FetchSavedInstrumentCryptogram](#fetchsavedinstrumentcryptogram) - Fetch cryptogram for saved instrument
* [GetAllSavedInstruments](#getallsavedinstruments) - Fetch All Saved Instruments

## DeleteSavedInstrument

To delete a saved instrument for a customer id and instrument id

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
    customerID := "ipsa"
    instrumentID := "omnis"
    xAPIVersion := "voluptate"
    xRequestID := "cum"

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

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `customerID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The customer_id for which instrument needs to be deleted                                                                   |
| `instrumentID`                                                                                                             | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The instrument_id which needs to be deleted                                                                                |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.DeleteSpecificSavedInstrumentResponse](../../models/operations/deletespecificsavedinstrumentresponse.md), error**


## FetchSavedInstrument

To get specific saved instrument for a customer id and instrument id

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
    customerID := "perferendis"
    instrumentID := "doloremque"
    xAPIVersion := "reprehenderit"
    xRequestID := "ut"

    ctx := context.Background()
    res, err := s.TokenVault.FetchSavedInstrument(ctx, customerID, instrumentID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchAllSavedInstruments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `customerID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The customer_id for which all the saved cards are queried                                                                  |
| `instrumentID`                                                                                                             | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The instrument_id of the saved instrument which needs to be queried                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.FetchSpecificSavedInstrumentResponse](../../models/operations/fetchspecificsavedinstrumentresponse.md), error**


## FetchSavedInstrumentCryptogram

To get the card network token, token expiry and cryptogram for a saved instrument using instrument id

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
    customerID := "maiores"
    instrumentID := "dicta"
    xAPIVersion := "corporis"
    xRequestID := "dolore"

    ctx := context.Background()
    res, err := s.TokenVault.FetchSavedInstrumentCryptogram(ctx, customerID, instrumentID, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Cryptogram != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `customerID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The customer_id of saved instrument                                                                                        |
| `instrumentID`                                                                                                             | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The instrument_id of the saved instrument which needs to be queried                                                        |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.FetchCryptogramResponse](../../models/operations/fetchcryptogramresponse.md), error**


## GetAllSavedInstruments

To get all saved instruments for a customer id

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
    customerID := "iusto"
    instrumentType := "dicta"
    xAPIVersion := "harum"
    xRequestID := "enim"

    ctx := context.Background()
    res, err := s.TokenVault.GetAllSavedInstruments(ctx, customerID, instrumentType, xAPIVersion, xRequestID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchAllSavedInstruments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `customerID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The customer_id for which all the saved cards are queried                                                                  |
| `instrumentType`                                                                                                           | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | type to instrument to query                                                                                                |
| `xAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `xRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.FetchAllSavedInstrumentsResponse](../../models/operations/fetchallsavedinstrumentsresponse.md), error**

