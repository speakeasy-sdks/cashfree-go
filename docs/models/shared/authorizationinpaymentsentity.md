# AuthorizationInPaymentsEntity

The authorization details are present for payments which go through the preauthorization workflow. Or else this parameter will be null.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Action`                                                              | [*shared.Action](../../models/shared/action.md)                       | :heavy_minus_sign:                                                    | One of CAPTURE or VOID                                                |
| `ActionReference`                                                     | **string*                                                             | :heavy_minus_sign:                                                    | CAPTURE or VOID reference number based on action                      |
| `ActionTime`                                                          | **string*                                                             | :heavy_minus_sign:                                                    | Time of action (CAPTURE or VOID)                                      |
| `ApproveBy`                                                           | **string*                                                             | :heavy_minus_sign:                                                    | Approve by time as passed in the authorization request (only for UPI) |
| `CapturedAmount`                                                      | **float64*                                                            | :heavy_minus_sign:                                                    | The captured amount for this authorization request                    |
| `EndTime`                                                             | **string*                                                             | :heavy_minus_sign:                                                    | End time of this authorization hold (only for UPI)                    |
| `StartTime`                                                           | **string*                                                             | :heavy_minus_sign:                                                    | Start time of this authorization hold (only for UPI)                  |
| `Status`                                                              | [*shared.Status](../../models/shared/status.md)                       | :heavy_minus_sign:                                                    | One of SUCCESS or PENDING                                             |