package request


type TaobaoPictureCategoryUpdateRequest struct {
    /*
        要更新的图片分类的id     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        图片分类的新名字，最大长度20字符，不能为空     */
    CategoryName  *string `json:"category_name,omitempty" required:"false" `
    /*
        图片分类的新父分类id     */
    ParentId  *int64 `json:"parent_id,omitempty" required:"false" `
}

func (s *TaobaoPictureCategoryUpdateRequest) SetCategoryId(v int64) *TaobaoPictureCategoryUpdateRequest {
    s.CategoryId = &v
    return s
}
func (s *TaobaoPictureCategoryUpdateRequest) SetCategoryName(v string) *TaobaoPictureCategoryUpdateRequest {
    s.CategoryName = &v
    return s
}
func (s *TaobaoPictureCategoryUpdateRequest) SetParentId(v int64) *TaobaoPictureCategoryUpdateRequest {
    s.ParentId = &v
    return s
}

func (req *TaobaoPictureCategoryUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.CategoryName != nil) {
        paramMap["category_name"] = *req.CategoryName
    }
    if(req.ParentId != nil) {
        paramMap["parent_id"] = *req.ParentId
    }
    return paramMap
}

func (req *TaobaoPictureCategoryUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}