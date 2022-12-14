package response

import (
    "topsdk/ability181/domain"
)

type TaobaoItemCatpropsModificationGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        θΏεη»ζ
    */
    Results  []domain.TaobaoItemCatpropsModificationGetPropsModificationResult `json:"results,omitempty" `
}
