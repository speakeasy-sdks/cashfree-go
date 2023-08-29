# CreateOfferResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ContentType`                                             | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `Headers`                                                 | map[string][]*string*                                     | :heavy_minus_sign:                                        | N/A                                                       |
| `OfferEntity`                                             | [*shared.OfferEntity](../../models/shared/offerentity.md) | :heavy_minus_sign:                                        | OK                                                        |
| `StatusCode`                                              | *int*                                                     | :heavy_check_mark:                                        | N/A                                                       |
| `RawResponse`                                             | [*http.Response](https://pkg.go.dev/net/http#Response)    | :heavy_minus_sign:                                        | N/A                                                       |