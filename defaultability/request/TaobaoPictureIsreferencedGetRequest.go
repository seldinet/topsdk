package request


type TaobaoPictureIsreferencedGetRequest struct {
    /*
        图片id     */
    PictureId  *int64 `json:"picture_id" required:"true" `
}

func (s *TaobaoPictureIsreferencedGetRequest) SetPictureId(v int64) *TaobaoPictureIsreferencedGetRequest {
    s.PictureId = &v
    return s
}

func (req *TaobaoPictureIsreferencedGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureId != nil) {
        paramMap["picture_id"] = *req.PictureId
    }
    return paramMap
}

func (req *TaobaoPictureIsreferencedGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}