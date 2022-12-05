package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemBarcodeUpdateResponse struct {

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
    Item  domain.TaobaoItemBarcodeUpdateItem `json:"item,omitempty" `
}
