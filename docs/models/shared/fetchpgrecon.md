# FetchPGRecon

OK


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Cursor`                                                                   | **string*                                                                  | :heavy_minus_sign:                                                         | Specifies from where the next set of settlement details should be fetched. |
| `Data`                                                                     | [][FetchPGReconData](../../models/shared/fetchpgrecondata.md)              | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Limit`                                                                    | **int64*                                                                   | :heavy_minus_sign:                                                         | Number of settlements you want to fetch in the next iteration.             |