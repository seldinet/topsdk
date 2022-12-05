package domain


type TaobaoPictureTokenUploadPolicy struct {
    /*
        限制用户上传文件的大小。 若用户上传文件大小超过size_limit，无法上传成功。     */
    SizeLimit  *int64 `json:"size_limit,omitempty" `

    /*
        限制用户上传文件的类型，多个值用；分隔。 可设置的值为：image/jpeg,image/png,image/webp等。 若用户上传文件的mime类型不在mime_limit范围内，无法上传成功。     */
    MimeLimit  *string `json:"mime_limit,omitempty" `

    /*
        图片分类ID，可通过taobao.picture.category.get获取。根目录值为0。     */
    DirId  *int64 `json:"dir_id,omitempty" `

    /*
        token有效期的截止时间，值为自1970年1月1日0时起的毫秒数。若当前时间晚于expired_time，token失效，上传文件失败。     */
    ExpiredTime  *int64 `json:"expired_time,omitempty" `

}

func (s *TaobaoPictureTokenUploadPolicy) SetSizeLimit(v int64) *TaobaoPictureTokenUploadPolicy {
    s.SizeLimit = &v
    return s
}
func (s *TaobaoPictureTokenUploadPolicy) SetMimeLimit(v string) *TaobaoPictureTokenUploadPolicy {
    s.MimeLimit = &v
    return s
}
func (s *TaobaoPictureTokenUploadPolicy) SetDirId(v int64) *TaobaoPictureTokenUploadPolicy {
    s.DirId = &v
    return s
}
func (s *TaobaoPictureTokenUploadPolicy) SetExpiredTime(v int64) *TaobaoPictureTokenUploadPolicy {
    s.ExpiredTime = &v
    return s
}
