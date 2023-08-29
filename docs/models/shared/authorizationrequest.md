# AuthorizationRequest

Request body to capture or void a transaction


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Action`                                                                         | [*AuthorizationRequestAction](../../models/shared/authorizationrequestaction.md) | :heavy_minus_sign:                                                               | Type of authorization to run. Can be one of 'CAPTURE' , 'VOID'                   |
| `Amount`                                                                         | **float64*                                                                       | :heavy_minus_sign:                                                               | The amount if you are running a 'CAPTURE'                                        |