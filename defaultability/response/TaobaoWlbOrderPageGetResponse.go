package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbOrderPageGetResponse struct {

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
    OrderList  []domain.TaobaoWlbOrderPageGetWlbOrder `json:"order_list,omitempty" `
    /*
        总条数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
