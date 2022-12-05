package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoRequisitionsGetRequest struct {
    /*
        申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）     */
    Status  *int64 `json:"status,omitempty" required:"false" `
    /*
        申请开始时间yyyy-MM-dd     */
    ApplyStart  *util.LocalTime `json:"apply_start,omitempty" required:"false" `
    /*
        申请结束时间yyyy-MM-dd     */
    ApplyEnd  *util.LocalTime `json:"apply_end,omitempty" required:"false" `
    /*
        页码（大于0的整数，默认1）     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页记录数（默认20，最大50）     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoRequisitionsGetRequest) SetStatus(v int64) *TaobaoFenxiaoRequisitionsGetRequest {
    s.Status = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequest) SetApplyStart(v util.LocalTime) *TaobaoFenxiaoRequisitionsGetRequest {
    s.ApplyStart = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequest) SetApplyEnd(v util.LocalTime) *TaobaoFenxiaoRequisitionsGetRequest {
    s.ApplyEnd = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequest) SetPageNo(v int64) *TaobaoFenxiaoRequisitionsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequest) SetPageSize(v int64) *TaobaoFenxiaoRequisitionsGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoFenxiaoRequisitionsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.ApplyStart != nil) {
        paramMap["apply_start"] = *req.ApplyStart
    }
    if(req.ApplyEnd != nil) {
        paramMap["apply_end"] = *req.ApplyEnd
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoFenxiaoRequisitionsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}