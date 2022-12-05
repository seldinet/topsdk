package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoDealerRequisitionorderGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        采购申请/经销采购单结果列表
    */
    DealerOrders  []domain.TaobaoFenxiaoDealerRequisitionorderGetDealerOrder `json:"dealer_orders,omitempty" `
    /*
        按查询条件查到的记录总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
