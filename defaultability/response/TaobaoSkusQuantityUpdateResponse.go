package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSkusQuantityUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    */
    Item  domain.TaobaoSkusQuantityUpdateItem `json:"item,omitempty" `
}
