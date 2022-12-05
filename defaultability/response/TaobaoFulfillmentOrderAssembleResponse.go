package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFulfillmentOrderAssembleResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        调用结果
    */
    Result  bool `json:"result,omitempty" `
    /*
        错误码
    */
    CallErrorCode  string `json:"call_error_code,omitempty" `
    /*
        错误信息描述
    */
    CallErrorMsg  string `json:"call_error_msg,omitempty" `
    /*
        回传结果
    */
    Model  domain.TaobaoFulfillmentOrderAssembleOrderAssembleResponse `json:"model,omitempty" `
}
