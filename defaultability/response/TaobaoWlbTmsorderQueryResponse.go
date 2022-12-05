package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbTmsorderQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        物流订单运单信息列表
    */
    TmsOrderList  []domain.TaobaoWlbTmsorderQueryWlbTmsOrder `json:"tms_order_list,omitempty" `
    /*
        结果总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
