package domain

import (
        "topsdk/util"
    )

type TaobaoItemImgUploadItemImg struct {
    /*
        商品图片的id，和商品相对应（主图id默认为0）     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片链接地址     */
    Url  *string `json:"url,omitempty" `

    /*
        图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoItemImgUploadItemImg) SetId(v int64) *TaobaoItemImgUploadItemImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemImgUploadItemImg) SetUrl(v string) *TaobaoItemImgUploadItemImg {
    s.Url = &v
    return s
}
func (s *TaobaoItemImgUploadItemImg) SetCreated(v util.LocalTime) *TaobaoItemImgUploadItemImg {
    s.Created = &v
    return s
}
