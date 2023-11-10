# Cryptogram


## Fields

| Field                                 | Type                                  | Required                              | Description                           |
| ------------------------------------- | ------------------------------------- | ------------------------------------- | ------------------------------------- |
| `CardDisplay`                         | **string*                             | :heavy_minus_sign:                    | last 4 digits of original card number |
| `CardExpiryMm`                        | **string*                             | :heavy_minus_sign:                    | token pan expiry month                |
| `CardExpiryYy`                        | **string*                             | :heavy_minus_sign:                    | token pan expiry year                 |
| `CardNumber`                          | **string*                             | :heavy_minus_sign:                    | token pan number                      |
| `Cryptogram`                          | **string*                             | :heavy_minus_sign:                    | cryptogram                            |
| `InstrumentID`                        | **string*                             | :heavy_minus_sign:                    | instrument_id of saved instrument     |
| `TokenRequestorID`                    | **string*                             | :heavy_minus_sign:                    | TRID issued by card networks          |