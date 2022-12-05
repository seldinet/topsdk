package response

import (
)

type TaobaoVmarketEticketTimeExpandResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        0:失败；1:成功
    */
    RetCode  int64 `json:"ret_code,omitempty" `
}
