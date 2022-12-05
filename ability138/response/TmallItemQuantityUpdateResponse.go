package response

import (
)

type TmallItemQuantityUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        库存更新结果，商品id
    */
    QuantityUpdateResult  string `json:"quantity_update_result,omitempty" `
}
