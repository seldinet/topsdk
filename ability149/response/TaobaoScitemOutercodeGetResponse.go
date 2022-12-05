package response

import (
    "topsdk/ability149/domain"
)

type TaobaoScitemOutercodeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        后台商品
    */
    ScItem  domain.TaobaoScitemOutercodeGetScItem `json:"sc_item,omitempty" `
}
