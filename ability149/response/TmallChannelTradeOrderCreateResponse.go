package response

import (
)

type TmallChannelTradeOrderCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        采购单号
    */
    MainPurchaseOrderList  []string `json:"main_purchase_order_list,omitempty" `
}
