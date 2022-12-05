package request

import (
        "topsdk/util"
    )

type TaobaoWlbOrderdetailDateGetRequest struct {
    /*
        创建时间起始     */
    StartTime  *util.LocalTime `json:"start_time" required:"true" `
    /*
        创建时间结束     */
    EndTime  *util.LocalTime `json:"end_time" required:"true" `
    /*
        分页大小 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        分页下标 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoWlbOrderdetailDateGetRequest) SetStartTime(v util.LocalTime) *TaobaoWlbOrderdetailDateGetRequest {
    s.StartTime = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetRequest) SetEndTime(v util.LocalTime) *TaobaoWlbOrderdetailDateGetRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetRequest) SetPageSize(v int64) *TaobaoWlbOrderdetailDateGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetRequest) SetPageNo(v int64) *TaobaoWlbOrderdetailDateGetRequest {
    s.PageNo = &v
    return s
}

func (req *TaobaoWlbOrderdetailDateGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    return paramMap
}

func (req *TaobaoWlbOrderdetailDateGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}