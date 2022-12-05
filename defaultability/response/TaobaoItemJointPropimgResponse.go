package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemJointPropimgResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        属性图片对象信息
    */
    PropImg  domain.TaobaoItemJointPropimgPropImg `json:"prop_img,omitempty" `
}
