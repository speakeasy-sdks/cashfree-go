# OTPResponseEntityAuthenticateStatus

Status of the is action. Will be either failed or successful. If the action is successful, you should still call the authorization status to verify the final payment status.


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `OTPResponseEntityAuthenticateStatusFailed`  | FAILED                                       |
| `OTPResponseEntityAuthenticateStatusSuccess` | SUCCESS                                      |