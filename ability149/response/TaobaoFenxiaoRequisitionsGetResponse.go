package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoRequisitionsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        结果记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        合作申请
    */
    Requisitions  []domain.TaobaoFenxiaoRequisitionsGetRequisition `json:"requisitions,omitempty" `
}
