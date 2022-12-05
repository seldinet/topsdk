package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbOrderstatusGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        订单流转信息状态列表
    */
    WlbOrderStatus  []domain.TaobaoWlbOrderstatusGetWlbProcessStatus `json:"wlb_order_status,omitempty" `
}
