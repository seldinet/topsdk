package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsInventoryProfitlossGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        损益信息
    */
    ProfitLossInfo  domain.TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info,omitempty" `
}
