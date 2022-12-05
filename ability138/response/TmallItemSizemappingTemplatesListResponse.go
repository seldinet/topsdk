package response

import (
    "topsdk/ability138/domain"
)

type TmallItemSizemappingTemplatesListResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        尺码表模板列表
    */
    SizeMappingTemplates  []domain.TmallItemSizemappingTemplatesListSizeMappingTemplate `json:"size_mapping_templates,omitempty" `
}
