package request

import (
        "topsdk/util"
    )

type TaobaoPictureReplaceRequest struct {
    /*
        要替换的图片的id，必须大于0     */
    PictureId  *int64 `json:"picture_id" required:"true" `
    /*
        图片二进制文件流,不能为空,允许png、jpg、gif图片格式     */
    ImageData  *[]byte `json:"image_data" required:"true" `
}

func (s *TaobaoPictureReplaceRequest) SetPictureId(v int64) *TaobaoPictureReplaceRequest {
    s.PictureId = &v
    return s
}
func (s *TaobaoPictureReplaceRequest) SetImageData(v []byte) *TaobaoPictureReplaceRequest {
    s.ImageData = &v
    return s
}

func (req *TaobaoPictureReplaceRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureId != nil) {
        paramMap["picture_id"] = *req.PictureId
    }
    return paramMap
}

func (req *TaobaoPictureReplaceRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.ImageData != nil {
        fileMap["image_data"] = *req.ImageData
    }
    return fileMap
}