package domain

import (
        "topsdk/util"
    )

type TaobaoPictureUploadPicture struct {
    /*
        图片ID     */
    PictureId  *int64 `json:"picture_id,omitempty" `

    /*
        图片分类ID     */
    PictureCategoryId  *int64 `json:"picture_category_id,omitempty" `

    /*
        返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg     */
    PicturePath  *string `json:"picture_path,omitempty" `

    /*
        图片标题     */
    Title  *string `json:"title,omitempty" `

    /*
        图片大小,bite单位     */
    Sizes  *int64 `json:"sizes,omitempty" `

    /*
        图片相素,格式:长x宽，如450x150     */
    Pixel  *string `json:"pixel,omitempty" `

    /*
        图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过     */
    Status  *string `json:"status,omitempty" `

    /*
        图片是否删除的标记     */
    Deleted  *string `json:"deleted,omitempty" `

    /*
        图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布     */
    ClientType  *string `json:"client_type,omitempty" `

    /*
        图片的创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        图片的修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoPictureUploadPicture) SetPictureId(v int64) *TaobaoPictureUploadPicture {
    s.PictureId = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetPictureCategoryId(v int64) *TaobaoPictureUploadPicture {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetPicturePath(v string) *TaobaoPictureUploadPicture {
    s.PicturePath = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetTitle(v string) *TaobaoPictureUploadPicture {
    s.Title = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetSizes(v int64) *TaobaoPictureUploadPicture {
    s.Sizes = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetPixel(v string) *TaobaoPictureUploadPicture {
    s.Pixel = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetStatus(v string) *TaobaoPictureUploadPicture {
    s.Status = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetDeleted(v string) *TaobaoPictureUploadPicture {
    s.Deleted = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetClientType(v string) *TaobaoPictureUploadPicture {
    s.ClientType = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetCreated(v util.LocalTime) *TaobaoPictureUploadPicture {
    s.Created = &v
    return s
}
func (s *TaobaoPictureUploadPicture) SetModified(v util.LocalTime) *TaobaoPictureUploadPicture {
    s.Modified = &v
    return s
}
