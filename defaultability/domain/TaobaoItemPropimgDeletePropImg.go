package domain

import (
        "topsdk/util"
    )

type TaobaoItemPropimgDeletePropImg struct {
    /*
        属性图片的id，和商品相对应     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoItemPropimgDeletePropImg) SetId(v int64) *TaobaoItemPropimgDeletePropImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemPropimgDeletePropImg) SetCreated(v util.LocalTime) *TaobaoItemPropimgDeletePropImg {
    s.Created = &v
    return s
}
