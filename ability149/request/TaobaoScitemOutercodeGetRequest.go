package request


type TaobaoScitemOutercodeGetRequest struct {
    /*
        商品编码     */
    OuterCode  *string `json:"outer_code" required:"true" `
}

func (s *TaobaoScitemOutercodeGetRequest) SetOuterCode(v string) *TaobaoScitemOutercodeGetRequest {
    s.OuterCode = &v
    return s
}

func (req *TaobaoScitemOutercodeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterCode != nil) {
        paramMap["outer_code"] = *req.OuterCode
    }
    return paramMap
}

func (req *TaobaoScitemOutercodeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}