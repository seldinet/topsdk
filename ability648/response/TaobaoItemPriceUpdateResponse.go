package response

import (
    "topsdk/ability648/domain"
)

type TaobaoItemPriceUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品结构里的num_iid，modified
    */
    Item  domain.TaobaoItemPriceUpdateItem `json:"item,omitempty" `
}
