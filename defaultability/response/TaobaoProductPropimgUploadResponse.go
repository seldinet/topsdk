package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoProductPropimgUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        支持返回产品属性图片中的：url,id,created,modified
    */
    ProductPropImg  domain.TaobaoProductPropimgUploadProductPropImg `json:"product_prop_img,omitempty" `
}
