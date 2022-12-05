package response

import (
    "topsdk/ability149/domain"
)

type TaobaoInventoryInitialResponse struct {

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
    TipInfos  []domain.TaobaoInventoryInitialTipInfo `json:"tip_infos,omitempty" `
}
