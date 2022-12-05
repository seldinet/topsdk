package response

import (
    "topsdk/ability6/domain"
)

type TmallDisputeReceiveGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        result
    */
    Result  domain.TmallDisputeReceiveGetResultSet `json:"result,omitempty" `
}
