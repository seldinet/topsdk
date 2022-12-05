package response

import (
)

type TaobaoFenxiaoDealerRequisitionorderCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        经销采购申请编号
    */
    DealerOrderId  int64 `json:"dealer_order_id,omitempty" `
}
