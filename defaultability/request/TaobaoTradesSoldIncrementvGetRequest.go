package request

import (
        "topsdk/util"
    )

type TaobaoTradesSoldIncrementvGetRequest struct {
    /*
        需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss     */
    StartCreate  *util.LocalTime `json:"start_create" required:"true" `
    /*
        查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。     */
    EndCreate  *util.LocalTime `json:"end_create" required:"true" `
    /*
        交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。     */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型     */
    ExtType  *string `json:"ext_type,omitempty" required:"false" `
    /*
        卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）     */
    Tag  *string `json:"tag,omitempty" required:"false" `
    /*
        页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 defalutValue��false    */
    UseHasNext  *bool `json:"use_has_next,omitempty" required:"false" `
}

func (s *TaobaoTradesSoldIncrementvGetRequest) SetFields(v []string) *TaobaoTradesSoldIncrementvGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetStartCreate(v util.LocalTime) *TaobaoTradesSoldIncrementvGetRequest {
    s.StartCreate = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetEndCreate(v util.LocalTime) *TaobaoTradesSoldIncrementvGetRequest {
    s.EndCreate = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetStatus(v string) *TaobaoTradesSoldIncrementvGetRequest {
    s.Status = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetType(v string) *TaobaoTradesSoldIncrementvGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetExtType(v string) *TaobaoTradesSoldIncrementvGetRequest {
    s.ExtType = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetTag(v string) *TaobaoTradesSoldIncrementvGetRequest {
    s.Tag = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetPageNo(v int64) *TaobaoTradesSoldIncrementvGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetPageSize(v int64) *TaobaoTradesSoldIncrementvGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetRequest) SetUseHasNext(v bool) *TaobaoTradesSoldIncrementvGetRequest {
    s.UseHasNext = &v
    return s
}

func (req *TaobaoTradesSoldIncrementvGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.StartCreate != nil) {
        paramMap["start_create"] = *req.StartCreate
    }
    if(req.EndCreate != nil) {
        paramMap["end_create"] = *req.EndCreate
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.ExtType != nil) {
        paramMap["ext_type"] = *req.ExtType
    }
    if(req.Tag != nil) {
        paramMap["tag"] = *req.Tag
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.UseHasNext != nil) {
        paramMap["use_has_next"] = *req.UseHasNext
    }
    return paramMap
}

func (req *TaobaoTradesSoldIncrementvGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}