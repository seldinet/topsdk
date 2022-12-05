package request

import (
        "topsdk/util"
    )

type TaobaoSpecialRefundGetRequest struct {
    /*
        需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        退款单号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoSpecialRefundGetRequest) SetFields(v []string) *TaobaoSpecialRefundGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoSpecialRefundGetRequest) SetRefundId(v int64) *TaobaoSpecialRefundGetRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoSpecialRefundGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoSpecialRefundGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}