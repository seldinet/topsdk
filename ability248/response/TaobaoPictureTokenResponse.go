package response

import (
    "topsdk/ability248/domain"
)

type TaobaoPictureTokenResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        响应dto
    */
    Result  domain.TaobaoPictureTokenUsResultDto `json:"result,omitempty" `
}
