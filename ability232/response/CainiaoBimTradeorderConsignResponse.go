package response

import (
)

type CainiaoBimTradeorderConsignResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        菜鸟仓库作业单据号
    */
    StoreOrderCode  string `json:"store_order_code,omitempty" `
    /*
        菜鸟物流订单号,格式为LP开头
    */
    LgOrderCode  string `json:"lg_order_code,omitempty" `
}
