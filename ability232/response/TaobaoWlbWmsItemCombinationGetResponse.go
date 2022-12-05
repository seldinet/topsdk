package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsItemCombinationGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回结果
    */
    Result  domain.TaobaoWlbWmsItemCombinationGetResult `json:"result,omitempty" `
}
