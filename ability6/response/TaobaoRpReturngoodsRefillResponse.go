package response

import (
)

type TaobaoRpReturngoodsRefillResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        验货操作是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
