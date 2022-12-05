package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemImgUploadResponse struct {

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
    ItemImg  domain.TaobaoItemImgUploadItemImg `json:"item_img,omitempty" `
}
