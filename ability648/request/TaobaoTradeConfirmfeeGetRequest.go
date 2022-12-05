package request


type TaobaoTradeConfirmfeeGetRequest struct {
    /*
        交易主订单或子订单ID     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoTradeConfirmfeeGetRequest) SetTid(v int64) *TaobaoTradeConfirmfeeGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoTradeConfirmfeeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoTradeConfirmfeeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}