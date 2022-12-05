package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemTemplatesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店）
    */
    ItemTemplateList  []domain.TaobaoItemTemplatesGetItemTemplate `json:"item_template_list,omitempty" `
}
