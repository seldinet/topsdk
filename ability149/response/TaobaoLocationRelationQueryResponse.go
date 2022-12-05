package response

import (
    "topsdk/ability149/domain"
)

type TaobaoLocationRelationQueryResponse struct {

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
    Result  domain.TaobaoLocationRelationQuerySingleResult `json:"result,omitempty" `
}
