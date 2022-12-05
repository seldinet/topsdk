package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoCooperationGetRequest struct {
    /*
        合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        合作开始时间yyyy-MM-dd HH:mm:ss     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" required:"false" `
    /*
        合作结束时间yyyy-MM-dd HH:mm:ss     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" required:"false" `
    /*
        分销方式：AGENT(代销) 、DEALER（经销）     */
    TradeType  *string `json:"trade_type,omitempty" required:"false" `
    /*
        页码（大于0的整数，默认1）     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页记录数（默认20，最大50）     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        渠道code     */
    ChannelCode  *string `json:"channel_code,omitempty" required:"false" `
    /*
        1是供应商，2 是分销商     */
    RoleType  *int64 `json:"role_type,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoCooperationGetRequest) SetStatus(v string) *TaobaoFenxiaoCooperationGetRequest {
    s.Status = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetStartDate(v util.LocalTime) *TaobaoFenxiaoCooperationGetRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetEndDate(v util.LocalTime) *TaobaoFenxiaoCooperationGetRequest {
    s.EndDate = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetTradeType(v string) *TaobaoFenxiaoCooperationGetRequest {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetPageNo(v int64) *TaobaoFenxiaoCooperationGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetPageSize(v int64) *TaobaoFenxiaoCooperationGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetChannelCode(v string) *TaobaoFenxiaoCooperationGetRequest {
    s.ChannelCode = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetRequest) SetRoleType(v int64) *TaobaoFenxiaoCooperationGetRequest {
    s.RoleType = &v
    return s
}

func (req *TaobaoFenxiaoCooperationGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.ChannelCode != nil) {
        paramMap["channel_code"] = *req.ChannelCode
    }
    if(req.RoleType != nil) {
        paramMap["role_type"] = *req.RoleType
    }
    return paramMap
}

func (req *TaobaoFenxiaoCooperationGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}