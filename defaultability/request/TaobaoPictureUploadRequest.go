package request

import (
        "topsdk/util"
    )

type TaobaoPictureUploadRequest struct {
    /*
        如果此参数大于0，而且在后台能查到对应的图片，则此次上传为原图替换 defalutValue��0    */
    PictureId  *int64 `json:"picture_id,omitempty" required:"false" `
    /*
        图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类     */
    PictureCategoryId  *int64 `json:"picture_category_id" required:"true" `
    /*
        图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。     */
    Img  *[]byte `json:"img" required:"true" `
    /*
        包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名     */
    ImageInputTitle  *string `json:"image_input_title" required:"true" `
    /*
        图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。 defalutValue��client:computer    */
    ClientType  *string `json:"client_type,omitempty" required:"false" `
    /*
        是否获取https连接 defalutValue��false    */
    IsHttps  *bool `json:"is_https,omitempty" required:"false" `
}

func (s *TaobaoPictureUploadRequest) SetPictureId(v int64) *TaobaoPictureUploadRequest {
    s.PictureId = &v
    return s
}
func (s *TaobaoPictureUploadRequest) SetPictureCategoryId(v int64) *TaobaoPictureUploadRequest {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureUploadRequest) SetImg(v []byte) *TaobaoPictureUploadRequest {
    s.Img = &v
    return s
}
func (s *TaobaoPictureUploadRequest) SetImageInputTitle(v string) *TaobaoPictureUploadRequest {
    s.ImageInputTitle = &v
    return s
}
func (s *TaobaoPictureUploadRequest) SetTitle(v string) *TaobaoPictureUploadRequest {
    s.Title = &v
    return s
}
func (s *TaobaoPictureUploadRequest) SetClientType(v string) *TaobaoPictureUploadRequest {
    s.ClientType = &v
    return s
}
func (s *TaobaoPictureUploadRequest) SetIsHttps(v bool) *TaobaoPictureUploadRequest {
    s.IsHttps = &v
    return s
}

func (req *TaobaoPictureUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureId != nil) {
        paramMap["picture_id"] = *req.PictureId
    }
    if(req.PictureCategoryId != nil) {
        paramMap["picture_category_id"] = *req.PictureCategoryId
    }
    if(req.ImageInputTitle != nil) {
        paramMap["image_input_title"] = *req.ImageInputTitle
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.ClientType != nil) {
        paramMap["client_type"] = *req.ClientType
    }
    if(req.IsHttps != nil) {
        paramMap["is_https"] = *req.IsHttps
    }
    return paramMap
}

func (req *TaobaoPictureUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Img != nil {
        fileMap["img"] = *req.Img
    }
    return fileMap
}