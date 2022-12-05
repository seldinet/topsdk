package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbTradeorderGetResponse struct {

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
    WlbOrderList  []domain.TaobaoWlbTradeorderGetWlbOrder `json:"wlb_order_list,omitempty" `
}
