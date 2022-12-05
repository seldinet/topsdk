package response

import (
    "topsdk/ability152/domain"
)

type TaobaoVmarketEticketTasksGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        任务列表查询结果信息
    */
    EticketTasks  []domain.TaobaoVmarketEticketTasksGetEticketTask `json:"eticket_tasks,omitempty" `
    /*
        任务列表查询结果的总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
