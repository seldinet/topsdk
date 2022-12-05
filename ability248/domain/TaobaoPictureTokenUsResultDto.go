package domain


type TaobaoPictureTokenUsResultDto struct {
    /*
        返回码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        响应response     */
    Data  *TaobaoPictureTokenGenerateTokenResponse `json:"data,omitempty" `

    /*
        响应结果     */
    Success  *bool `json:"success,omitempty" `

    /*
        响应结果说明     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoPictureTokenUsResultDto) SetCode(v int64) *TaobaoPictureTokenUsResultDto {
    s.Code = &v
    return s
}
func (s *TaobaoPictureTokenUsResultDto) SetData(v TaobaoPictureTokenGenerateTokenResponse) *TaobaoPictureTokenUsResultDto {
    s.Data = &v
    return s
}
func (s *TaobaoPictureTokenUsResultDto) SetSuccess(v bool) *TaobaoPictureTokenUsResultDto {
    s.Success = &v
    return s
}
func (s *TaobaoPictureTokenUsResultDto) SetErrorMsg(v string) *TaobaoPictureTokenUsResultDto {
    s.ErrorMsg = &v
    return s
}
