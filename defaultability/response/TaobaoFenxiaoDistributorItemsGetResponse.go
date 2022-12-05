package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoDistributorItemsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询结果记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        下载记录对象
    */
    Records  []domain.TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord `json:"records,omitempty" `
}
