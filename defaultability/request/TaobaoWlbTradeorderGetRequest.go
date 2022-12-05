package request


type TaobaoWlbTradeorderGetRequest struct {
    /*
        交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易     */
    TradeType  *string `json:"trade_type" required:"true" `
    /*
        子交易号     */
    SubTradeId  *string `json:"sub_trade_id,omitempty" required:"false" `
    /*
        指定交易类型的交易号     */
    TradeId  *string `json:"trade_id" required:"true" `
}

func (s *TaobaoWlbTradeorderGetRequest) SetTradeType(v string) *TaobaoWlbTradeorderGetRequest {
    s.TradeType = &v
    return s
}
func (s *TaobaoWlbTradeorderGetRequest) SetSubTradeId(v string) *TaobaoWlbTradeorderGetRequest {
    s.SubTradeId = &v
    return s
}
func (s *TaobaoWlbTradeorderGetRequest) SetTradeId(v string) *TaobaoWlbTradeorderGetRequest {
    s.TradeId = &v
    return s
}

func (req *TaobaoWlbTradeorderGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    if(req.SubTradeId != nil) {
        paramMap["sub_trade_id"] = *req.SubTradeId
    }
    if(req.TradeId != nil) {
        paramMap["trade_id"] = *req.TradeId
    }
    return paramMap
}

func (req *TaobaoWlbTradeorderGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}