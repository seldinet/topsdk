package request

import (
        "topsdk/util"
    )

type TaobaoTradeFullinfoGetRequest struct {
    /*
        include_oaid defalutValue��false    */
    IncludeOaid  *string `json:"include_oaid,omitempty" required:"false" `
    /*
        需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        交易编号     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoTradeFullinfoGetRequest) SetIncludeOaid(v string) *TaobaoTradeFullinfoGetRequest {
    s.IncludeOaid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetRequest) SetFields(v []string) *TaobaoTradeFullinfoGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTradeFullinfoGetRequest) SetTid(v int64) *TaobaoTradeFullinfoGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoTradeFullinfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.IncludeOaid != nil) {
        paramMap["include_oaid"] = *req.IncludeOaid
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoTradeFullinfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}