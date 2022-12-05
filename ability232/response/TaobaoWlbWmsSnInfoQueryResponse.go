package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsSnInfoQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回
    */
    Result  domain.TaobaoWlbWmsSnInfoQueryResult `json:"result,omitempty" `
}
