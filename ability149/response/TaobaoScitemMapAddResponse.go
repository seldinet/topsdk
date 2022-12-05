package response

import (
)

type TaobaoScitemMapAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口调用返回结果信息：商家编码
    */
    OuterCode  string `json:"outer_code,omitempty" `
}
