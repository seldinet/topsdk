package request


type TaobaoItemAnchorGetRequest struct {
    /*
        宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)     */
    Type  *int64 `json:"type" required:"true" `
    /*
        对应类目编号     */
    CatId  *int64 `json:"cat_id" required:"true" `
}

func (s *TaobaoItemAnchorGetRequest) SetType(v int64) *TaobaoItemAnchorGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoItemAnchorGetRequest) SetCatId(v int64) *TaobaoItemAnchorGetRequest {
    s.CatId = &v
    return s
}

func (req *TaobaoItemAnchorGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    return paramMap
}

func (req *TaobaoItemAnchorGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}