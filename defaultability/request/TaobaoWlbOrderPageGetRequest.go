package request

import (
        "topsdk/util"
    )

type TaobaoWlbOrderPageGetRequest struct {
    /*
        订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库     */
    OrderType  *string `json:"order_type,omitempty" required:"false" `
    /*
        订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单     */
    OrderSubType  *string `json:"order_sub_type,omitempty" required:"false" `
    /*
        每页多少条 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        分页的第几页 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        TMS拒签：-100 接收方拒签：-200     */
    OrderStatus  *int64 `json:"order_status,omitempty" required:"false" `
    /*
        物流订单编号     */
    OrderCode  *string `json:"order_code,omitempty" required:"false" `
    /*
        查询截止时间     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" required:"false" `
    /*
        查询开始时间     */
    StartTime  *util.LocalTime `json:"start_time,omitempty" required:"false" `
}

func (s *TaobaoWlbOrderPageGetRequest) SetOrderType(v string) *TaobaoWlbOrderPageGetRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetOrderSubType(v string) *TaobaoWlbOrderPageGetRequest {
    s.OrderSubType = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetPageSize(v int64) *TaobaoWlbOrderPageGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetPageNo(v int64) *TaobaoWlbOrderPageGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetOrderStatus(v int64) *TaobaoWlbOrderPageGetRequest {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetOrderCode(v string) *TaobaoWlbOrderPageGetRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetEndTime(v util.LocalTime) *TaobaoWlbOrderPageGetRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoWlbOrderPageGetRequest) SetStartTime(v util.LocalTime) *TaobaoWlbOrderPageGetRequest {
    s.StartTime = &v
    return s
}

func (req *TaobaoWlbOrderPageGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.OrderSubType != nil) {
        paramMap["order_sub_type"] = *req.OrderSubType
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.OrderStatus != nil) {
        paramMap["order_status"] = *req.OrderStatus
    }
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    return paramMap
}

func (req *TaobaoWlbOrderPageGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}