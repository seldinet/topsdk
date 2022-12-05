package response

import (
    "topsdk/ability149/domain"
)

type TaobaoInventoryInitialItemResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        提示信息
    */
    TipInfos  []domain.TaobaoInventoryInitialItemTipInfo `json:"tip_infos,omitempty" `
}
