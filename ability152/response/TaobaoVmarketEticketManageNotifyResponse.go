package response

import (
)

type TaobaoVmarketEticketManageNotifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        1:成功
    */
    RetCode  int64 `json:"ret_code,omitempty" `
}
