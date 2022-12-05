package request

import (
        "topsdk/util"
    )

type TaobaoWlbNotifyMessagePageGetRequest struct {
    /*
        通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功     */
    MsgCode  *string `json:"msg_code,omitempty" required:"false" `
    /*
        分页查询页数     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页查询的每页页数     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        记录开始时间     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" required:"false" `
    /*
        记录截至时间     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" required:"false" `
    /*
        消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM     */
    Status  *string `json:"status,omitempty" required:"false" `
}

func (s *TaobaoWlbNotifyMessagePageGetRequest) SetMsgCode(v string) *TaobaoWlbNotifyMessagePageGetRequest {
    s.MsgCode = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetRequest) SetPageNo(v int64) *TaobaoWlbNotifyMessagePageGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetRequest) SetPageSize(v int64) *TaobaoWlbNotifyMessagePageGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetRequest) SetStartDate(v util.LocalTime) *TaobaoWlbNotifyMessagePageGetRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetRequest) SetEndDate(v util.LocalTime) *TaobaoWlbNotifyMessagePageGetRequest {
    s.EndDate = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetRequest) SetStatus(v string) *TaobaoWlbNotifyMessagePageGetRequest {
    s.Status = &v
    return s
}

func (req *TaobaoWlbNotifyMessagePageGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MsgCode != nil) {
        paramMap["msg_code"] = *req.MsgCode
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoWlbNotifyMessagePageGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}