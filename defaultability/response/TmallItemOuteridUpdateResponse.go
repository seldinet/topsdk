package response

import (
)

type TmallItemOuteridUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商家编码更新结果
    */
    OuteridUpdateResult  string `json:"outerid_update_result,omitempty" `
}
