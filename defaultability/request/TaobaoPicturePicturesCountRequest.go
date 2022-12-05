package request

import (
        "topsdk/util"
    )

type TaobaoPicturePicturesCountRequest struct {
    /*
        图片ID     */
    PictureId  *int64 `json:"picture_id,omitempty" required:"false" `
    /*
        图片分类     */
    PictureCategoryId  *int64 `json:"picture_category_id,omitempty" required:"false" `
    /*
        是否删除,undeleted代表没有删除,deleted表示删除     */
    Deleted  *string `json:"deleted,omitempty" required:"false" `
    /*
        图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。     */
    StartModifiedDate  *util.LocalTime `json:"start_modified_date,omitempty" required:"false" `
    /*
        查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" required:"false" `
    /*
        查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" required:"false" `
    /*
        图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 defalutValue��client:all    */
    ClientType  *string `json:"client_type,omitempty" required:"false" `
    /*
        文件名     */
    Title  *string `json:"title,omitempty" required:"false" `
}

func (s *TaobaoPicturePicturesCountRequest) SetPictureId(v int64) *TaobaoPicturePicturesCountRequest {
    s.PictureId = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetPictureCategoryId(v int64) *TaobaoPicturePicturesCountRequest {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetDeleted(v string) *TaobaoPicturePicturesCountRequest {
    s.Deleted = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetStartModifiedDate(v util.LocalTime) *TaobaoPicturePicturesCountRequest {
    s.StartModifiedDate = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetStartDate(v util.LocalTime) *TaobaoPicturePicturesCountRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetEndDate(v util.LocalTime) *TaobaoPicturePicturesCountRequest {
    s.EndDate = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetClientType(v string) *TaobaoPicturePicturesCountRequest {
    s.ClientType = &v
    return s
}
func (s *TaobaoPicturePicturesCountRequest) SetTitle(v string) *TaobaoPicturePicturesCountRequest {
    s.Title = &v
    return s
}

func (req *TaobaoPicturePicturesCountRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PictureId != nil) {
        paramMap["picture_id"] = *req.PictureId
    }
    if(req.PictureCategoryId != nil) {
        paramMap["picture_category_id"] = *req.PictureCategoryId
    }
    if(req.Deleted != nil) {
        paramMap["deleted"] = *req.Deleted
    }
    if(req.StartModifiedDate != nil) {
        paramMap["start_modified_date"] = *req.StartModifiedDate
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.ClientType != nil) {
        paramMap["client_type"] = *req.ClientType
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    return paramMap
}

func (req *TaobaoPicturePicturesCountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}