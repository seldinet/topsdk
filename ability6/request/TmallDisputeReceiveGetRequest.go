package request

import (
        "topsdk/util"
    )

type TmallDisputeReceiveGetRequest struct {
    /*
        退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);CLOSED(退款关闭); SUCCESS(退款成功);SELLER_REFUSE_BUYER(卖家拒绝退款);WAIT_BUYER_CONFIRM_REDO_SEND_GOODS(等待买家确认重新邮寄的货物);WAIT_SELLER_CONFIRM_RETURN_ADDRESS(等待卖家确认退货地址);WAIT_SELLER_CONSIGN_GOOGDS(卖家确认收货,等待卖家发货);EXCHANGE_TRANSFORM_TO_REFUND(换货关闭,转退货退款);EXCHANGE_WAIT_BUYER_CONFIRM_GOODS(卖家已发货,等待买家确认收货);POST_FEE_DISPUTE_WAIT_ACTIVATE(邮费单已创建,待激活)     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        每页条数。取值范围:大于零的整数; 默认值:20;最大值:100 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。     */
    UseHasNext  *bool `json:"use_has_next,omitempty" required:"false" `
    /*
        交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，查看可选值 defalutValue��fixed    */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        逆向纠纷单号id     */
    RefundId  *int64 `json:"refund_id,omitempty" required:"false" `
    /*
        页码。取值范围:大于零的整数; 默认值:1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" required:"false" `
    /*
        查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        需要返回的字段。目前支持有：refund_id, alipay_no, tid, buyer_nick, seller_nick, status, created, modified, order_status, refund_fee, good_status, show_return_logistic(展现买家退货的物流信息), show_exchange_logistic(展现换货的物流信息), time_out, oid, refund_version, title, num, dispute_request, reason, desc     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        买家openId     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" required:"false" `
}

func (s *TmallDisputeReceiveGetRequest) SetStatus(v string) *TmallDisputeReceiveGetRequest {
    s.Status = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetPageSize(v int64) *TmallDisputeReceiveGetRequest {
    s.PageSize = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetUseHasNext(v bool) *TmallDisputeReceiveGetRequest {
    s.UseHasNext = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetType(v string) *TmallDisputeReceiveGetRequest {
    s.Type = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetRefundId(v int64) *TmallDisputeReceiveGetRequest {
    s.RefundId = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetPageNo(v int64) *TmallDisputeReceiveGetRequest {
    s.PageNo = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetBuyerNick(v string) *TmallDisputeReceiveGetRequest {
    s.BuyerNick = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetStartModified(v util.LocalTime) *TmallDisputeReceiveGetRequest {
    s.StartModified = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetEndModified(v util.LocalTime) *TmallDisputeReceiveGetRequest {
    s.EndModified = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetFields(v []string) *TmallDisputeReceiveGetRequest {
    s.Fields = &v
    return s
}
func (s *TmallDisputeReceiveGetRequest) SetBuyerOpenUid(v string) *TmallDisputeReceiveGetRequest {
    s.BuyerOpenUid = &v
    return s
}

func (req *TmallDisputeReceiveGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.UseHasNext != nil) {
        paramMap["use_has_next"] = *req.UseHasNext
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.BuyerOpenUid != nil) {
        paramMap["buyer_open_uid"] = *req.BuyerOpenUid
    }
    return paramMap
}

func (req *TmallDisputeReceiveGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}