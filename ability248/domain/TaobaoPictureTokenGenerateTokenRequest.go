package domain


type TaobaoPictureTokenGenerateTokenRequest struct {
    /*
        请求策略     */
    UploadPolicy  *TaobaoPictureTokenUploadPolicy `json:"upload_policy,omitempty" `

}

func (s *TaobaoPictureTokenGenerateTokenRequest) SetUploadPolicy(v TaobaoPictureTokenUploadPolicy) *TaobaoPictureTokenGenerateTokenRequest {
    s.UploadPolicy = &v
    return s
}
