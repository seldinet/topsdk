package response

import (
    "topsdk/ability149/domain"
)

type TmallChannelProductsQueryResponse struct {

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
    Result  domain.TmallChannelProductsQueryPageResultDto `json:"result,omitempty" `
}
