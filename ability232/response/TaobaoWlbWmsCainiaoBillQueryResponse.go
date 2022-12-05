package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsCainiaoBillQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        订单列表信息
    */
    OrderInfoList  []domain.TaobaoWlbWmsCainiaoBillQueryCainiaoBillQueryOrderinfolist `json:"order_info_list,omitempty" `
    /*
        总条数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
