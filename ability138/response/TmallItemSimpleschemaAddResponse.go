package response

import (
    "topsdk/util"
)

type TmallItemSimpleschemaAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        发布成功后返回商品ID
    */
    Result  string `json:"result,omitempty" `
    /*
        商品最后发布时间。
    */
    GmtModified  util.LocalTime `json:"gmt_modified,omitempty" `
}
