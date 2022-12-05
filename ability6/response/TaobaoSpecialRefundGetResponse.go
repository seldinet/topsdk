package response

import (
    "topsdk/ability6/domain"
)

type TaobaoSpecialRefundGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        退款详情
    */
    Refund  domain.TaobaoSpecialRefundGetRefund `json:"refund,omitempty" `
}
