package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemJointImgResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品图片信息
    */
    ItemImg  domain.TaobaoItemJointImgItemImg `json:"item_img,omitempty" `
}
