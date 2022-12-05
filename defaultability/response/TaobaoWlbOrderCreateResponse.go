package response

import (
    "topsdk/util"
)

type TaobaoWlbOrderCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        物流宝订单创建成功后，返回物流宝的订单编号；如果订单创建失败，该字段为空。
    */
    OrderCode  string `json:"order_code,omitempty" `
    /*
        订单创建时间
    */
    CreateTime  util.LocalTime `json:"create_time,omitempty" `
}
