package response

import (
    "topsdk/ability138/domain"
)

type TaobaoItemUpdateListingTmallResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        上架后返回的商品信息：返回的结果就是:num_iid和modified
    */
    Item  domain.TaobaoItemUpdateListingTmallItem `json:"item,omitempty" `
}
