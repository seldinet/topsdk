package request

import (
        "topsdk/util"
    )

type TaobaoVmarketEticketOplogsGetRequest struct {
    /*
        0:全部 1:核销 2:冲正     */
    Type  *int64 `json:"type" required:"true" `
    /*
        开始时间     */
    StartTime  *util.LocalTime `json:"start_time,omitempty" required:"false" `
    /*
        结束时间     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" required:"false" `
    /*
        核销码     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        手机号后四位     */
    Mobile  *string `json:"mobile,omitempty" required:"false" `
    /*
        核销身份     */
    Posid  *string `json:"posid,omitempty" required:"false" `
    /*
        码商ID     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
    /*
        当前页码 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页显示的记录数，最大为40，默认为40 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        排序方式     */
    Sort  *string `json:"sort,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketOplogsGetRequest) SetType(v int64) *TaobaoVmarketEticketOplogsGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetStartTime(v util.LocalTime) *TaobaoVmarketEticketOplogsGetRequest {
    s.StartTime = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetEndTime(v util.LocalTime) *TaobaoVmarketEticketOplogsGetRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetCode(v string) *TaobaoVmarketEticketOplogsGetRequest {
    s.Code = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetMobile(v string) *TaobaoVmarketEticketOplogsGetRequest {
    s.Mobile = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetPosid(v string) *TaobaoVmarketEticketOplogsGetRequest {
    s.Posid = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketOplogsGetRequest {
    s.CodemerchantId = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetPageNo(v int64) *TaobaoVmarketEticketOplogsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetPageSize(v int64) *TaobaoVmarketEticketOplogsGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetRequest) SetSort(v string) *TaobaoVmarketEticketOplogsGetRequest {
    s.Sort = &v
    return s
}

func (req *TaobaoVmarketEticketOplogsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.Mobile != nil) {
        paramMap["mobile"] = *req.Mobile
    }
    if(req.Posid != nil) {
        paramMap["posid"] = *req.Posid
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Sort != nil) {
        paramMap["sort"] = *req.Sort
    }
    return paramMap
}

func (req *TaobaoVmarketEticketOplogsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}