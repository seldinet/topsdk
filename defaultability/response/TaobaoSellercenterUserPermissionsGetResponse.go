package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSellercenterUserPermissionsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        ๆ้ๅ่กจ
    */
    Permissions  []domain.TaobaoSellercenterUserPermissionsGetPermission `json:"permissions,omitempty" `
}
