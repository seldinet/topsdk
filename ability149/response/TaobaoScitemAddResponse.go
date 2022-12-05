package response

import (
    "topsdk/ability149/domain"
)

type TaobaoScitemAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        后台商品信息
    */
    ScItem  domain.TaobaoScitemAddScItem `json:"sc_item,omitempty" `
}
