package response

import (
)

type TmallItemSizemappingTemplateDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        尺码表模板ID
    */
    TemplateId  int64 `json:"template_id,omitempty" `
}
