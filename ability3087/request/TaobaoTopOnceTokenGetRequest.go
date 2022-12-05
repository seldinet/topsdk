package request


type TaobaoTopOnceTokenGetRequest struct {
    /*
        sec_token     */
    SecToken  *string `json:"sec_token" required:"true" `
}

func (s *TaobaoTopOnceTokenGetRequest) SetSecToken(v string) *TaobaoTopOnceTokenGetRequest {
    s.SecToken = &v
    return s
}

func (req *TaobaoTopOnceTokenGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SecToken != nil) {
        paramMap["sec_token"] = *req.SecToken
    }
    return paramMap
}

func (req *TaobaoTopOnceTokenGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}