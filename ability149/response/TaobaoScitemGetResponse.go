package response

import (
    "topsdk/ability149/domain"
)

type TaobaoScitemGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        后端商品
    */
    ScItem  domain.TaobaoScitemGetScItem `json:"sc_item,omitempty" `
}
