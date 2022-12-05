package domain

import (
        "topsdk/util"
    )

type TaobaoPictureCategoryGetPictureCategory struct {
    /*
        图片分类ID     */
    PictureCategoryId  *int64 `json:"picture_category_id,omitempty" `

    /*
        图片分类名     */
    PictureCategoryName  *string `json:"picture_category_name,omitempty" `

    /*
        图片分类排序     */
    Position  *int64 `json:"position,omitempty" `

    /*
        图片分类型别，sys-fixture代表店铺装修分类(系统分类)，sys-auction代表宝贝图片分类(系统分类)，user-define代表用户自定义分类     */
    Type  *string `json:"type,omitempty" `

    /*
        分类下的图片数     */
    Total  *int64 `json:"total,omitempty" `

    /*
        图片类目的创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        图片分类的修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        一级分类的parent_id为0
二级分类的则为其父分类的picture_category_id     */
    ParentId  *int64 `json:"parent_id,omitempty" `

}

func (s *TaobaoPictureCategoryGetPictureCategory) SetPictureCategoryId(v int64) *TaobaoPictureCategoryGetPictureCategory {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetPictureCategoryName(v string) *TaobaoPictureCategoryGetPictureCategory {
    s.PictureCategoryName = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetPosition(v int64) *TaobaoPictureCategoryGetPictureCategory {
    s.Position = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetType(v string) *TaobaoPictureCategoryGetPictureCategory {
    s.Type = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetTotal(v int64) *TaobaoPictureCategoryGetPictureCategory {
    s.Total = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetCreated(v util.LocalTime) *TaobaoPictureCategoryGetPictureCategory {
    s.Created = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetModified(v util.LocalTime) *TaobaoPictureCategoryGetPictureCategory {
    s.Modified = &v
    return s
}
func (s *TaobaoPictureCategoryGetPictureCategory) SetParentId(v int64) *TaobaoPictureCategoryGetPictureCategory {
    s.ParentId = &v
    return s
}
