package request


type TaobaoFenxiaoProductcatsGetRequest struct {
    /*
        返回字段列表     */
    Fields  *string `json:"fields,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductcatsGetRequest) SetFields(v string) *TaobaoFenxiaoProductcatsGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoFenxiaoProductcatsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductcatsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}