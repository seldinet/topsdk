package response

import (
)

type TaobaoScitemMapDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        失效条数
    */
    Module  int64 `json:"module,omitempty" `
}
