package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoCooperationGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        合作分销关系
    */
    Cooperations  []domain.TaobaoFenxiaoCooperationGetCooperation `json:"cooperations,omitempty" `
}
