package response

import (
)

type TaobaoRpReturngoodsAgreeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
