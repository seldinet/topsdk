package request

import (
        "topsdk/util"
    )

type TaobaoPictureCategoryGetRequest struct {
    /*
        图片分类ID     */
    PictureCategoryId  *int64 `json:"picture_category_id,omitempty" required:"false" `
    /*
        图片分类名，不支持模糊查询     */
    PictureCategoryName  *string `json:"picture_category_name,omitempty" required:"false" `
    /*
        1     */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        取二级分类时设置为对应父分类id
取一级分类时父分类id设为0
取全部分类的时候不设或设为-1 defalutValue��-1    */
    ParentId  *int64 `json:"parent_id,omitempty" required:"false" `
    /*
        图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" required:"false" `
}

func (s *TaobaoPictureCategoryGetRequest) SetPictureCategoryId(v int64) *TaobaoPictureCategoryGetRequest {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureCategoryGetRequest) SetPictureCategoryName(v string) *TaobaoPictureCategoryGetRequest {
    s.PictureCategoryName = &v
    return s
}
func (s *TaobaoPictureCategoryGetRequest) SetType(v string) *TaobaoPictureCategoryGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoPictureCategoryGetRequest) SetParentId(v int64) *TaobaoPictureCategoryGetRequest {
    s.ParentId = &v
    return s
}
func (s *TaobaoPictureCategoryGetRequest) SetModifiedTime(v util.LocalTime) *TaobaoPictureCategoryGetRequest {
    s.ModifiedTime = &v
    return s
}

func (req *TaobaoPictureCategoryGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureCategoryId != nil) {
        paramMap["picture_category_id"] = *req.PictureCategoryId
    }
    if(req.PictureCategoryName != nil) {
        paramMap["picture_category_name"] = *req.PictureCategoryName
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.ParentId != nil) {
        paramMap["parent_id"] = *req.ParentId
    }
    if(req.ModifiedTime != nil) {
        paramMap["modified_time"] = *req.ModifiedTime
    }
    return paramMap
}

func (req *TaobaoPictureCategoryGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}