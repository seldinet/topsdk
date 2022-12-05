package domain


type TaobaoPictureTokenGenerateTokenResponse struct {
    /*
        返回的token     */
    Token  *string `json:"token,omitempty" `

}

func (s *TaobaoPictureTokenGenerateTokenResponse) SetToken(v string) *TaobaoPictureTokenGenerateTokenResponse {
    s.Token = &v
    return s
}
