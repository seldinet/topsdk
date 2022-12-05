package request

import (
        "topsdk/util"
    )

type TaobaoSpecialRefundsReceiveGetRequest struct {
    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" required:"false" `
    /*
        查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        页码。取值范围:大于零的整数; 默认值:1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>     */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 defalutValue��false    */
    UseHasNext  *bool `json:"use_has_next,omitempty" required:"false" `
    /*
        买家openId     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" required:"false" `
}

func (s *TaobaoSpecialRefundsReceiveGetRequest) SetBuyerNick(v string) *TaobaoSpecialRefundsReceiveGetRequest {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetEndModified(v util.LocalTime) *TaobaoSpecialRefundsReceiveGetRequest {
    s.EndModified = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetFields(v []string) *TaobaoSpecialRefundsReceiveGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetPageNo(v int64) *TaobaoSpecialRefundsReceiveGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetPageSize(v int64) *TaobaoSpecialRefundsReceiveGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetStartModified(v util.LocalTime) *TaobaoSpecialRefundsReceiveGetRequest {
    s.StartModified = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetStatus(v string) *TaobaoSpecialRefundsReceiveGetRequest {
    s.Status = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetType(v string) *TaobaoSpecialRefundsReceiveGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetUseHasNext(v bool) *TaobaoSpecialRefundsReceiveGetRequest {
    s.UseHasNext = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRequest) SetBuyerOpenUid(v string) *TaobaoSpecialRefundsReceiveGetRequest {
    s.BuyerOpenUid = &v
    return s
}

func (req *TaobaoSpecialRefundsReceiveGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.UseHasNext != nil) {
        paramMap["use_has_next"] = *req.UseHasNext
    }
    if(req.BuyerOpenUid != nil) {
        paramMap["buyer_open_uid"] = *req.BuyerOpenUid
    }
    return paramMap
}

func (req *TaobaoSpecialRefundsReceiveGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}