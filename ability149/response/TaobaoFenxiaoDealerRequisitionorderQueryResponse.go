package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoDealerRequisitionorderQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        经销采购单结果列表
    */
    DealerOrders  []domain.TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder `json:"dealer_orders,omitempty" `
}
