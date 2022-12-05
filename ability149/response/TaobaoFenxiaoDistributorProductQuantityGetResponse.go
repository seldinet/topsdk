package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoDistributorProductQuantityGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询结果
    */
    Result  domain.TaobaoFenxiaoDistributorProductQuantityGetResultDTO `json:"result,omitempty" `
}
