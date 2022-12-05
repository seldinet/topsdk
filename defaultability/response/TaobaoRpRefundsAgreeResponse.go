package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoRpRefundsAgreeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        信息
    */
    Message  string `json:"message,omitempty" `
    /*
        操作成功
    */
    Succ  bool `json:"succ,omitempty" `
    /*
        退款操作结果列表
    */
    Results  []domain.TaobaoRpRefundsAgreeRefundMappingResult `json:"results,omitempty" `
    /*
        批量退款操作情况，可选值：OP_SUCC（全部成功），SOME_OP_SUCC（部分成功），OP_FAILURE_UE（全部失败）
    */
    MsgCode  string `json:"msg_code,omitempty" `
}
