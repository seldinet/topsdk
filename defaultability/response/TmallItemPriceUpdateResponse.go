package response

import (
)

type TmallItemPriceUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        价格更新结果
    */
    PriceUpdateResult  string `json:"price_update_result,omitempty" `
}
