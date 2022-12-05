package response

import (
    "topsdk/ability149/domain"
)

type TaobaoScitemQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        List<ScItemDO>
    */
    ScItemList  []domain.TaobaoScitemQueryScItem `json:"sc_item_list,omitempty" `
    /*
        商品条数
    */
    TotalPage  int64 `json:"total_page,omitempty" `
    /*
        分页
    */
    QueryPagination  domain.TaobaoScitemQueryQueryPagination `json:"query_pagination,omitempty" `
}
