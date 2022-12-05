package response

import (
)

type TaobaoRpRefundReviewResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        success
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
