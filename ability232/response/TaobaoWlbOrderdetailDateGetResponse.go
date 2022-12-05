package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbOrderdetailDateGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        物流宝订单，并且包含订单详情
    */
    OrderDetailList  []domain.TaobaoWlbOrderdetailDateGetWlbOrderDetail `json:"order_detail_list,omitempty" `
    /*
        总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
