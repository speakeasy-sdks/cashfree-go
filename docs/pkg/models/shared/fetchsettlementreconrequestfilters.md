# FetchSettlementReconRequestFilters

Specify either the Settlement ID, Settlement UTR, or start date and end date to fetch the settlement details.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `CfSettlementIds`                                                                 | []*int64*                                                                         | :heavy_minus_sign:                                                                | List of settlement IDs for which you want the settlement reconciliation details.  |
| `EndDate`                                                                         | **string*                                                                         | :heavy_minus_sign:                                                                | Specify the end date till when you want the settlement reconciliation details.    |
| `SettlementUtrs`                                                                  | []*string*                                                                        | :heavy_minus_sign:                                                                | List of settlement UTRs for which you want the settlement reconciliation details. |
| `StartDate`                                                                       | **string*                                                                         | :heavy_minus_sign:                                                                | Specify the start date from when you want the settlement reconciliation details.  |