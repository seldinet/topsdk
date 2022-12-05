package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoRefundGetResponse struct {

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
    RefundDetail  domain.TaobaoFenxiaoRefundGetTopDpRefundDetailDO `json:"refund_detail,omitempty" `
}
