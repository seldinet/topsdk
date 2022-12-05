package response

import (
)

type TaobaoVmarketEticketFailsendResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        成功
    */
    RetCode  int64 `json:"ret_code,omitempty" `
}
