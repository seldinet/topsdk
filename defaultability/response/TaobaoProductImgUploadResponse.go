package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoProductImgUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回产品图片结构中的：url,id,created,modified
    */
    ProductImg  domain.TaobaoProductImgUploadProductImg `json:"product_img,omitempty" `
}
