# VendorSplit

Use to split order when cashfree's Easy Split is enabled for your account.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Amount`                                                           | **float64*                                                         | :heavy_minus_sign:                                                 | Amount which will be associated with this vendor                   |
| `Percentage`                                                       | **float64*                                                         | :heavy_minus_sign:                                                 | Percentage of order amount which shall get added to vendor account |
| `VendorID`                                                         | **string*                                                          | :heavy_minus_sign:                                                 | Vendor id created in Cashfree system                               |