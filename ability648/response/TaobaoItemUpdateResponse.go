package response

import (
    "topsdk/ability648/domain"
)

type TaobaoItemUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品结构
    */
    Item  domain.TaobaoItemUpdateItem `json:"item,omitempty" `
}
