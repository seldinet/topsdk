package response

import (
    "topsdk/ability149/domain"
)

type TaobaoRegionSaleQueryResponse struct {

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
    Result  domain.TaobaoRegionSaleQueryBaseResult `json:"result,omitempty" `
}
