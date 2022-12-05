package response

import (
)

type TaobaoPicturePicturesCountResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询的文件总数
    */
    Totals  int64 `json:"totals,omitempty" `
}
