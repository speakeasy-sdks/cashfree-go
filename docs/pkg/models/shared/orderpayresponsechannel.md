# OrderPayResponseChannel

One of ["link", "collect", "qrcode"]. In an older version we used to support different channels like 'gpay', 'phonepe' etc. However, we now support only the following channels - link, collect and qrcode. To process payments using gpay, you will have to provide channel as 'link' and provider as 'gpay'


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `OrderPayResponseChannelLink`    | link                             |
| `OrderPayResponseChannelCollect` | collect                          |
| `OrderPayResponseChannelQrcode`  | qrcode                           |
| `OrderPayResponseChannelPost`    | post                             |