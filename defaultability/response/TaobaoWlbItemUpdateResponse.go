package response

import (
)

type TaobaoWlbItemUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        修改时间
    */
    GmtModified  bool `json:"gmt_modified,omitempty" `
}
