package response

import (
    "topsdk/ability149/domain"
)

type TaobaoInventoryWarehouseManageResponse struct {

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
    Result  domain.TaobaoInventoryWarehouseManageResult `json:"result,omitempty" `
}
