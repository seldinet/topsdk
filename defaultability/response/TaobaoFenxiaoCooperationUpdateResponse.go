package response

import (
)

type TaobaoFenxiaoCooperationUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        更新结果成功失败
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
