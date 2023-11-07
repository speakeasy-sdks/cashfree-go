# AuthorizationRequest


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Action`                                                                                | [*shared.AuthorizationRequestAction](../../models/shared/authorizationrequestaction.md) | :heavy_minus_sign:                                                                      | Type of authorization to run. Can be one of 'CAPTURE' , 'VOID'                          |
| `Amount`                                                                                | **float64*                                                                              | :heavy_minus_sign:                                                                      | The amount if you are running a 'CAPTURE'                                               |