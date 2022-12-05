package response

import (
    "topsdk/ability138/domain"
)

type TaobaoItemUpdateDelistingTmallResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回商品更新信息：返回的结果是:num_iid和modified
    */
    Item  domain.TaobaoItemUpdateDelistingTmallItem `json:"item,omitempty" `
}
