package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemPropimgDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        属性图片结构
    */
    PropImg  domain.TaobaoItemPropimgDeletePropImg `json:"prop_img,omitempty" `
}
