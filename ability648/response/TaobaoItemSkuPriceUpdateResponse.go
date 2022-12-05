package response

import (
    "topsdk/ability648/domain"
)

type TaobaoItemSkuPriceUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品SKU信息（只包含num_iid和modified）
    */
    Sku  domain.TaobaoItemSkuPriceUpdateSku `json:"sku,omitempty" `
}
