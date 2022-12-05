package response

import (
)

type TaobaoFenxiaoProductGradepriceUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回操作结果：成功或失败
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
