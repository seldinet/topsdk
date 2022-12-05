package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoTradePrepayOfflineReduceResponse struct {

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
    Result  domain.TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo `json:"result,omitempty" `
}
