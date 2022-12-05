package domain

import (
        "topsdk/util"
    )

type TaobaoItemImgDeleteItemImg struct {
    /*
        商品图片的id，和商品相对应（主图id默认为0）     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoItemImgDeleteItemImg) SetId(v int64) *TaobaoItemImgDeleteItemImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemImgDeleteItemImg) SetCreated(v util.LocalTime) *TaobaoItemImgDeleteItemImg {
    s.Created = &v
    return s
}
