# OTPRequest


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Action`                                                                  | [shared.OTPRequestAction](../../../pkg/models/shared/otprequestaction.md) | :heavy_check_mark:                                                        | The action for this workflow. Could be either SUBMIT_OTP or RESEND_OTP    |
| `Otp`                                                                     | *string*                                                                  | :heavy_check_mark:                                                        | OTP to be submitted                                                       |