package request

import (
        "topsdk/util"
    )

type TaobaoTradeAmountGetRequest struct {
    /*
        订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        交易编号     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoTradeAmountGetRequest) SetFields(v []string) *TaobaoTradeAmountGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTradeAmountGetRequest) SetTid(v int64) *TaobaoTradeAmountGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoTradeAmountGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoTradeAmountGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}