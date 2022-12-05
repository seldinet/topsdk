package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoProductSkusGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        记录数量
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        sku信息
    */
    Skus  []domain.TaobaoFenxiaoProductSkusGetFenxiaoSku `json:"skus,omitempty" `
}
