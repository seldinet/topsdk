package response

import (
    "topsdk/ability138/domain"
)

type TmallItemSizemappingTemplateCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        尺码表模板
    */
    SizeMappingTemplate  domain.TmallItemSizemappingTemplateCreateSizeMappingTemplateDo `json:"size_mapping_template,omitempty" `
}
