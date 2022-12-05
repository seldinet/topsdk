package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbWlborderGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        物流宝订单对象
    */
    WlbOrder  domain.TaobaoWlbWlborderGetWlbOrder `json:"wlb_order,omitempty" `
}
