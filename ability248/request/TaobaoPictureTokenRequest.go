package request

import (
        "topsdk/ability248/domain"
        "topsdk/util"
    )

type TaobaoPictureTokenRequest struct {
    /*
        请求参数     */
    GenerateTokenRequest  *domain.TaobaoPictureTokenGenerateTokenRequest `json:"generate_token_request" required:"true" `
}

func (s *TaobaoPictureTokenRequest) SetGenerateTokenRequest(v domain.TaobaoPictureTokenGenerateTokenRequest) *TaobaoPictureTokenRequest {
    s.GenerateTokenRequest = &v
    return s
}

func (req *TaobaoPictureTokenRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GenerateTokenRequest != nil) {
        paramMap["generate_token_request"] = util.ConvertStruct(*req.GenerateTokenRequest)
    }
    return paramMap
}

func (req *TaobaoPictureTokenRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}