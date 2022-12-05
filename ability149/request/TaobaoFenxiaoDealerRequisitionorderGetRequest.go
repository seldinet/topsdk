package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDealerRequisitionorderGetRequest struct {
    /*
        采购申请/经销采购单最早修改时间     */
    StartDate  *util.LocalTime `json:"start_date" required:"true" `
    /*
        采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天     */
    EndDate  *util.LocalTime `json:"end_date" required:"true" `
    /*
        页码（大于0的整数。无值或小于1的值按默认值1计） defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） defalutValue��50    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        采购申请/经销采购单状态。
0：全部状态。
1：分销商提交申请，待供应商审核。
2：供应商驳回申请，待分销商确认。
3：供应商修改后，待分销商确认。
4：分销商拒绝修改，待供应商再审核。
5：审核通过下单成功，待分销商付款。
7：付款成功，待供应商发货。
8：供应商发货，待分销商收货。
9：分销商收货，交易成功。
10：采购申请/经销采购单关闭。

注：无值按默认值0计，超出状态范围返回错误信息。 defalutValue��0    */
    OrderStatus  *int64 `json:"order_status,omitempty" required:"false" `
    /*
        查询者自己在所要查询的采购申请/经销采购单中的身份。
1：供应商。
2：分销商。
注：填写其他值当做错误处理。     */
    Identity  *int64 `json:"identity" required:"true" `
    /*
        多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表     */
    Fields  *string `json:"fields,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetStartDate(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetEndDate(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.EndDate = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetPageNo(v int64) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetPageSize(v int64) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetOrderStatus(v int64) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetIdentity(v int64) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.Identity = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetFields(v string) *TaobaoFenxiaoDealerRequisitionorderGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoFenxiaoDealerRequisitionorderGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.OrderStatus != nil) {
        paramMap["order_status"] = *req.OrderStatus
    }
    if(req.Identity != nil) {
        paramMap["identity"] = *req.Identity
    }
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    return paramMap
}

func (req *TaobaoFenxiaoDealerRequisitionorderGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}