package response

import (
    "topsdk/ability648/domain"
)

type TaobaoItemSkuAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        sku
    */
    Sku  domain.TaobaoItemSkuAddSku `json:"sku,omitempty" `
}
