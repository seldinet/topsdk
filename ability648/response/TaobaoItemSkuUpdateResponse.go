package response

import (
    "topsdk/ability648/domain"
)

type TaobaoItemSkuUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品Sku
    */
    Sku  domain.TaobaoItemSkuUpdateSku `json:"sku,omitempty" `
}
