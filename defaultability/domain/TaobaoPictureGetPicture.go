package domain

import (
        "topsdk/util"
    )

type TaobaoPictureGetPicture struct {
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
        图片状态,0 未审核没冻结 1  冻结 2 审核通过     */
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

    /*
        图片在后台处理之后的md5值当md5值为32位长度的字符串时为图片搬家后的文件md5验证码md5值为长整数时为图片替换后的时间戳     */
    Md5  *string `json:"md5,omitempty" `

    /*
        图片是否被引用     */
    Referenced  *bool `json:"referenced,omitempty" `

}

func (s *TaobaoPictureGetPicture) SetPictureId(v int64) *TaobaoPictureGetPicture {
    s.PictureId = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetPictureCategoryId(v int64) *TaobaoPictureGetPicture {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetPicturePath(v string) *TaobaoPictureGetPicture {
    s.PicturePath = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetTitle(v string) *TaobaoPictureGetPicture {
    s.Title = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetSizes(v int64) *TaobaoPictureGetPicture {
    s.Sizes = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetPixel(v string) *TaobaoPictureGetPicture {
    s.Pixel = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetStatus(v string) *TaobaoPictureGetPicture {
    s.Status = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetDeleted(v string) *TaobaoPictureGetPicture {
    s.Deleted = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetClientType(v string) *TaobaoPictureGetPicture {
    s.ClientType = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetCreated(v util.LocalTime) *TaobaoPictureGetPicture {
    s.Created = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetModified(v util.LocalTime) *TaobaoPictureGetPicture {
    s.Modified = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetMd5(v string) *TaobaoPictureGetPicture {
    s.Md5 = &v
    return s
}
func (s *TaobaoPictureGetPicture) SetReferenced(v bool) *TaobaoPictureGetPicture {
    s.Referenced = &v
    return s
}
