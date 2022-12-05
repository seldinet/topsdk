package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbSubscriptionQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        卖家定购的服务列表
    */
    SellerSubscriptionList  []domain.TaobaoWlbSubscriptionQuerySellerSubscriptionList `json:"seller_subscription_list,omitempty" `
}
