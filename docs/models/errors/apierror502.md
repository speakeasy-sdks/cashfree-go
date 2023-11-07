# APIError502


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Code`                                                                             | **string*                                                                          | :heavy_minus_sign:                                                                 | `bank_processing_failure` will be returned here to denote failure at bank.<br/>    |
| `Message`                                                                          | **string*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Type`                                                                             | [*sdkerrors.SchemasAPIError502Type](../../models/errors/schemasapierror502type.md) | :heavy_minus_sign:                                                                 | api_error                                                                          |