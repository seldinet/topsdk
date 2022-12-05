package request

import (
        "topsdk/util"
    )

type TaobaoPictureGetRequest struct {
    /*
        图片url查询接口     */
    Urls  *string `json:"urls,omitempty" required:"false" `
    /*
        是否获取https的链接 defalutValue��false    */
    IsHttps  *bool `json:"is_https,omitempty" required:"false" `
    /*
        图片ID     */
    PictureId  *int64 `json:"picture_id,omitempty" required:"false" `
    /*
        图片分类ID     */
    PictureCategoryId  *int64 `json:"picture_category_id,omitempty" required:"false" `
    /*
        是否删除,unfroze代表没有删除     */
    Deleted  *string `json:"deleted,omitempty" required:"false" `
    /*
        图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" required:"false" `
    /*
        图片标题,最大长度50字符,中英文都算一字符     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" required:"false" `
    /*
        查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" required:"false" `
    /*
        页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数.每页返回最多返回100条,默认值40 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 defalutValue��client:all    */
    ClientType  *string `json:"client_type,omitempty" required:"false" `
    /*
        图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc defalutValue��time:desc    */
    OrderBy  *string `json:"order_by,omitempty" required:"false" `
}

func (s *TaobaoPictureGetRequest) SetUrls(v string) *TaobaoPictureGetRequest {
    s.Urls = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetIsHttps(v bool) *TaobaoPictureGetRequest {
    s.IsHttps = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetPictureId(v int64) *TaobaoPictureGetRequest {
    s.PictureId = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetPictureCategoryId(v int64) *TaobaoPictureGetRequest {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetDeleted(v string) *TaobaoPictureGetRequest {
    s.Deleted = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetModifiedTime(v util.LocalTime) *TaobaoPictureGetRequest {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetTitle(v string) *TaobaoPictureGetRequest {
    s.Title = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetStartDate(v util.LocalTime) *TaobaoPictureGetRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetEndDate(v util.LocalTime) *TaobaoPictureGetRequest {
    s.EndDate = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetPageNo(v int64) *TaobaoPictureGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetPageSize(v int64) *TaobaoPictureGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetClientType(v string) *TaobaoPictureGetRequest {
    s.ClientType = &v
    return s
}
func (s *TaobaoPictureGetRequest) SetOrderBy(v string) *TaobaoPictureGetRequest {
    s.OrderBy = &v
    return s
}

func (req *TaobaoPictureGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Urls != nil) {
        paramMap["urls"] = *req.Urls
    }
    if(req.IsHttps != nil) {
        paramMap["is_https"] = *req.IsHttps
    }
    if(req.PictureId != nil) {
        paramMap["picture_id"] = *req.PictureId
    }
    if(req.PictureCategoryId != nil) {
        paramMap["picture_category_id"] = *req.PictureCategoryId
    }
    if(req.Deleted != nil) {
        paramMap["deleted"] = *req.Deleted
    }
    if(req.ModifiedTime != nil) {
        paramMap["modified_time"] = *req.ModifiedTime
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.ClientType != nil) {
        paramMap["client_type"] = *req.ClientType
    }
    if(req.OrderBy != nil) {
        paramMap["order_by"] = *req.OrderBy
    }
    return paramMap
}

func (req *TaobaoPictureGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}