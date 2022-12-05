package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemImgDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品图片结构
    */
    ItemImg  domain.TaobaoItemImgDeleteItemImg `json:"item_img,omitempty" `
}
