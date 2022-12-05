package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbOrderitemPageGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        系统自动生成
    */
    OrderItemList  []domain.TaobaoWlbOrderitemPageGetWlbOrderItem `json:"order_item_list,omitempty" `
    /*
        总数量
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
