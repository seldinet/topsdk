package response

import (
)

type TaobaoFenxiaoProductcatUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
