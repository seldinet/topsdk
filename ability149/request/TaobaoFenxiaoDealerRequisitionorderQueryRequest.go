package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDealerRequisitionorderQueryRequest struct {
    /*
        经销采购单编号。
多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。     */
    DealerOrderIds  *[]int64 `json:"dealer_order_ids" required:"true" `
    /*
        多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表     */
    Fields  *string `json:"fields,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDealerRequisitionorderQueryRequest) SetDealerOrderIds(v []int64) *TaobaoFenxiaoDealerRequisitionorderQueryRequest {
    s.DealerOrderIds = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryRequest) SetFields(v string) *TaobaoFenxiaoDealerRequisitionorderQueryRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoFenxiaoDealerRequisitionorderQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DealerOrderIds != nil) {
        paramMap["dealer_order_ids"] = util.ConvertBasicList(*req.DealerOrderIds)
    }
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    return paramMap
}

func (req *TaobaoFenxiaoDealerRequisitionorderQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}