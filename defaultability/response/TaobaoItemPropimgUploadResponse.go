package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemPropimgUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        PropImg属性图片结构
    */
    PropImg  domain.TaobaoItemPropimgUploadPropImg `json:"prop_img,omitempty" `
}
