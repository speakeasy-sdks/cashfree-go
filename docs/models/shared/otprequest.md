# OTPRequest

Request body to submit/resend headless OTP. To use this API make sure you have headless OTP enabled for your account


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Action`                                                               | [OTPRequestAction](../../models/shared/otprequestaction.md)            | :heavy_check_mark:                                                     | The action for this workflow. Could be either SUBMIT_OTP or RESEND_OTP |
| `Otp`                                                                  | *string*                                                               | :heavy_check_mark:                                                     | OTP to be submitted                                                    |