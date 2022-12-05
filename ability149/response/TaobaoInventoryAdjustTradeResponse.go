package response

import (
    "topsdk/ability149/domain"
)

type TaobaoInventoryAdjustTradeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作返回码
    */
    OperateCode  string `json:"operate_code,omitempty" `
    /*
        提示信息
    */
    TipInfos  []domain.TaobaoInventoryAdjustTradeTipInfo `json:"tip_infos,omitempty" `
}
