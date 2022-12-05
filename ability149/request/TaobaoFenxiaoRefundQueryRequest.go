package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoRefundQueryRequest struct {
    /*
        查询退款单的修改时间开始,格式如:yyyy-MM-dd HH:mm:ss     */
    StartDate  *util.LocalTime `json:"start_date" required:"true" `
    /*
        查询退款单的修改时间结束,格式如:yyyy-MM-dd HH:mm:ss     */
    EndDate  *util.LocalTime `json:"end_date" required:"true" `
    /*
        页码（大于0的整数。无值或小于1的值按默认值1计） defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） defalutValue��50    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        是否查询下游消费者的退款信息     */
    QuerySellerRefund  *bool `json:"query_seller_refund,omitempty" required:"false" `
    /*
        查询的经营模式，当不指定时，默认查询代销订单 代销：1 经销：2 寄售（自营寄售）：5 平台寄售     */
    TradeTypes  *[]int32 `json:"trade_types,omitempty" required:"false" `
    /*
        角色，供应商：2，分销商：1     */
    UserRoleType  *int64 `json:"user_role_type,omitempty" required:"false" `
    /*
        渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场     */
    ChannelCodes  *[]int32 `json:"channel_codes,omitempty" required:"false" `
    /*
        退款流程类型：4：发货前退款；1：发货后退款不退货；2：发货后退款退货；3：售后退款；5：拒收；6：售后退货退款     */
    RefundFlowTypes  *[]int32 `json:"refund_flow_types,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoRefundQueryRequest) SetStartDate(v util.LocalTime) *TaobaoFenxiaoRefundQueryRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetEndDate(v util.LocalTime) *TaobaoFenxiaoRefundQueryRequest {
    s.EndDate = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetPageNo(v int64) *TaobaoFenxiaoRefundQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetPageSize(v int64) *TaobaoFenxiaoRefundQueryRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetQuerySellerRefund(v bool) *TaobaoFenxiaoRefundQueryRequest {
    s.QuerySellerRefund = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetTradeTypes(v []int32) *TaobaoFenxiaoRefundQueryRequest {
    s.TradeTypes = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetUserRoleType(v int64) *TaobaoFenxiaoRefundQueryRequest {
    s.UserRoleType = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetChannelCodes(v []int32) *TaobaoFenxiaoRefundQueryRequest {
    s.ChannelCodes = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRequest) SetRefundFlowTypes(v []int32) *TaobaoFenxiaoRefundQueryRequest {
    s.RefundFlowTypes = &v
    return s
}

func (req *TaobaoFenxiaoRefundQueryRequest) ToMap() map[string]interface{} {
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
    if(req.QuerySellerRefund != nil) {
        paramMap["query_seller_refund"] = *req.QuerySellerRefund
    }
    if(req.TradeTypes != nil) {
        paramMap["trade_types"] = util.ConvertBasicList(*req.TradeTypes)
    }
    if(req.UserRoleType != nil) {
        paramMap["user_role_type"] = *req.UserRoleType
    }
    if(req.ChannelCodes != nil) {
        paramMap["channel_codes"] = util.ConvertBasicList(*req.ChannelCodes)
    }
    if(req.RefundFlowTypes != nil) {
        paramMap["refund_flow_types"] = util.ConvertBasicList(*req.RefundFlowTypes)
    }
    return paramMap
}

func (req *TaobaoFenxiaoRefundQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}