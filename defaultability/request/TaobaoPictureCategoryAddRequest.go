package request


type TaobaoPictureCategoryAddRequest struct {
    /*
        图片分类名称，最大长度20字符，中文字符算2个字符，不能为空     */
    PictureCategoryName  *string `json:"picture_category_name" required:"true" `
    /*
        图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id defalutValue��0    */
    ParentId  *int64 `json:"parent_id,omitempty" required:"false" `
}

func (s *TaobaoPictureCategoryAddRequest) SetPictureCategoryName(v string) *TaobaoPictureCategoryAddRequest {
    s.PictureCategoryName = &v
    return s
}
func (s *TaobaoPictureCategoryAddRequest) SetParentId(v int64) *TaobaoPictureCategoryAddRequest {
    s.ParentId = &v
    return s
}

func (req *TaobaoPictureCategoryAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureCategoryName != nil) {
        paramMap["picture_category_name"] = *req.PictureCategoryName
    }
    if(req.ParentId != nil) {
        paramMap["parent_id"] = *req.ParentId
    }
    return paramMap
}

func (req *TaobaoPictureCategoryAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}