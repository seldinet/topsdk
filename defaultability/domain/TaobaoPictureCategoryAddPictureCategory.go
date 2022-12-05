package domain

import (
        "topsdk/util"
    )

type TaobaoPictureCategoryAddPictureCategory struct {
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

func (s *TaobaoPictureCategoryAddPictureCategory) SetPictureCategoryId(v int64) *TaobaoPictureCategoryAddPictureCategory {
    s.PictureCategoryId = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetPictureCategoryName(v string) *TaobaoPictureCategoryAddPictureCategory {
    s.PictureCategoryName = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetPosition(v int64) *TaobaoPictureCategoryAddPictureCategory {
    s.Position = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetType(v string) *TaobaoPictureCategoryAddPictureCategory {
    s.Type = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetTotal(v int64) *TaobaoPictureCategoryAddPictureCategory {
    s.Total = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetCreated(v util.LocalTime) *TaobaoPictureCategoryAddPictureCategory {
    s.Created = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetModified(v util.LocalTime) *TaobaoPictureCategoryAddPictureCategory {
    s.Modified = &v
    return s
}
func (s *TaobaoPictureCategoryAddPictureCategory) SetParentId(v int64) *TaobaoPictureCategoryAddPictureCategory {
    s.ParentId = &v
    return s
}
