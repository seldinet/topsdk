package request

import (
        "topsdk/util"
    )

type TaobaoTradeGetRequest struct {
    /*
        需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        交易编号     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoTradeGetRequest) SetFields(v []string) *TaobaoTradeGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTradeGetRequest) SetTid(v int64) *TaobaoTradeGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoTradeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoTradeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}