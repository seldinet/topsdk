package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemSkuDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        Sku结构
    */
    Sku  domain.TaobaoItemSkuDeleteSku `json:"sku,omitempty" `
}
