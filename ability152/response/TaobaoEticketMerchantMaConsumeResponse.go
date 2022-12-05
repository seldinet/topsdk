package response

import (
    "topsdk/ability152/domain"
)

type TaobaoEticketMerchantMaConsumeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        系统自动生成
    */
    RespBody  domain.TaobaoEticketMerchantMaConsumeConsumeMaCallbackResp `json:"resp_body,omitempty" `
    /*
        子结果码
    */
    RetCode  string `json:"ret_code,omitempty" `
    /*
        子结果信息
    */
    RetMsg  string `json:"ret_msg,omitempty" `
}
