package response

import (
)

type TmallItemShiptimeUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        被修改的商品ID
    */
    Result  string `json:"result,omitempty" `
}
