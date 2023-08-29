# GetEligibilityOfferRequest


## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `EligibilityOffersRequest`                                                                                                 | [*shared.EligibilityOffersRequest](../../models/shared/eligibilityoffersrequest.md)                                        | :heavy_minus_sign:                                                                                                         | Request body to check for eligibility for offers                                                                           |
| `XAPIVersion`                                                                                                              | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | API version to be used. Format is in YYYY-MM-DD                                                                            |
| `XRequestID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |