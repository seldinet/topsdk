package response

import (
)

type TaobaoFenxiaoProductcatAddResponse struct {

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
    /*
        产品线ID
    */
    ProductLineId  int64 `json:"product_line_id,omitempty" `
}
