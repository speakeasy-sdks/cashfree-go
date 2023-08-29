# GetRefundResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `Headers`                                                     | map[string][]*string*                                         | :heavy_minus_sign:                                            | N/A                                                           |
| `RefundsEntity`                                               | [*shared.RefundsEntity](../../models/shared/refundsentity.md) | :heavy_minus_sign:                                            | OK                                                            |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |