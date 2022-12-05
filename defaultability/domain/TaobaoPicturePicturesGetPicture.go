package domain

import (
        "topsdk/util"
    )

type TaobaoPicturePicturesGetPicture struct {
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

    /*
        图片在后台处理之后的md5值当md5值为32位长度的字符串时为图片搬家后的文件md5验证码md5值为长整数时为图片替换后的时间戳     */
    Md5  *string `json:"md5,omitempty" `

    /*
        图片是否被引用     */
    Referenced  *bool `json:"referenced,omitempty" `

}

func (s *TaobaoPicturePicturesGetPicture) SetPictureId(v int64) *TaobaoPicturePicturesGetPicture {
    s.PictureId = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetPictureCategoryId(v int64) *TaobaoPicturePicturesGetPicture {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetPicturePath(v string) *TaobaoPicturePicturesGetPicture {
    s.PicturePath = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetTitle(v string) *TaobaoPicturePicturesGetPicture {
    s.Title = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetSizes(v int64) *TaobaoPicturePicturesGetPicture {
    s.Sizes = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetPixel(v string) *TaobaoPicturePicturesGetPicture {
    s.Pixel = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetStatus(v string) *TaobaoPicturePicturesGetPicture {
    s.Status = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetDeleted(v string) *TaobaoPicturePicturesGetPicture {
    s.Deleted = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetClientType(v string) *TaobaoPicturePicturesGetPicture {
    s.ClientType = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetCreated(v util.LocalTime) *TaobaoPicturePicturesGetPicture {
    s.Created = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetModified(v util.LocalTime) *TaobaoPicturePicturesGetPicture {
    s.Modified = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetMd5(v string) *TaobaoPicturePicturesGetPicture {
    s.Md5 = &v
    return s
}
func (s *TaobaoPicturePicturesGetPicture) SetReferenced(v bool) *TaobaoPicturePicturesGetPicture {
    s.Referenced = &v
    return s
}
