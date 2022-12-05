package response

import (
)

type TaobaoFenxiaoOrderConfirmPaidResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        确认结果成功与否
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
