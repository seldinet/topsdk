package response

import (
    "topsdk/ability149/domain"
)

type TaobaoInventoryWarehouseQueryResponse struct {

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
    Result  domain.TaobaoInventoryWarehouseQueryPaginationResult `json:"result,omitempty" `
}
