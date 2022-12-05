package domain

import (
        "topsdk/util"
    )

type TaobaoItemJointImgItemImg struct {
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

func (s *TaobaoItemJointImgItemImg) SetId(v int64) *TaobaoItemJointImgItemImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemJointImgItemImg) SetUrl(v string) *TaobaoItemJointImgItemImg {
    s.Url = &v
    return s
}
func (s *TaobaoItemJointImgItemImg) SetCreated(v util.LocalTime) *TaobaoItemJointImgItemImg {
    s.Created = &v
    return s
}
