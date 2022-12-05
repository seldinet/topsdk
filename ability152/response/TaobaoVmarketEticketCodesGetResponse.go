package response

import (
    "topsdk/ability152/domain"
)

type TaobaoVmarketEticketCodesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        电子凭证码列表
    */
    EticketCodes  []domain.TaobaoVmarketEticketCodesGetEticketCode `json:"eticket_codes,omitempty" `
    /*
        记录总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
