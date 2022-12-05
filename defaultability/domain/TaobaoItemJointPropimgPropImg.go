package domain

import (
        "topsdk/util"
    )

type TaobaoItemJointPropimgPropImg struct {
    /*
        属性图片的id，和商品相对应     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片链接地址     */
    Url  *string `json:"url,omitempty" `

    /*
        图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoItemJointPropimgPropImg) SetId(v int64) *TaobaoItemJointPropimgPropImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemJointPropimgPropImg) SetUrl(v string) *TaobaoItemJointPropimgPropImg {
    s.Url = &v
    return s
}
func (s *TaobaoItemJointPropimgPropImg) SetCreated(v util.LocalTime) *TaobaoItemJointPropimgPropImg {
    s.Created = &v
    return s
}
