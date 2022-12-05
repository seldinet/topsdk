package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbItemGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品信息
    */
    Item  domain.TaobaoWlbItemGetWlbItem `json:"item,omitempty" `
}
