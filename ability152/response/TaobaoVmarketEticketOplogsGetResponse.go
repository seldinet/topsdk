package response

import (
    "topsdk/ability152/domain"
)

type TaobaoVmarketEticketOplogsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作日志列表
    */
    EticketOpLogs  []domain.TaobaoVmarketEticketOplogsGetEticketOpLog `json:"eticket_op_logs,omitempty" `
    /*
        符合条件的记录总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
