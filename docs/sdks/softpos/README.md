# SoftPOS

## Overview

softPOS' agent and order management system now supported by APIs

### Available Operations

* [TerminalStatus](#terminalstatus) - Get terminal status using phone number
* [CreateTerminals](#createterminals) - Create Terminal

## TerminalStatus

Use this API to view all details of a terminal.

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
    res, err := s.SoftPOS.TerminalStatus(ctx, operations.GetTerminalByMobileNumberRequest{
        TerminalPhoneNo: "dolorum",
        XAPIVersion: "excepturi",
        XRequestID: cashfree.String("tempora"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TerminalDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetTerminalByMobileNumberRequest](../../models/operations/getterminalbymobilenumberrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetTerminalByMobileNumberResponse](../../models/operations/getterminalbymobilenumberresponse.md), error**


## CreateTerminals

Use this API to create new terminals to use softPOS.

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
    res, err := s.SoftPOS.CreateTerminals(ctx, operations.CreateTerminalsRequest{
        CreateTerminalRequest: &shared.CreateTerminalRequest{
            TerminalID: cashfree.String("facilis"),
            TerminalName: "tempore",
            TerminalPhoneNo: "labore",
        },
        XAPIVersion: "delectus",
        XIdempotencyKey: cashfree.String("eum"),
        XRequestID: cashfree.String("non"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TerminalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateTerminalsRequest](../../models/operations/createterminalsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateTerminalsResponse](../../models/operations/createterminalsresponse.md), error**

