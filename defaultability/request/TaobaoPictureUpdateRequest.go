package request


type TaobaoPictureUpdateRequest struct {
    /*
        要更改名字的图片的id     */
    PictureId  *int64 `json:"picture_id" required:"true" `
    /*
        新的图片名，最大长度50字符，不能为空     */
    NewName  *string `json:"new_name" required:"true" `
}

func (s *TaobaoPictureUpdateRequest) SetPictureId(v int64) *TaobaoPictureUpdateRequest {
    s.PictureId = &v
    return s
}
func (s *TaobaoPictureUpdateRequest) SetNewName(v string) *TaobaoPictureUpdateRequest {
    s.NewName = &v
    return s
}

func (req *TaobaoPictureUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureId != nil) {
        paramMap["picture_id"] = *req.PictureId
    }
    if(req.NewName != nil) {
        paramMap["new_name"] = *req.NewName
    }
    return paramMap
}

func (req *TaobaoPictureUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}