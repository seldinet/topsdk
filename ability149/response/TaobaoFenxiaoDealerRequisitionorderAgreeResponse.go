package response

import (
)

type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作是否成功。true：成功；false：失败。
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
